// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvPodMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateEnvPodMonitorResponseBody
	GetCode() *int32
	SetData(v *UpdateEnvPodMonitorResponseBodyData) *UpdateEnvPodMonitorResponseBody
	GetData() *UpdateEnvPodMonitorResponseBodyData
	SetMessage(v string) *UpdateEnvPodMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEnvPodMonitorResponseBody
	GetRequestId() *string
}

type UpdateEnvPodMonitorResponseBody struct {
	// The response code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *UpdateEnvPodMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnvPodMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvPodMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvPodMonitorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateEnvPodMonitorResponseBody) GetData() *UpdateEnvPodMonitorResponseBodyData {
	return s.Data
}

func (s *UpdateEnvPodMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEnvPodMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnvPodMonitorResponseBody) SetCode(v int32) *UpdateEnvPodMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvPodMonitorResponseBody) SetData(v *UpdateEnvPodMonitorResponseBodyData) *UpdateEnvPodMonitorResponseBody {
	s.Data = v
	return s
}

func (s *UpdateEnvPodMonitorResponseBody) SetMessage(v string) *UpdateEnvPodMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEnvPodMonitorResponseBody) SetRequestId(v string) *UpdateEnvPodMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnvPodMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEnvPodMonitorResponseBodyData struct {
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

func (s UpdateEnvPodMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvPodMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateEnvPodMonitorResponseBodyData) GetMatchedMsg() *string {
	return s.MatchedMsg
}

func (s *UpdateEnvPodMonitorResponseBodyData) GetMatchedTargetCount() *string {
	return s.MatchedTargetCount
}

func (s *UpdateEnvPodMonitorResponseBodyData) SetMatchedMsg(v string) *UpdateEnvPodMonitorResponseBodyData {
	s.MatchedMsg = &v
	return s
}

func (s *UpdateEnvPodMonitorResponseBodyData) SetMatchedTargetCount(v string) *UpdateEnvPodMonitorResponseBodyData {
	s.MatchedTargetCount = &v
	return s
}

func (s *UpdateEnvPodMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
