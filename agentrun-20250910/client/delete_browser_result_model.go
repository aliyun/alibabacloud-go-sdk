// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBrowserResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBrowserResult
	GetCode() *string
	SetData(v *Browser) *DeleteBrowserResult
	GetData() *Browser
	SetRequestId(v string) *DeleteBrowserResult
	GetRequestId() *string
}

type DeleteBrowserResult struct {
	// The operation result code. `SUCCESS` indicates a successful operation. Other values indicate an error type, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Details about the deleted browser.
	//
	// example:
	//
	// {}
	Data *Browser `json:"data,omitempty" xml:"data,omitempty"`
	// A unique request ID for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteBrowserResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteBrowserResult) GoString() string {
	return s.String()
}

func (s *DeleteBrowserResult) GetCode() *string {
	return s.Code
}

func (s *DeleteBrowserResult) GetData() *Browser {
	return s.Data
}

func (s *DeleteBrowserResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBrowserResult) SetCode(v string) *DeleteBrowserResult {
	s.Code = &v
	return s
}

func (s *DeleteBrowserResult) SetData(v *Browser) *DeleteBrowserResult {
	s.Data = v
	return s
}

func (s *DeleteBrowserResult) SetRequestId(v string) *DeleteBrowserResult {
	s.RequestId = &v
	return s
}

func (s *DeleteBrowserResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
