// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLog interface {
	dara.Model
	String() string
	GoString() string
	SetDriverStartup(v string) *RunLog
	GetDriverStartup() *string
	SetDriverStdError(v string) *RunLog
	GetDriverStdError() *string
	SetDriverStdOut(v string) *RunLog
	GetDriverStdOut() *string
	SetDriverSyslog(v string) *RunLog
	GetDriverSyslog() *string
}

type RunLog struct {
	DriverStartup  *string `json:"driverStartup,omitempty" xml:"driverStartup,omitempty"`
	DriverStdError *string `json:"driverStdError,omitempty" xml:"driverStdError,omitempty"`
	DriverStdOut   *string `json:"driverStdOut,omitempty" xml:"driverStdOut,omitempty"`
	DriverSyslog   *string `json:"driverSyslog,omitempty" xml:"driverSyslog,omitempty"`
}

func (s RunLog) String() string {
	return dara.Prettify(s)
}

func (s RunLog) GoString() string {
	return s.String()
}

func (s *RunLog) GetDriverStartup() *string {
	return s.DriverStartup
}

func (s *RunLog) GetDriverStdError() *string {
	return s.DriverStdError
}

func (s *RunLog) GetDriverStdOut() *string {
	return s.DriverStdOut
}

func (s *RunLog) GetDriverSyslog() *string {
	return s.DriverSyslog
}

func (s *RunLog) SetDriverStartup(v string) *RunLog {
	s.DriverStartup = &v
	return s
}

func (s *RunLog) SetDriverStdError(v string) *RunLog {
	s.DriverStdError = &v
	return s
}

func (s *RunLog) SetDriverStdOut(v string) *RunLog {
	s.DriverStdOut = &v
	return s
}

func (s *RunLog) SetDriverSyslog(v string) *RunLog {
	s.DriverSyslog = &v
	return s
}

func (s *RunLog) Validate() error {
	return dara.Validate(s)
}
