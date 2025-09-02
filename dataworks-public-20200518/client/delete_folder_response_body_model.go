// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteFolderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteFolderResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteFolderResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteFolderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFolderResponseBody
	GetSuccess() *bool
}

type DeleteFolderResponseBody struct {
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

func (s DeleteFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteFolderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFolderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFolderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFolderResponseBody) SetErrorCode(v string) *DeleteFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFolderResponseBody) SetErrorMessage(v string) *DeleteFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFolderResponseBody) SetHttpStatusCode(v int32) *DeleteFolderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFolderResponseBody) SetRequestId(v string) *DeleteFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFolderResponseBody) SetSuccess(v bool) *DeleteFolderResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
