// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateSlrPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateSlrPermissionResponseBody
	GetRequestId() *string
	SetResult(v bool) *ValidateSlrPermissionResponseBody
	GetResult() *bool
}

type ValidateSlrPermissionResponseBody struct {
	// example:
	//
	// BC4ED7DD-8C84-49B5-8A95-456F82E44D13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateSlrPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateSlrPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateSlrPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateSlrPermissionResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ValidateSlrPermissionResponseBody) SetRequestId(v string) *ValidateSlrPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateSlrPermissionResponseBody) SetResult(v bool) *ValidateSlrPermissionResponseBody {
	s.Result = &v
	return s
}

func (s *ValidateSlrPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
