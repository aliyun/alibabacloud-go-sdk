// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateFolderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateFolderResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateFolderResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateFolderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFolderResponseBody
	GetSuccess() *bool
}

type UpdateFolderResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFolderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateFolderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateFolderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFolderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFolderResponseBody) SetErrorCode(v string) *UpdateFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateFolderResponseBody) SetErrorMessage(v string) *UpdateFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFolderResponseBody) SetHttpStatusCode(v int32) *UpdateFolderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateFolderResponseBody) SetRequestId(v string) *UpdateFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFolderResponseBody) SetSuccess(v bool) *UpdateFolderResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
