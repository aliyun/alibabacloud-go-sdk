// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendTraceByConnectionAndChannelAndDeliveryTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
	GetChannelId() *string
	SetClientToken(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
	GetClientToken() *string
	SetConnectionId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
	GetConnectionId() *string
	SetConsoleSessionId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
	GetConsoleSessionId() *string
	SetDeliveryTag(v int64) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
	GetDeliveryTag() *int64
	SetEndTime(v int32) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
	GetEndTime() *int32
	SetInstanceId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
	GetInstanceId() *string
	SetStartTime(v int32) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
	GetStartTime() *int32
	SetVhostName(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest
	GetVhostName() *string
}

type GetSendTraceByConnectionAndChannelAndDeliveryTagRequest struct {
	// This parameter is required.
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	DeliveryTag *int64 `json:"DeliveryTag,omitempty" xml:"DeliveryTag,omitempty"`
	// This parameter is required.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GoString() string {
	return s.String()
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GetDeliveryTag() *int64 {
	return s.DeliveryTag
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GetEndTime() *int32 {
	return s.EndTime
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) SetChannelId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest {
	s.ChannelId = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) SetClientToken(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest {
	s.ClientToken = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) SetConnectionId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest {
	s.ConnectionId = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) SetConsoleSessionId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) SetDeliveryTag(v int64) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest {
	s.DeliveryTag = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) SetEndTime(v int32) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest {
	s.EndTime = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) SetInstanceId(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) SetStartTime(v int32) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest {
	s.StartTime = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) SetVhostName(v string) *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest {
	s.VhostName = &v
	return s
}

func (s *GetSendTraceByConnectionAndChannelAndDeliveryTagRequest) Validate() error {
	return dara.Validate(s)
}
