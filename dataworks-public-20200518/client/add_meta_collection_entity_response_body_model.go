// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMetaCollectionEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddMetaCollectionEntityResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddMetaCollectionEntityResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *AddMetaCollectionEntityResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AddMetaCollectionEntityResponseBody
	GetRequestId() *string
	SetStatus(v bool) *AddMetaCollectionEntityResponseBody
	GetStatus() *bool
	SetSuccess(v bool) *AddMetaCollectionEntityResponseBody
	GetSuccess() *bool
}

type AddMetaCollectionEntityResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 9999
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
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
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation. Valid values:
	//
	// 	- true: succeeded
	//
	// 	- false: failed
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMetaCollectionEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMetaCollectionEntityResponseBody) GoString() string {
	return s.String()
}

func (s *AddMetaCollectionEntityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddMetaCollectionEntityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddMetaCollectionEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddMetaCollectionEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMetaCollectionEntityResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *AddMetaCollectionEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddMetaCollectionEntityResponseBody) SetErrorCode(v string) *AddMetaCollectionEntityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddMetaCollectionEntityResponseBody) SetErrorMessage(v string) *AddMetaCollectionEntityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddMetaCollectionEntityResponseBody) SetHttpStatusCode(v int32) *AddMetaCollectionEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddMetaCollectionEntityResponseBody) SetRequestId(v string) *AddMetaCollectionEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMetaCollectionEntityResponseBody) SetStatus(v bool) *AddMetaCollectionEntityResponseBody {
	s.Status = &v
	return s
}

func (s *AddMetaCollectionEntityResponseBody) SetSuccess(v bool) *AddMetaCollectionEntityResponseBody {
	s.Success = &v
	return s
}

func (s *AddMetaCollectionEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
