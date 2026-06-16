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
	// The operation status code. `SUCCESS` indicates success. A failed operation returns an error code, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned browser object.
	//
	// example:
	//
	// {}
	Data *Browser `json:"data,omitempty" xml:"data,omitempty"`
	// The unique request identifier. Use it for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
