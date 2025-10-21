// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogging interface {
	dara.Model
	String() string
	GoString() string
	SetLog4j2ConfigurationTemplate(v string) *Logging
	GetLog4j2ConfigurationTemplate() *string
	SetLog4jLoggers(v []*Log4jLogger) *Logging
	GetLog4jLoggers() []*Log4jLogger
	SetLogReservePolicy(v *LogReservePolicy) *Logging
	GetLogReservePolicy() *LogReservePolicy
	SetLoggingProfile(v string) *Logging
	GetLoggingProfile() *string
}

type Logging struct {
	// example:
	//
	// xml格式文本
	Log4j2ConfigurationTemplate *string           `json:"log4j2ConfigurationTemplate,omitempty" xml:"log4j2ConfigurationTemplate,omitempty"`
	Log4jLoggers                []*Log4jLogger    `json:"log4jLoggers,omitempty" xml:"log4jLoggers,omitempty" type:"Repeated"`
	LogReservePolicy            *LogReservePolicy `json:"logReservePolicy,omitempty" xml:"logReservePolicy,omitempty"`
	// example:
	//
	// oss
	LoggingProfile *string `json:"loggingProfile,omitempty" xml:"loggingProfile,omitempty"`
}

func (s Logging) String() string {
	return dara.Prettify(s)
}

func (s Logging) GoString() string {
	return s.String()
}

func (s *Logging) GetLog4j2ConfigurationTemplate() *string {
	return s.Log4j2ConfigurationTemplate
}

func (s *Logging) GetLog4jLoggers() []*Log4jLogger {
	return s.Log4jLoggers
}

func (s *Logging) GetLogReservePolicy() *LogReservePolicy {
	return s.LogReservePolicy
}

func (s *Logging) GetLoggingProfile() *string {
	return s.LoggingProfile
}

func (s *Logging) SetLog4j2ConfigurationTemplate(v string) *Logging {
	s.Log4j2ConfigurationTemplate = &v
	return s
}

func (s *Logging) SetLog4jLoggers(v []*Log4jLogger) *Logging {
	s.Log4jLoggers = v
	return s
}

func (s *Logging) SetLogReservePolicy(v *LogReservePolicy) *Logging {
	s.LogReservePolicy = v
	return s
}

func (s *Logging) SetLoggingProfile(v string) *Logging {
	s.LoggingProfile = &v
	return s
}

func (s *Logging) Validate() error {
	if s.Log4jLoggers != nil {
		for _, item := range s.Log4jLoggers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogReservePolicy != nil {
		if err := s.LogReservePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
