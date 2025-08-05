// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTunnelQuotaTimerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*UpdateTunnelQuotaTimerRequestBody) *UpdateTunnelQuotaTimerRequest
	GetBody() []*UpdateTunnelQuotaTimerRequestBody
	SetTimezone(v string) *UpdateTunnelQuotaTimerRequest
	GetTimezone() *string
}

type UpdateTunnelQuotaTimerRequest struct {
	// The request body.
	Body     []*UpdateTunnelQuotaTimerRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	Timezone *string                              `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s UpdateTunnelQuotaTimerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTunnelQuotaTimerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerRequest) GetBody() []*UpdateTunnelQuotaTimerRequestBody {
	return s.Body
}

func (s *UpdateTunnelQuotaTimerRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateTunnelQuotaTimerRequest) SetBody(v []*UpdateTunnelQuotaTimerRequestBody) *UpdateTunnelQuotaTimerRequest {
	s.Body = v
	return s
}

func (s *UpdateTunnelQuotaTimerRequest) SetTimezone(v string) *UpdateTunnelQuotaTimerRequest {
	s.Timezone = &v
	return s
}

func (s *UpdateTunnelQuotaTimerRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTunnelQuotaTimerRequestBody struct {
	// The start time of the time-specific configuration.
	//
	// example:
	//
	// 00:00
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The end time of the time-specific configuration.
	//
	// example:
	//
	// 08:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The parameters for the time-specific configuration.
	TunnelQuotaParameter *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter `json:"tunnelQuotaParameter,omitempty" xml:"tunnelQuotaParameter,omitempty" type:"Struct"`
}

func (s UpdateTunnelQuotaTimerRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTunnelQuotaTimerRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerRequestBody) GetBeginTime() *string {
	return s.BeginTime
}

func (s *UpdateTunnelQuotaTimerRequestBody) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateTunnelQuotaTimerRequestBody) GetTunnelQuotaParameter() *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter {
	return s.TunnelQuotaParameter
}

func (s *UpdateTunnelQuotaTimerRequestBody) SetBeginTime(v string) *UpdateTunnelQuotaTimerRequestBody {
	s.BeginTime = &v
	return s
}

func (s *UpdateTunnelQuotaTimerRequestBody) SetEndTime(v string) *UpdateTunnelQuotaTimerRequestBody {
	s.EndTime = &v
	return s
}

func (s *UpdateTunnelQuotaTimerRequestBody) SetTunnelQuotaParameter(v *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) *UpdateTunnelQuotaTimerRequestBody {
	s.TunnelQuotaParameter = v
	return s
}

func (s *UpdateTunnelQuotaTimerRequestBody) Validate() error {
	return dara.Validate(s)
}

type UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter struct {
	// The number of elastically reserved slots.
	//
	// example:
	//
	// 100
	ElasticReservedSlotNum *int64 `json:"elasticReservedSlotNum,omitempty" xml:"elasticReservedSlotNum,omitempty"`
	// The number of reserved slots.
	//
	// example:
	//
	// 100
	SlotNum *int64 `json:"slotNum,omitempty" xml:"slotNum,omitempty"`
}

func (s UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) String() string {
	return dara.Prettify(s)
}

func (s UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) GoString() string {
	return s.String()
}

func (s *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) GetElasticReservedSlotNum() *int64 {
	return s.ElasticReservedSlotNum
}

func (s *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) GetSlotNum() *int64 {
	return s.SlotNum
}

func (s *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) SetElasticReservedSlotNum(v int64) *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter {
	s.ElasticReservedSlotNum = &v
	return s
}

func (s *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) SetSlotNum(v int64) *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter {
	s.SlotNum = &v
	return s
}

func (s *UpdateTunnelQuotaTimerRequestBodyTunnelQuotaParameter) Validate() error {
	return dara.Validate(s)
}
