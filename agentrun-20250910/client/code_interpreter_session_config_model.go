// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeInterpreterSessionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v map[string]*string) *CodeInterpreterSessionConfig
	GetEnvironment() map[string]*string
	SetTimeout(v int32) *CodeInterpreterSessionConfig
	GetTimeout() *int32
	SetWorkingDirectory(v string) *CodeInterpreterSessionConfig
	GetWorkingDirectory() *string
}

type CodeInterpreterSessionConfig struct {
	// The environment variables for the session, specified as key-value pairs.
	//
	// example:
	//
	// PYTHONPATH=/app,DEBUG=true
	Environment map[string]*string `json:"environment" xml:"environment"`
	// The timeout for the session, in seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The working directory path for the session.
	//
	// example:
	//
	// /app
	WorkingDirectory *string `json:"workingDirectory,omitempty" xml:"workingDirectory,omitempty"`
}

func (s CodeInterpreterSessionConfig) String() string {
	return dara.Prettify(s)
}

func (s CodeInterpreterSessionConfig) GoString() string {
	return s.String()
}

func (s *CodeInterpreterSessionConfig) GetEnvironment() map[string]*string {
	return s.Environment
}

func (s *CodeInterpreterSessionConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CodeInterpreterSessionConfig) GetWorkingDirectory() *string {
	return s.WorkingDirectory
}

func (s *CodeInterpreterSessionConfig) SetEnvironment(v map[string]*string) *CodeInterpreterSessionConfig {
	s.Environment = v
	return s
}

func (s *CodeInterpreterSessionConfig) SetTimeout(v int32) *CodeInterpreterSessionConfig {
	s.Timeout = &v
	return s
}

func (s *CodeInterpreterSessionConfig) SetWorkingDirectory(v string) *CodeInterpreterSessionConfig {
	s.WorkingDirectory = &v
	return s
}

func (s *CodeInterpreterSessionConfig) Validate() error {
	return dara.Validate(s)
}
