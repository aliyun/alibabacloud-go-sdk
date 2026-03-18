// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTunnelQuotaTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListTunnelQuotaTimerResponseBodyData) *ListTunnelQuotaTimerResponseBody
	GetData() []*ListTunnelQuotaTimerResponseBodyData
	SetErrorCode(v string) *ListTunnelQuotaTimerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListTunnelQuotaTimerResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListTunnelQuotaTimerResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListTunnelQuotaTimerResponseBody
	GetRequestId() *string
}

type ListTunnelQuotaTimerResponseBody struct {
	Data      []*ListTunnelQuotaTimerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorCode *string                                 `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                 `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                                  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListTunnelQuotaTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTunnelQuotaTimerResponseBody) GoString() string {
	return s.String()
}

func (s *ListTunnelQuotaTimerResponseBody) GetData() []*ListTunnelQuotaTimerResponseBodyData {
	return s.Data
}

func (s *ListTunnelQuotaTimerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTunnelQuotaTimerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListTunnelQuotaTimerResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListTunnelQuotaTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTunnelQuotaTimerResponseBody) SetData(v []*ListTunnelQuotaTimerResponseBodyData) *ListTunnelQuotaTimerResponseBody {
	s.Data = v
	return s
}

func (s *ListTunnelQuotaTimerResponseBody) SetErrorCode(v string) *ListTunnelQuotaTimerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBody) SetErrorMsg(v string) *ListTunnelQuotaTimerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBody) SetHttpCode(v int32) *ListTunnelQuotaTimerResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBody) SetRequestId(v string) *ListTunnelQuotaTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTunnelQuotaTimerResponseBodyData struct {
	BeginTime            *string                                                   `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	EndTime              *string                                                   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Timezone             *string                                                   `json:"timezone,omitempty" xml:"timezone,omitempty"`
	TunnelQuotaParameter *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter `json:"tunnelQuotaParameter,omitempty" xml:"tunnelQuotaParameter,omitempty" type:"Struct"`
}

func (s ListTunnelQuotaTimerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTunnelQuotaTimerResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTunnelQuotaTimerResponseBodyData) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ListTunnelQuotaTimerResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTunnelQuotaTimerResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *ListTunnelQuotaTimerResponseBodyData) GetTunnelQuotaParameter() *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter {
	return s.TunnelQuotaParameter
}

func (s *ListTunnelQuotaTimerResponseBodyData) SetBeginTime(v string) *ListTunnelQuotaTimerResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyData) SetEndTime(v string) *ListTunnelQuotaTimerResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyData) SetTimezone(v string) *ListTunnelQuotaTimerResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyData) SetTunnelQuotaParameter(v *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) *ListTunnelQuotaTimerResponseBodyData {
	s.TunnelQuotaParameter = v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyData) Validate() error {
	if s.TunnelQuotaParameter != nil {
		if err := s.TunnelQuotaParameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter struct {
	ElasticReservedSlotNum *int64 `json:"elasticReservedSlotNum,omitempty" xml:"elasticReservedSlotNum,omitempty"`
	SlotNum                *int64 `json:"slotNum,omitempty" xml:"slotNum,omitempty"`
}

func (s ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) String() string {
	return dara.Prettify(s)
}

func (s ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) GoString() string {
	return s.String()
}

func (s *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) GetElasticReservedSlotNum() *int64 {
	return s.ElasticReservedSlotNum
}

func (s *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) GetSlotNum() *int64 {
	return s.SlotNum
}

func (s *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) SetElasticReservedSlotNum(v int64) *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter {
	s.ElasticReservedSlotNum = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) SetSlotNum(v int64) *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter {
	s.SlotNum = &v
	return s
}

func (s *ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter) Validate() error {
	return dara.Validate(s)
}
