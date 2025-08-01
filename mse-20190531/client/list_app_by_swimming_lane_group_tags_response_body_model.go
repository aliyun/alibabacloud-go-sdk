// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppBySwimmingLaneGroupTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string][]*DataValue) *ListAppBySwimmingLaneGroupTagsResponseBody
	GetData() map[string][]*DataValue
	SetErrorCode(v string) *ListAppBySwimmingLaneGroupTagsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListAppBySwimmingLaneGroupTagsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAppBySwimmingLaneGroupTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAppBySwimmingLaneGroupTagsResponseBody
	GetSuccess() *bool
}

type ListAppBySwimmingLaneGroupTagsResponseBody struct {
	// The returned data.
	Data map[string][]*DataValue `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional request information.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6832e76b-bb5f-4dea-****-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppBySwimmingLaneGroupTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppBySwimmingLaneGroupTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) GetData() map[string][]*DataValue {
	return s.Data
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) SetData(v map[string][]*DataValue) *ListAppBySwimmingLaneGroupTagsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) SetErrorCode(v string) *ListAppBySwimmingLaneGroupTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) SetMessage(v string) *ListAppBySwimmingLaneGroupTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) SetRequestId(v string) *ListAppBySwimmingLaneGroupTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) SetSuccess(v bool) *ListAppBySwimmingLaneGroupTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
