// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowsersResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBrowsersResult
	GetCode() *string
	SetData(v *ListBrowsersOutput) *ListBrowsersResult
	GetData() *ListBrowsersOutput
	SetRequestId(v string) *ListBrowsersResult
	GetRequestId() *string
}

type ListBrowsersResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 浏览器列表的详细信息
	//
	// example:
	//
	// {}
	Data *ListBrowsersOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListBrowsersResult) String() string {
	return dara.Prettify(s)
}

func (s ListBrowsersResult) GoString() string {
	return s.String()
}

func (s *ListBrowsersResult) GetCode() *string {
	return s.Code
}

func (s *ListBrowsersResult) GetData() *ListBrowsersOutput {
	return s.Data
}

func (s *ListBrowsersResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBrowsersResult) SetCode(v string) *ListBrowsersResult {
	s.Code = &v
	return s
}

func (s *ListBrowsersResult) SetData(v *ListBrowsersOutput) *ListBrowsersResult {
	s.Data = v
	return s
}

func (s *ListBrowsersResult) SetRequestId(v string) *ListBrowsersResult {
	s.RequestId = &v
	return s
}

func (s *ListBrowsersResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
