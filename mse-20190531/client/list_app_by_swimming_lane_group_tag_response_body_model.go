// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppBySwimmingLaneGroupTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListAppBySwimmingLaneGroupTagResponseBody
	GetData() interface{}
	SetErrorCode(v string) *ListAppBySwimmingLaneGroupTagResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListAppBySwimmingLaneGroupTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAppBySwimmingLaneGroupTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAppBySwimmingLaneGroupTagResponseBody
	GetSuccess() *bool
}

type ListAppBySwimmingLaneGroupTagResponseBody struct {
	// The details of the data.
	//
	// example:
	//
	// [{appName:"test",appId:"hkhon1po62@958bba95910341f
	//
	// "}]
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 7466566F-F30F-4A29-965D-3E0AF21D****
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

func (s ListAppBySwimmingLaneGroupTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppBySwimmingLaneGroupTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) SetData(v interface{}) *ListAppBySwimmingLaneGroupTagResponseBody {
	s.Data = v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) SetErrorCode(v string) *ListAppBySwimmingLaneGroupTagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) SetMessage(v string) *ListAppBySwimmingLaneGroupTagResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) SetRequestId(v string) *ListAppBySwimmingLaneGroupTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) SetSuccess(v bool) *ListAppBySwimmingLaneGroupTagResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagResponseBody) Validate() error {
	return dara.Validate(s)
}
