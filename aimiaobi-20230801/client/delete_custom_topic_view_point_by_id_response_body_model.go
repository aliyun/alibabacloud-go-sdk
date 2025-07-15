// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTopicViewPointByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomTopicViewPointByIdResponseBody
	GetCode() *string
	SetData(v int32) *DeleteCustomTopicViewPointByIdResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *DeleteCustomTopicViewPointByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCustomTopicViewPointByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomTopicViewPointByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomTopicViewPointByIdResponseBody
	GetSuccess() *bool
}

type DeleteCustomTopicViewPointByIdResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 7
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomTopicViewPointByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTopicViewPointByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetCode(v string) *DeleteCustomTopicViewPointByIdResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetData(v int32) *DeleteCustomTopicViewPointByIdResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetHttpStatusCode(v int32) *DeleteCustomTopicViewPointByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetMessage(v string) *DeleteCustomTopicViewPointByIdResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetRequestId(v string) *DeleteCustomTopicViewPointByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetSuccess(v bool) *DeleteCustomTopicViewPointByIdResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
