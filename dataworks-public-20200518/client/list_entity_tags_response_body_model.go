// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntityTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*UserEntityTag) *ListEntityTagsResponseBody
	GetData() []*UserEntityTag
	SetErrorCode(v string) *ListEntityTagsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListEntityTagsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListEntityTagsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListEntityTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEntityTagsResponseBody
	GetSuccess() *bool
}

type ListEntityTagsResponseBody struct {
	// The tags.
	Data []*UserEntityTag `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 101011005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid.Entity.EntityTypeNotSupported
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true\\
	//
	// false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEntityTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEntityTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntityTagsResponseBody) GetData() []*UserEntityTag {
	return s.Data
}

func (s *ListEntityTagsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListEntityTagsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListEntityTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListEntityTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEntityTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEntityTagsResponseBody) SetData(v []*UserEntityTag) *ListEntityTagsResponseBody {
	s.Data = v
	return s
}

func (s *ListEntityTagsResponseBody) SetErrorCode(v string) *ListEntityTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEntityTagsResponseBody) SetErrorMessage(v string) *ListEntityTagsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListEntityTagsResponseBody) SetHttpStatusCode(v int32) *ListEntityTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEntityTagsResponseBody) SetRequestId(v string) *ListEntityTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntityTagsResponseBody) SetSuccess(v bool) *ListEntityTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEntityTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
