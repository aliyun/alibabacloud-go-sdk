// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvPodMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateEnvPodMonitorResponseBody
	GetCode() *int32
	SetData(v *CreateEnvPodMonitorResponseBodyData) *CreateEnvPodMonitorResponseBody
	GetData() *CreateEnvPodMonitorResponseBodyData
	SetMessage(v string) *CreateEnvPodMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEnvPodMonitorResponseBody
	GetRequestId() *string
}

type CreateEnvPodMonitorResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result, which indicates whether the operation was successful.
	Data *CreateEnvPodMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnvPodMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvPodMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvPodMonitorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateEnvPodMonitorResponseBody) GetData() *CreateEnvPodMonitorResponseBodyData {
	return s.Data
}

func (s *CreateEnvPodMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEnvPodMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnvPodMonitorResponseBody) SetCode(v int32) *CreateEnvPodMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnvPodMonitorResponseBody) SetData(v *CreateEnvPodMonitorResponseBodyData) *CreateEnvPodMonitorResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnvPodMonitorResponseBody) SetMessage(v string) *CreateEnvPodMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEnvPodMonitorResponseBody) SetRequestId(v string) *CreateEnvPodMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnvPodMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateEnvPodMonitorResponseBodyData struct {
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
	// The namespace.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the created PodMonitor.
	//
	// example:
	//
	// arms-admin-pm1
	PodMonitorName *string `json:"PodMonitorName,omitempty" xml:"PodMonitorName,omitempty"`
}

func (s CreateEnvPodMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvPodMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnvPodMonitorResponseBodyData) GetMatchedMsg() *string {
	return s.MatchedMsg
}

func (s *CreateEnvPodMonitorResponseBodyData) GetMatchedTargetCount() *string {
	return s.MatchedTargetCount
}

func (s *CreateEnvPodMonitorResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateEnvPodMonitorResponseBodyData) GetPodMonitorName() *string {
	return s.PodMonitorName
}

func (s *CreateEnvPodMonitorResponseBodyData) SetMatchedMsg(v string) *CreateEnvPodMonitorResponseBodyData {
	s.MatchedMsg = &v
	return s
}

func (s *CreateEnvPodMonitorResponseBodyData) SetMatchedTargetCount(v string) *CreateEnvPodMonitorResponseBodyData {
	s.MatchedTargetCount = &v
	return s
}

func (s *CreateEnvPodMonitorResponseBodyData) SetNamespace(v string) *CreateEnvPodMonitorResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateEnvPodMonitorResponseBodyData) SetPodMonitorName(v string) *CreateEnvPodMonitorResponseBodyData {
	s.PodMonitorName = &v
	return s
}

func (s *CreateEnvPodMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
