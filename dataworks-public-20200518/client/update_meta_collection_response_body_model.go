// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateMetaCollectionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateMetaCollectionResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *UpdateMetaCollectionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateMetaCollectionResponseBody
	GetRequestId() *string
	SetStatus(v bool) *UpdateMetaCollectionResponseBody
	GetStatus() *bool
	SetSuccess(v bool) *UpdateMetaCollectionResponseBody
	GetSuccess() *bool
}

type UpdateMetaCollectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 9999
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The specified product does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation. Valid values: true false
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s UpdateMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaCollectionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMetaCollectionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateMetaCollectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetaCollectionResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *UpdateMetaCollectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMetaCollectionResponseBody) SetErrorCode(v string) *UpdateMetaCollectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMetaCollectionResponseBody) SetErrorMessage(v string) *UpdateMetaCollectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMetaCollectionResponseBody) SetHttpStatusCode(v int32) *UpdateMetaCollectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMetaCollectionResponseBody) SetRequestId(v string) *UpdateMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetaCollectionResponseBody) SetStatus(v bool) *UpdateMetaCollectionResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateMetaCollectionResponseBody) SetSuccess(v bool) *UpdateMetaCollectionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMetaCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
