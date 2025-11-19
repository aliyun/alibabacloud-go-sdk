// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSandboxHealthCheckResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SandboxHealthCheckResult
	GetCode() *string
	SetData(v *SandboxHealthCheckOut) *SandboxHealthCheckResult
	GetData() *SandboxHealthCheckOut
	SetRequestId(v string) *SandboxHealthCheckResult
	GetRequestId() *string
}

type SandboxHealthCheckResult struct {
	// SUCCESS 为成功
	Code *string                `json:"code,omitempty" xml:"code,omitempty"`
	Data *SandboxHealthCheckOut `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SandboxHealthCheckResult) String() string {
	return dara.Prettify(s)
}

func (s SandboxHealthCheckResult) GoString() string {
	return s.String()
}

func (s *SandboxHealthCheckResult) GetCode() *string {
	return s.Code
}

func (s *SandboxHealthCheckResult) GetData() *SandboxHealthCheckOut {
	return s.Data
}

func (s *SandboxHealthCheckResult) GetRequestId() *string {
	return s.RequestId
}

func (s *SandboxHealthCheckResult) SetCode(v string) *SandboxHealthCheckResult {
	s.Code = &v
	return s
}

func (s *SandboxHealthCheckResult) SetData(v *SandboxHealthCheckOut) *SandboxHealthCheckResult {
	s.Data = v
	return s
}

func (s *SandboxHealthCheckResult) SetRequestId(v string) *SandboxHealthCheckResult {
	s.RequestId = &v
	return s
}

func (s *SandboxHealthCheckResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
