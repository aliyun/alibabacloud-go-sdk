// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvServiceMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateEnvServiceMonitorResponseBody
	GetCode() *int32
	SetData(v *CreateEnvServiceMonitorResponseBodyData) *CreateEnvServiceMonitorResponseBody
	GetData() *CreateEnvServiceMonitorResponseBodyData
	SetMessage(v string) *CreateEnvServiceMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEnvServiceMonitorResponseBody
	GetRequestId() *string
}

type CreateEnvServiceMonitorResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *CreateEnvServiceMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateEnvServiceMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvServiceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvServiceMonitorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateEnvServiceMonitorResponseBody) GetData() *CreateEnvServiceMonitorResponseBodyData {
	return s.Data
}

func (s *CreateEnvServiceMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEnvServiceMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnvServiceMonitorResponseBody) SetCode(v int32) *CreateEnvServiceMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnvServiceMonitorResponseBody) SetData(v *CreateEnvServiceMonitorResponseBodyData) *CreateEnvServiceMonitorResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnvServiceMonitorResponseBody) SetMessage(v string) *CreateEnvServiceMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEnvServiceMonitorResponseBody) SetRequestId(v string) *CreateEnvServiceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnvServiceMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateEnvServiceMonitorResponseBodyData struct {
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
	MatchedTargetCount *int32 `json:"MatchedTargetCount,omitempty" xml:"MatchedTargetCount,omitempty"`
	// The namespace.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the created ServiceMonitor.
	//
	// example:
	//
	// arms-admin1
	ServiceMonitorName *string `json:"ServiceMonitorName,omitempty" xml:"ServiceMonitorName,omitempty"`
}

func (s CreateEnvServiceMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvServiceMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnvServiceMonitorResponseBodyData) GetMatchedMsg() *string {
	return s.MatchedMsg
}

func (s *CreateEnvServiceMonitorResponseBodyData) GetMatchedTargetCount() *int32 {
	return s.MatchedTargetCount
}

func (s *CreateEnvServiceMonitorResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateEnvServiceMonitorResponseBodyData) GetServiceMonitorName() *string {
	return s.ServiceMonitorName
}

func (s *CreateEnvServiceMonitorResponseBodyData) SetMatchedMsg(v string) *CreateEnvServiceMonitorResponseBodyData {
	s.MatchedMsg = &v
	return s
}

func (s *CreateEnvServiceMonitorResponseBodyData) SetMatchedTargetCount(v int32) *CreateEnvServiceMonitorResponseBodyData {
	s.MatchedTargetCount = &v
	return s
}

func (s *CreateEnvServiceMonitorResponseBodyData) SetNamespace(v string) *CreateEnvServiceMonitorResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateEnvServiceMonitorResponseBodyData) SetServiceMonitorName(v string) *CreateEnvServiceMonitorResponseBodyData {
	s.ServiceMonitorName = &v
	return s
}

func (s *CreateEnvServiceMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
