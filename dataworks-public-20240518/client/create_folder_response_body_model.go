// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateFolderResponseBody
	GetData() *string
	SetErrorCode(v string) *CreateFolderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateFolderResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *CreateFolderResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateFolderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFolderResponseBody
	GetSuccess() *bool
}

type CreateFolderResponseBody struct {
	// The unique identifier of the newly created folder.
	//
	// example:
	//
	// bdfd68****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. Used for troubleshooting when errors occur.
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded. Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFolderResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateFolderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateFolderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateFolderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFolderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFolderResponseBody) SetData(v string) *CreateFolderResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFolderResponseBody) SetErrorCode(v string) *CreateFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFolderResponseBody) SetErrorMessage(v string) *CreateFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFolderResponseBody) SetHttpStatusCode(v int32) *CreateFolderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFolderResponseBody) SetRequestId(v string) *CreateFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFolderResponseBody) SetSuccess(v bool) *CreateFolderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
