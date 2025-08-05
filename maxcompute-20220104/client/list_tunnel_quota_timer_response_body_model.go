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
	// The data returned.
	Data []*ListTunnelQuotaTimerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0b716671885050924814e3623
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	return dara.Validate(s)
}

type ListTunnelQuotaTimerResponseBodyData struct {
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
	// The time zone property for the time-specific configuration.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The parameters for the time-specific configuration.
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
	return dara.Validate(s)
}

type ListTunnelQuotaTimerResponseBodyDataTunnelQuotaParameter struct {
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
