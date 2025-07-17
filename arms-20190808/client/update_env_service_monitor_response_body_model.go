// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvServiceMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateEnvServiceMonitorResponseBody
	GetCode() *int32
	SetData(v *UpdateEnvServiceMonitorResponseBodyData) *UpdateEnvServiceMonitorResponseBody
	GetData() *UpdateEnvServiceMonitorResponseBodyData
	SetMessage(v string) *UpdateEnvServiceMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEnvServiceMonitorResponseBody
	GetRequestId() *string
}

type UpdateEnvServiceMonitorResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *UpdateEnvServiceMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A474FF8-7861-4D00-81B5-5BC3DA4E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnvServiceMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvServiceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvServiceMonitorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateEnvServiceMonitorResponseBody) GetData() *UpdateEnvServiceMonitorResponseBodyData {
	return s.Data
}

func (s *UpdateEnvServiceMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEnvServiceMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnvServiceMonitorResponseBody) SetCode(v int32) *UpdateEnvServiceMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvServiceMonitorResponseBody) SetData(v *UpdateEnvServiceMonitorResponseBodyData) *UpdateEnvServiceMonitorResponseBody {
	s.Data = v
	return s
}

func (s *UpdateEnvServiceMonitorResponseBody) SetMessage(v string) *UpdateEnvServiceMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEnvServiceMonitorResponseBody) SetRequestId(v string) *UpdateEnvServiceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnvServiceMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEnvServiceMonitorResponseBodyData struct {
	// Indicates whether targets are matched.
	//
	// example:
	//
	// Match successful.
	MatchedMsg *string `json:"MatchedMsg,omitempty" xml:"MatchedMsg,omitempty"`
	// The number of matched targets.
	//
	// example:
	//
	// 1
	MatchedTargetCount *string `json:"MatchedTargetCount,omitempty" xml:"MatchedTargetCount,omitempty"`
}

func (s UpdateEnvServiceMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvServiceMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateEnvServiceMonitorResponseBodyData) GetMatchedMsg() *string {
	return s.MatchedMsg
}

func (s *UpdateEnvServiceMonitorResponseBodyData) GetMatchedTargetCount() *string {
	return s.MatchedTargetCount
}

func (s *UpdateEnvServiceMonitorResponseBodyData) SetMatchedMsg(v string) *UpdateEnvServiceMonitorResponseBodyData {
	s.MatchedMsg = &v
	return s
}

func (s *UpdateEnvServiceMonitorResponseBodyData) SetMatchedTargetCount(v string) *UpdateEnvServiceMonitorResponseBodyData {
	s.MatchedTargetCount = &v
	return s
}

func (s *UpdateEnvServiceMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
