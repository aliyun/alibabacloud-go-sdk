// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSandboxesResult
	GetCode() *string
	SetData(v *ListSandboxesOutput) *ListSandboxesResult
	GetData() *ListSandboxesOutput
	SetRequestId(v string) *ListSandboxesResult
	GetRequestId() *string
}

type ListSandboxesResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 沙箱列表的详细信息
	Data *ListSandboxesOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSandboxesResult) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxesResult) GoString() string {
	return s.String()
}

func (s *ListSandboxesResult) GetCode() *string {
	return s.Code
}

func (s *ListSandboxesResult) GetData() *ListSandboxesOutput {
	return s.Data
}

func (s *ListSandboxesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSandboxesResult) SetCode(v string) *ListSandboxesResult {
	s.Code = &v
	return s
}

func (s *ListSandboxesResult) SetData(v *ListSandboxesOutput) *ListSandboxesResult {
	s.Data = v
	return s
}

func (s *ListSandboxesResult) SetRequestId(v string) *ListSandboxesResult {
	s.RequestId = &v
	return s
}

func (s *ListSandboxesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
