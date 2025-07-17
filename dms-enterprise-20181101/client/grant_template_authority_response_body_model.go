// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantTemplateAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GrantTemplateAuthorityResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GrantTemplateAuthorityResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GrantTemplateAuthorityResponseBody
	GetRequestId() *string
	SetResult(v bool) *GrantTemplateAuthorityResponseBody
	GetResult() *bool
	SetSuccess(v bool) *GrantTemplateAuthorityResponseBody
	GetSuccess() *bool
}

type GrantTemplateAuthorityResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34E01EDD-6A16-4CF0-9541-C644D1BE01AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the permissions on resources were granted to the users by using the permission template.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s GrantTemplateAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantTemplateAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *GrantTemplateAuthorityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GrantTemplateAuthorityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GrantTemplateAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantTemplateAuthorityResponseBody) GetResult() *bool {
	return s.Result
}

func (s *GrantTemplateAuthorityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GrantTemplateAuthorityResponseBody) SetErrorCode(v string) *GrantTemplateAuthorityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GrantTemplateAuthorityResponseBody) SetErrorMessage(v string) *GrantTemplateAuthorityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GrantTemplateAuthorityResponseBody) SetRequestId(v string) *GrantTemplateAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantTemplateAuthorityResponseBody) SetResult(v bool) *GrantTemplateAuthorityResponseBody {
	s.Result = &v
	return s
}

func (s *GrantTemplateAuthorityResponseBody) SetSuccess(v bool) *GrantTemplateAuthorityResponseBody {
	s.Success = &v
	return s
}

func (s *GrantTemplateAuthorityResponseBody) Validate() error {
	return dara.Validate(s)
}
