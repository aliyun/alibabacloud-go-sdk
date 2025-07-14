// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneGroupTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSwimmingLaneGroupTagsResponseBody
	GetCode() *string
	SetData(v []*ListSwimmingLaneGroupTagsResponseBodyData) *ListSwimmingLaneGroupTagsResponseBody
	GetData() []*ListSwimmingLaneGroupTagsResponseBodyData
	SetErrorCode(v string) *ListSwimmingLaneGroupTagsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListSwimmingLaneGroupTagsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSwimmingLaneGroupTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSwimmingLaneGroupTagsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListSwimmingLaneGroupTagsResponseBody
	GetTraceId() *string
}

type ListSwimmingLaneGroupTagsResponseBody struct {
	// example:
	//
	// 200
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListSwimmingLaneGroupTagsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 30375C38-F4ED-4135-A0AE-5C75DC7F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// ac1a0b2215622920113732501e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListSwimmingLaneGroupTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupTagsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSwimmingLaneGroupTagsResponseBody) GetData() []*ListSwimmingLaneGroupTagsResponseBodyData {
	return s.Data
}

func (s *ListSwimmingLaneGroupTagsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSwimmingLaneGroupTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSwimmingLaneGroupTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSwimmingLaneGroupTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSwimmingLaneGroupTagsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListSwimmingLaneGroupTagsResponseBody) SetCode(v string) *ListSwimmingLaneGroupTagsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponseBody) SetData(v []*ListSwimmingLaneGroupTagsResponseBodyData) *ListSwimmingLaneGroupTagsResponseBody {
	s.Data = v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponseBody) SetErrorCode(v string) *ListSwimmingLaneGroupTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponseBody) SetMessage(v string) *ListSwimmingLaneGroupTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponseBody) SetRequestId(v string) *ListSwimmingLaneGroupTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponseBody) SetSuccess(v bool) *ListSwimmingLaneGroupTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponseBody) SetTraceId(v string) *ListSwimmingLaneGroupTagsResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSwimmingLaneGroupTagsResponseBodyData struct {
	// example:
	//
	// {"version":"1.0.0","owner":"team-a"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// {"alicloud.service.tag":"g1"}
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListSwimmingLaneGroupTagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupTagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupTagsResponseBodyData) GetMetadata() *string {
	return s.Metadata
}

func (s *ListSwimmingLaneGroupTagsResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *ListSwimmingLaneGroupTagsResponseBodyData) SetMetadata(v string) *ListSwimmingLaneGroupTagsResponseBodyData {
	s.Metadata = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponseBodyData) SetTag(v string) *ListSwimmingLaneGroupTagsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
