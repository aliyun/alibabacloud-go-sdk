// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTraceGetResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsTraceGetResultRequest
	GetInstanceId() *string
	SetQueryId(v string) *OnsTraceGetResultRequest
	GetQueryId() *string
	SetTopic(v string) *OnsTraceGetResultRequest
	GetTopic() *string
}

type OnsTraceGetResultRequest struct {
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the task that was created to query the trace of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 272967562652883649157096685****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The topic to which the message belongs.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTraceGetResultRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceGetResultRequest) GoString() string {
	return s.String()
}

func (s *OnsTraceGetResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTraceGetResultRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *OnsTraceGetResultRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTraceGetResultRequest) SetInstanceId(v string) *OnsTraceGetResultRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTraceGetResultRequest) SetQueryId(v string) *OnsTraceGetResultRequest {
	s.QueryId = &v
	return s
}

func (s *OnsTraceGetResultRequest) SetTopic(v string) *OnsTraceGetResultRequest {
	s.Topic = &v
	return s
}

func (s *OnsTraceGetResultRequest) Validate() error {
	return dara.Validate(s)
}
