// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTemplateAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RevokeTemplateAuthorityResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RevokeTemplateAuthorityResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RevokeTemplateAuthorityResponseBody
	GetRequestId() *string
	SetResult(v bool) *RevokeTemplateAuthorityResponseBody
	GetResult() *bool
	SetSuccess(v bool) *RevokeTemplateAuthorityResponseBody
	GetSuccess() *bool
}

type RevokeTemplateAuthorityResponseBody struct {
	// The error code returned if the request failed.
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
	// C51420E3-144A-4A94-B473-8662FCF4AD10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the permissions were revoked from the users.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeTemplateAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeTemplateAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeTemplateAuthorityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RevokeTemplateAuthorityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RevokeTemplateAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeTemplateAuthorityResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RevokeTemplateAuthorityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeTemplateAuthorityResponseBody) SetErrorCode(v string) *RevokeTemplateAuthorityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RevokeTemplateAuthorityResponseBody) SetErrorMessage(v string) *RevokeTemplateAuthorityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RevokeTemplateAuthorityResponseBody) SetRequestId(v string) *RevokeTemplateAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeTemplateAuthorityResponseBody) SetResult(v bool) *RevokeTemplateAuthorityResponseBody {
	s.Result = &v
	return s
}

func (s *RevokeTemplateAuthorityResponseBody) SetSuccess(v bool) *RevokeTemplateAuthorityResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeTemplateAuthorityResponseBody) Validate() error {
	return dara.Validate(s)
}
