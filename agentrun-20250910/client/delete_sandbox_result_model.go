// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSandboxResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSandboxResult
	GetCode() *string
	SetData(v *Sandbox) *DeleteSandboxResult
	GetData() *Sandbox
	SetRequestId(v string) *DeleteSandboxResult
	GetRequestId() *string
}

type DeleteSandboxResult struct {
	// SUCCESS 为成功
	Code *string  `json:"code,omitempty" xml:"code,omitempty"`
	Data *Sandbox `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSandboxResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteSandboxResult) GoString() string {
	return s.String()
}

func (s *DeleteSandboxResult) GetCode() *string {
	return s.Code
}

func (s *DeleteSandboxResult) GetData() *Sandbox {
	return s.Data
}

func (s *DeleteSandboxResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSandboxResult) SetCode(v string) *DeleteSandboxResult {
	s.Code = &v
	return s
}

func (s *DeleteSandboxResult) SetData(v *Sandbox) *DeleteSandboxResult {
	s.Data = v
	return s
}

func (s *DeleteSandboxResult) SetRequestId(v string) *DeleteSandboxResult {
	s.RequestId = &v
	return s
}

func (s *DeleteSandboxResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
