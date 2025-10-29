// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBrowserSessionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBrowserSessionResult
	GetCode() *string
	SetData(v *BrowserSessionOut) *GetBrowserSessionResult
	GetData() *BrowserSessionOut
	SetRequestId(v string) *GetBrowserSessionResult
	GetRequestId() *string
}

type GetBrowserSessionResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {}
	Data *BrowserSessionOut `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetBrowserSessionResult) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserSessionResult) GoString() string {
	return s.String()
}

func (s *GetBrowserSessionResult) GetCode() *string {
	return s.Code
}

func (s *GetBrowserSessionResult) GetData() *BrowserSessionOut {
	return s.Data
}

func (s *GetBrowserSessionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBrowserSessionResult) SetCode(v string) *GetBrowserSessionResult {
	s.Code = &v
	return s
}

func (s *GetBrowserSessionResult) SetData(v *BrowserSessionOut) *GetBrowserSessionResult {
	s.Data = v
	return s
}

func (s *GetBrowserSessionResult) SetRequestId(v string) *GetBrowserSessionResult {
	s.RequestId = &v
	return s
}

func (s *GetBrowserSessionResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
