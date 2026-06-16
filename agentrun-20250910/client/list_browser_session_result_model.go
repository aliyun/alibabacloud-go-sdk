// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserSessionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBrowserSessionResult
	GetCode() *string
	SetData(v *BrowserSessionListOut) *ListBrowserSessionResult
	GetData() *BrowserSessionListOut
	SetRequestId(v string) *ListBrowserSessionResult
	GetRequestId() *string
}

type ListBrowserSessionResult struct {
	// The response code. A value of SUCCESS indicates the request succeeded. On failure, an error code is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The operation\\"s payload.
	//
	// example:
	//
	// {}
	Data *BrowserSessionListOut `json:"data,omitempty" xml:"data,omitempty"`
	// The unique identifier for the request.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListBrowserSessionResult) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserSessionResult) GoString() string {
	return s.String()
}

func (s *ListBrowserSessionResult) GetCode() *string {
	return s.Code
}

func (s *ListBrowserSessionResult) GetData() *BrowserSessionListOut {
	return s.Data
}

func (s *ListBrowserSessionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBrowserSessionResult) SetCode(v string) *ListBrowserSessionResult {
	s.Code = &v
	return s
}

func (s *ListBrowserSessionResult) SetData(v *BrowserSessionListOut) *ListBrowserSessionResult {
	s.Data = v
	return s
}

func (s *ListBrowserSessionResult) SetRequestId(v string) *ListBrowserSessionResult {
	s.RequestId = &v
	return s
}

func (s *ListBrowserSessionResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
