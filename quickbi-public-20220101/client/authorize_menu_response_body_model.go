// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeMenuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeMenuResponseBody
	GetRequestId() *string
	SetResult(v int32) *AuthorizeMenuResponseBody
	GetResult() *int32
	SetSuccess(v bool) *AuthorizeMenuResponseBody
	GetSuccess() *bool
}

type AuthorizeMenuResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 188F0B12-00EF-41B3-944A-FB7EF06C9F43
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of authorized menus.
	//
	// example:
	//
	// 2
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeMenuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeMenuResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeMenuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeMenuResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *AuthorizeMenuResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthorizeMenuResponseBody) SetRequestId(v string) *AuthorizeMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeMenuResponseBody) SetResult(v int32) *AuthorizeMenuResponseBody {
	s.Result = &v
	return s
}

func (s *AuthorizeMenuResponseBody) SetSuccess(v bool) *AuthorizeMenuResponseBody {
	s.Success = &v
	return s
}

func (s *AuthorizeMenuResponseBody) Validate() error {
	return dara.Validate(s)
}
