// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTpsByTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v string) *GetTpsByTimeRequest
	GetApi() *string
	SetClientToken(v string) *GetTpsByTimeRequest
	GetClientToken() *string
	SetConsoleSessionId(v string) *GetTpsByTimeRequest
	GetConsoleSessionId() *string
	SetEndTime(v int32) *GetTpsByTimeRequest
	GetEndTime() *int32
	SetInstanceId(v string) *GetTpsByTimeRequest
	GetInstanceId() *string
	SetQueueName(v string) *GetTpsByTimeRequest
	GetQueueName() *string
	SetStartTime(v int32) *GetTpsByTimeRequest
	GetStartTime() *int32
	SetVhostName(v string) *GetTpsByTimeRequest
	GetVhostName() *string
}

type GetTpsByTimeRequest struct {
	Api              *string `json:"Api,omitempty" xml:"Api,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName  *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetTpsByTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTpsByTimeRequest) GoString() string {
	return s.String()
}

func (s *GetTpsByTimeRequest) GetApi() *string {
	return s.Api
}

func (s *GetTpsByTimeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetTpsByTimeRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetTpsByTimeRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *GetTpsByTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTpsByTimeRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *GetTpsByTimeRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *GetTpsByTimeRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetTpsByTimeRequest) SetApi(v string) *GetTpsByTimeRequest {
	s.Api = &v
	return s
}

func (s *GetTpsByTimeRequest) SetClientToken(v string) *GetTpsByTimeRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTpsByTimeRequest) SetConsoleSessionId(v string) *GetTpsByTimeRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetTpsByTimeRequest) SetEndTime(v int32) *GetTpsByTimeRequest {
	s.EndTime = &v
	return s
}

func (s *GetTpsByTimeRequest) SetInstanceId(v string) *GetTpsByTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTpsByTimeRequest) SetQueueName(v string) *GetTpsByTimeRequest {
	s.QueueName = &v
	return s
}

func (s *GetTpsByTimeRequest) SetStartTime(v int32) *GetTpsByTimeRequest {
	s.StartTime = &v
	return s
}

func (s *GetTpsByTimeRequest) SetVhostName(v string) *GetTpsByTimeRequest {
	s.VhostName = &v
	return s
}

func (s *GetTpsByTimeRequest) Validate() error {
	return dara.Validate(s)
}
