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
	// 代码解释器会话的环境变量配置
	Environment map[string]*string `json:"environment,omitempty" xml:"environment,omitempty"`
	// 代码解释器会话的超时时间，单位为秒
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 代码解释器会话的工作目录路径
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
