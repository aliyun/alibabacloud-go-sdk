// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeInterpretersResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCodeInterpretersResult
	GetCode() *string
	SetData(v *ListCodeInterpretersOutput) *ListCodeInterpretersResult
	GetData() *ListCodeInterpretersOutput
	SetRequestId(v string) *ListCodeInterpretersResult
	GetRequestId() *string
}

type ListCodeInterpretersResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 代码解释器列表的详细信息
	//
	// example:
	//
	// {}
	Data *ListCodeInterpretersOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListCodeInterpretersResult) String() string {
	return dara.Prettify(s)
}

func (s ListCodeInterpretersResult) GoString() string {
	return s.String()
}

func (s *ListCodeInterpretersResult) GetCode() *string {
	return s.Code
}

func (s *ListCodeInterpretersResult) GetData() *ListCodeInterpretersOutput {
	return s.Data
}

func (s *ListCodeInterpretersResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCodeInterpretersResult) SetCode(v string) *ListCodeInterpretersResult {
	s.Code = &v
	return s
}

func (s *ListCodeInterpretersResult) SetData(v *ListCodeInterpretersOutput) *ListCodeInterpretersResult {
	s.Data = v
	return s
}

func (s *ListCodeInterpretersResult) SetRequestId(v string) *ListCodeInterpretersResult {
	s.RequestId = &v
	return s
}

func (s *ListCodeInterpretersResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
