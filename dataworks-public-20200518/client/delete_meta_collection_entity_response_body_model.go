// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCollectionEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteMetaCollectionEntityResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteMetaCollectionEntityResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteMetaCollectionEntityResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteMetaCollectionEntityResponseBody
	GetRequestId() *string
	SetStatus(v bool) *DeleteMetaCollectionEntityResponseBody
	GetStatus() *bool
	SetSuccess(v bool) *DeleteMetaCollectionEntityResponseBody
	GetSuccess() *bool
}

type DeleteMetaCollectionEntityResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 999999
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// entity not exist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation. Valid values:
	//
	// true: succeeded
	//
	// false: failed
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s DeleteMetaCollectionEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCollectionEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetaCollectionEntityResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMetaCollectionEntityResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteMetaCollectionEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMetaCollectionEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetaCollectionEntityResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *DeleteMetaCollectionEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetaCollectionEntityResponseBody) SetErrorCode(v string) *DeleteMetaCollectionEntityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMetaCollectionEntityResponseBody) SetErrorMessage(v string) *DeleteMetaCollectionEntityResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMetaCollectionEntityResponseBody) SetHttpStatusCode(v int32) *DeleteMetaCollectionEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMetaCollectionEntityResponseBody) SetRequestId(v string) *DeleteMetaCollectionEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetaCollectionEntityResponseBody) SetStatus(v bool) *DeleteMetaCollectionEntityResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteMetaCollectionEntityResponseBody) SetSuccess(v bool) *DeleteMetaCollectionEntityResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetaCollectionEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
