// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTopicByTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomTopicByTopicResponseBody
	GetCode() *string
	SetData(v int32) *DeleteCustomTopicByTopicResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *DeleteCustomTopicByTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCustomTopicByTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomTopicByTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomTopicByTopicResponseBody
	GetSuccess() *bool
}

type DeleteCustomTopicByTopicResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 33
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

func (s DeleteCustomTopicByTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTopicByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicByTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomTopicByTopicResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteCustomTopicByTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCustomTopicByTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomTopicByTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomTopicByTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomTopicByTopicResponseBody) SetCode(v string) *DeleteCustomTopicByTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetData(v int32) *DeleteCustomTopicByTopicResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetHttpStatusCode(v int32) *DeleteCustomTopicByTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetMessage(v string) *DeleteCustomTopicByTopicResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetRequestId(v string) *DeleteCustomTopicByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetSuccess(v bool) *DeleteCustomTopicByTopicResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
