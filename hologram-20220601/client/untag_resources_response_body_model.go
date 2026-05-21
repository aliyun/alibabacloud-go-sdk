// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UntagResourcesResponseBody
	GetData() *bool
	SetErrorCode(v string) *UntagResourcesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UntagResourcesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UntagResourcesResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UntagResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UntagResourcesResponseBody
	GetSuccess() *bool
}

type UntagResourcesResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CE1C813B-3445-57EC-9504-342E45D1C7EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) GetData() *bool {
	return s.Data
}

func (s *UntagResourcesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UntagResourcesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UntagResourcesResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UntagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UntagResourcesResponseBody) SetData(v bool) *UntagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *UntagResourcesResponseBody) SetErrorCode(v string) *UntagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetErrorMessage(v string) *UntagResourcesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UntagResourcesResponseBody) SetHttpStatusCode(v string) *UntagResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
