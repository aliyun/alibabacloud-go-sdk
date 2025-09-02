// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCollectionDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v *Collection) *GetMetaCollectionDetailResponseBody
	GetCollection() *Collection
	SetErrorCode(v string) *GetMetaCollectionDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaCollectionDetailResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaCollectionDetailResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaCollectionDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaCollectionDetailResponseBody
	GetSuccess() *bool
}

type GetMetaCollectionDetailResponseBody struct {
	// The information about the collection.
	Collection *Collection `json:"Collection,omitempty" xml:"Collection,omitempty"`
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
	// album.xxxx does not exist.
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

func (s GetMetaCollectionDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCollectionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaCollectionDetailResponseBody) GetCollection() *Collection {
	return s.Collection
}

func (s *GetMetaCollectionDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaCollectionDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaCollectionDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaCollectionDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaCollectionDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaCollectionDetailResponseBody) SetCollection(v *Collection) *GetMetaCollectionDetailResponseBody {
	s.Collection = v
	return s
}

func (s *GetMetaCollectionDetailResponseBody) SetErrorCode(v string) *GetMetaCollectionDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaCollectionDetailResponseBody) SetErrorMessage(v string) *GetMetaCollectionDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaCollectionDetailResponseBody) SetHttpStatusCode(v int32) *GetMetaCollectionDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaCollectionDetailResponseBody) SetRequestId(v string) *GetMetaCollectionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaCollectionDetailResponseBody) SetSuccess(v bool) *GetMetaCollectionDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaCollectionDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
