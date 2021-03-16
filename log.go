package logsrus

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func InitialLog() {
	fmt.Println("Just initial")
	defer func() {
		fmt.Println("Here")
	}()

	fmt.Println("Ends")
}

func BasicLog() {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	// log.Formatter = new(logrus.TextFormatter)                     //default
	// log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	// log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }

	defer func() {
		fmt.Println("Where's the defer comes")
		err := recover()
		if err != nil {
			// fmt.Println("Where a u", err)
			entry := err.(*logrus.Entry)
			log.WithFields(logrus.Fields{
				"omg":         true,
				"err_animal":  entry.Data["animal"],
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
				"number":      100,
			}).Error("The ice breaks!") // or use Fatal() to force the process to exit with a nonzero code
		}
	}()

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 0,
	}).Trace("Went to the beach")

	// log.WithFields(logrus.Fields{
	// 	"animal": "walrus",
	// 	"number": 8,
	// }).Debug("Started observing beach")

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// log.WithFields(logrus.Fields{
	// 	"omg":    true,
	// 	"number": 122,
	// }).Warn("The group's number increased tremendously!")

	// log.WithFields(logrus.Fields{
	// 	"temperature": -4,
	// }).Debug("Temperature changes")

	log.WithFields(logrus.Fields{
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")
}

func BasicAPI() {
	app := fiber.New()
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Exit with :", err)
		}
	}()

	app.Use(requestid.New(requestid.Config{
		Header: fiber.HeaderXRequestID,
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	app.Get("/characters", func(c *fiber.Ctx) error {
		// log.WithField("RequestID", c.Response().Header.Peek(fiber.HeaderXRequestID))

		log.WithFields(logrus.Fields{
			"requestid": c.Response().Header.Peek(fiber.HeaderXRequestID),
			"endpoint":  string(c.Request().URI().Path()),
			"from":      c.IP(),
		}).Info()

		return c.Status(200).JSON("You Got .....!!!!")
	})

	log.Panic(app.Listen(":8080"))
}
