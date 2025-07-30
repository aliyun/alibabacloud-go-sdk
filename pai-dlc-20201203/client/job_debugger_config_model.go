// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobDebuggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDebuggerConfigContent(v string) *JobDebuggerConfig
	GetDebuggerConfigContent() *string
	SetDebuggerConfigId(v string) *JobDebuggerConfig
	GetDebuggerConfigId() *string
	SetGmtCreateTime(v string) *JobDebuggerConfig
	GetGmtCreateTime() *string
	SetJobId(v string) *JobDebuggerConfig
	GetJobId() *string
}

type JobDebuggerConfig struct {
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	DebuggerConfigId      *string `json:"DebuggerConfigId,omitempty" xml:"DebuggerConfigId,omitempty"`
	GmtCreateTime         *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	JobId                 *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s JobDebuggerConfig) String() string {
	return dara.Prettify(s)
}

func (s JobDebuggerConfig) GoString() string {
	return s.String()
}

func (s *JobDebuggerConfig) GetDebuggerConfigContent() *string {
	return s.DebuggerConfigContent
}

func (s *JobDebuggerConfig) GetDebuggerConfigId() *string {
	return s.DebuggerConfigId
}

func (s *JobDebuggerConfig) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *JobDebuggerConfig) GetJobId() *string {
	return s.JobId
}

func (s *JobDebuggerConfig) SetDebuggerConfigContent(v string) *JobDebuggerConfig {
	s.DebuggerConfigContent = &v
	return s
}

func (s *JobDebuggerConfig) SetDebuggerConfigId(v string) *JobDebuggerConfig {
	s.DebuggerConfigId = &v
	return s
}

func (s *JobDebuggerConfig) SetGmtCreateTime(v string) *JobDebuggerConfig {
	s.GmtCreateTime = &v
	return s
}

func (s *JobDebuggerConfig) SetJobId(v string) *JobDebuggerConfig {
	s.JobId = &v
	return s
}

func (s *JobDebuggerConfig) Validate() error {
	return dara.Validate(s)
}
