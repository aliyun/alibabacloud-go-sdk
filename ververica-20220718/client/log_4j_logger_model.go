// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLog4jLogger interface {
	dara.Model
	String() string
	GoString() string
	SetLoggerLevel(v string) *Log4jLogger
	GetLoggerLevel() *string
	SetLoggerName(v string) *Log4jLogger
	GetLoggerName() *string
}

type Log4jLogger struct {
	// example:
	//
	// ERROR
	LoggerLevel *string `json:"loggerLevel,omitempty" xml:"loggerLevel,omitempty"`
	// example:
	//
	// StdOutErrConsoleAppender
	LoggerName *string `json:"loggerName,omitempty" xml:"loggerName,omitempty"`
}

func (s Log4jLogger) String() string {
	return dara.Prettify(s)
}

func (s Log4jLogger) GoString() string {
	return s.String()
}

func (s *Log4jLogger) GetLoggerLevel() *string {
	return s.LoggerLevel
}

func (s *Log4jLogger) GetLoggerName() *string {
	return s.LoggerName
}

func (s *Log4jLogger) SetLoggerLevel(v string) *Log4jLogger {
	s.LoggerLevel = &v
	return s
}

func (s *Log4jLogger) SetLoggerName(v string) *Log4jLogger {
	s.LoggerName = &v
	return s
}

func (s *Log4jLogger) Validate() error {
	return dara.Validate(s)
}
