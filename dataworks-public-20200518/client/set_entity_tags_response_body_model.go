// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEntityTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetEntityTagsResponseBody
	GetData() *bool
	SetErrorCode(v string) *SetEntityTagsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SetEntityTagsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *SetEntityTagsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SetEntityTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetEntityTagsResponseBody
	GetSuccess() *bool
}

type SetEntityTagsResponseBody struct {
	// Indicates whether the call was successful. Valid values:
	//
	// true and false.
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
	// true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetEntityTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetEntityTagsResponseBody) GoString() string {
	return s.String()
}

func (s *SetEntityTagsResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetEntityTagsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SetEntityTagsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SetEntityTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SetEntityTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetEntityTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetEntityTagsResponseBody) SetData(v bool) *SetEntityTagsResponseBody {
	s.Data = &v
	return s
}

func (s *SetEntityTagsResponseBody) SetErrorCode(v string) *SetEntityTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetEntityTagsResponseBody) SetErrorMessage(v string) *SetEntityTagsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetEntityTagsResponseBody) SetHttpStatusCode(v int32) *SetEntityTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SetEntityTagsResponseBody) SetRequestId(v string) *SetEntityTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetEntityTagsResponseBody) SetSuccess(v bool) *SetEntityTagsResponseBody {
	s.Success = &v
	return s
}

func (s *SetEntityTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
