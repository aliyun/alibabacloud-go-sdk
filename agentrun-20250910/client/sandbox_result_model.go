// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSandboxResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SandboxResult
	GetCode() *string
	SetData(v *Sandbox) *SandboxResult
	GetData() *Sandbox
	SetRequestId(v string) *SandboxResult
	GetRequestId() *string
}

type SandboxResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 沙箱的详细信息
	//
	// This parameter is required.
	Data *Sandbox `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SandboxResult) String() string {
	return dara.Prettify(s)
}

func (s SandboxResult) GoString() string {
	return s.String()
}

func (s *SandboxResult) GetCode() *string {
	return s.Code
}

func (s *SandboxResult) GetData() *Sandbox {
	return s.Data
}

func (s *SandboxResult) GetRequestId() *string {
	return s.RequestId
}

func (s *SandboxResult) SetCode(v string) *SandboxResult {
	s.Code = &v
	return s
}

func (s *SandboxResult) SetData(v *Sandbox) *SandboxResult {
	s.Data = v
	return s
}

func (s *SandboxResult) SetRequestId(v string) *SandboxResult {
	s.RequestId = &v
	return s
}

func (s *SandboxResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
