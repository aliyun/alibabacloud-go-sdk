// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagsBySwimmingLaneGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *GetTagsBySwimmingLaneGroupIdResponseBody
	GetData() []*string
	SetErrorCode(v string) *GetTagsBySwimmingLaneGroupIdResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTagsBySwimmingLaneGroupIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTagsBySwimmingLaneGroupIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTagsBySwimmingLaneGroupIdResponseBody
	GetSuccess() *bool
}

type GetTagsBySwimmingLaneGroupIdResponseBody struct {
	// The data of the tag.
	//
	// example:
	//
	// ["gray"]
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTagsBySwimmingLaneGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTagsBySwimmingLaneGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) SetData(v []*string) *GetTagsBySwimmingLaneGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) SetErrorCode(v string) *GetTagsBySwimmingLaneGroupIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) SetMessage(v string) *GetTagsBySwimmingLaneGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) SetRequestId(v string) *GetTagsBySwimmingLaneGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) SetSuccess(v bool) *GetTagsBySwimmingLaneGroupIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetTagsBySwimmingLaneGroupIdResponseBody) Validate() error {
	return dara.Validate(s)
}
