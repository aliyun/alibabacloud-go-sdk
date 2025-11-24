// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSandboxResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopSandboxResult
	GetCode() *string
	SetData(v *Sandbox) *StopSandboxResult
	GetData() *Sandbox
	SetRequestId(v string) *StopSandboxResult
	GetRequestId() *string
}

type StopSandboxResult struct {
	Code      *string  `json:"code,omitempty" xml:"code,omitempty"`
	Data      *Sandbox `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopSandboxResult) String() string {
	return dara.Prettify(s)
}

func (s StopSandboxResult) GoString() string {
	return s.String()
}

func (s *StopSandboxResult) GetCode() *string {
	return s.Code
}

func (s *StopSandboxResult) GetData() *Sandbox {
	return s.Data
}

func (s *StopSandboxResult) GetRequestId() *string {
	return s.RequestId
}

func (s *StopSandboxResult) SetCode(v string) *StopSandboxResult {
	s.Code = &v
	return s
}

func (s *StopSandboxResult) SetData(v *Sandbox) *StopSandboxResult {
	s.Data = v
	return s
}

func (s *StopSandboxResult) SetRequestId(v string) *StopSandboxResult {
	s.RequestId = &v
	return s
}

func (s *StopSandboxResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
