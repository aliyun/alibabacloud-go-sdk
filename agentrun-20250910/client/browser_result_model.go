// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BrowserResult
	GetCode() *string
	SetData(v *Browser) *BrowserResult
	GetData() *Browser
	SetRequestId(v string) *BrowserResult
	GetRequestId() *string
}

type BrowserResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 浏览器的详细信息
	Data *Browser `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BrowserResult) String() string {
	return dara.Prettify(s)
}

func (s BrowserResult) GoString() string {
	return s.String()
}

func (s *BrowserResult) GetCode() *string {
	return s.Code
}

func (s *BrowserResult) GetData() *Browser {
	return s.Data
}

func (s *BrowserResult) GetRequestId() *string {
	return s.RequestId
}

func (s *BrowserResult) SetCode(v string) *BrowserResult {
	s.Code = &v
	return s
}

func (s *BrowserResult) SetData(v *Browser) *BrowserResult {
	s.Data = v
	return s
}

func (s *BrowserResult) SetRequestId(v string) *BrowserResult {
	s.RequestId = &v
	return s
}

func (s *BrowserResult) Validate() error {
	return dara.Validate(s)
}
