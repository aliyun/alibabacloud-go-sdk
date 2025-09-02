// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntityTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RemoveEntityTagsResponseBody
	GetData() *bool
	SetErrorCode(v string) *RemoveEntityTagsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RemoveEntityTagsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *RemoveEntityTagsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RemoveEntityTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveEntityTagsResponseBody
	GetSuccess() *bool
}

type RemoveEntityTagsResponseBody struct {
	// Indicates whether the call was successful. Valid values:
	//
	// true\\
	//
	// false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s RemoveEntityTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntityTagsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntityTagsResponseBody) GetData() *bool {
	return s.Data
}

func (s *RemoveEntityTagsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RemoveEntityTagsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RemoveEntityTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveEntityTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveEntityTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveEntityTagsResponseBody) SetData(v bool) *RemoveEntityTagsResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveEntityTagsResponseBody) SetErrorCode(v string) *RemoveEntityTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemoveEntityTagsResponseBody) SetErrorMessage(v string) *RemoveEntityTagsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveEntityTagsResponseBody) SetHttpStatusCode(v int32) *RemoveEntityTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveEntityTagsResponseBody) SetRequestId(v string) *RemoveEntityTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveEntityTagsResponseBody) SetSuccess(v bool) *RemoveEntityTagsResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveEntityTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
