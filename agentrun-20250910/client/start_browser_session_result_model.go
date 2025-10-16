// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBrowserSessionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartBrowserSessionResult
	GetCode() *string
	SetData(v *BrowserSessionOut) *StartBrowserSessionResult
	GetData() *BrowserSessionOut
	SetRequestId(v string) *StartBrowserSessionResult
	GetRequestId() *string
}

type StartBrowserSessionResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code      *string            `json:"code,omitempty" xml:"code,omitempty"`
	Data      *BrowserSessionOut `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartBrowserSessionResult) String() string {
	return dara.Prettify(s)
}

func (s StartBrowserSessionResult) GoString() string {
	return s.String()
}

func (s *StartBrowserSessionResult) GetCode() *string {
	return s.Code
}

func (s *StartBrowserSessionResult) GetData() *BrowserSessionOut {
	return s.Data
}

func (s *StartBrowserSessionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *StartBrowserSessionResult) SetCode(v string) *StartBrowserSessionResult {
	s.Code = &v
	return s
}

func (s *StartBrowserSessionResult) SetData(v *BrowserSessionOut) *StartBrowserSessionResult {
	s.Data = v
	return s
}

func (s *StartBrowserSessionResult) SetRequestId(v string) *StartBrowserSessionResult {
	s.RequestId = &v
	return s
}

func (s *StartBrowserSessionResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
