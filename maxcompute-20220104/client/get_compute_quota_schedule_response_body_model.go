// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeQuotaScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetComputeQuotaScheduleResponseBodyData) *GetComputeQuotaScheduleResponseBody
	GetData() []*GetComputeQuotaScheduleResponseBodyData
	SetErrorCode(v string) *GetComputeQuotaScheduleResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetComputeQuotaScheduleResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetComputeQuotaScheduleResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetComputeQuotaScheduleResponseBody
	GetRequestId() *string
}

type GetComputeQuotaScheduleResponseBody struct {
	Data      []*GetComputeQuotaScheduleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorCode *string                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                    `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                                     `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetComputeQuotaScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleResponseBody) GetData() []*GetComputeQuotaScheduleResponseBodyData {
	return s.Data
}

func (s *GetComputeQuotaScheduleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetComputeQuotaScheduleResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetComputeQuotaScheduleResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetComputeQuotaScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeQuotaScheduleResponseBody) SetData(v []*GetComputeQuotaScheduleResponseBodyData) *GetComputeQuotaScheduleResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeQuotaScheduleResponseBody) SetErrorCode(v string) *GetComputeQuotaScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBody) SetErrorMsg(v string) *GetComputeQuotaScheduleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBody) SetHttpCode(v int32) *GetComputeQuotaScheduleResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBody) SetRequestId(v string) *GetComputeQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBody) Validate() error {
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

type GetComputeQuotaScheduleResponseBodyData struct {
	Condition *GetComputeQuotaScheduleResponseBodyDataCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
	Id        *string                                           `json:"id,omitempty" xml:"id,omitempty"`
	Plan      *string                                           `json:"plan,omitempty" xml:"plan,omitempty"`
	Timezone  *string                                           `json:"timezone,omitempty" xml:"timezone,omitempty"`
	Type      *string                                           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetComputeQuotaScheduleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleResponseBodyData) GetCondition() *GetComputeQuotaScheduleResponseBodyDataCondition {
	return s.Condition
}

func (s *GetComputeQuotaScheduleResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetComputeQuotaScheduleResponseBodyData) GetPlan() *string {
	return s.Plan
}

func (s *GetComputeQuotaScheduleResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *GetComputeQuotaScheduleResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetCondition(v *GetComputeQuotaScheduleResponseBodyDataCondition) *GetComputeQuotaScheduleResponseBodyData {
	s.Condition = v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetId(v string) *GetComputeQuotaScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetPlan(v string) *GetComputeQuotaScheduleResponseBodyData {
	s.Plan = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetTimezone(v string) *GetComputeQuotaScheduleResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyData) SetType(v string) *GetComputeQuotaScheduleResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyData) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeQuotaScheduleResponseBodyDataCondition struct {
	At *string `json:"at,omitempty" xml:"at,omitempty"`
}

func (s GetComputeQuotaScheduleResponseBodyDataCondition) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaScheduleResponseBodyDataCondition) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleResponseBodyDataCondition) GetAt() *string {
	return s.At
}

func (s *GetComputeQuotaScheduleResponseBodyDataCondition) SetAt(v string) *GetComputeQuotaScheduleResponseBodyDataCondition {
	s.At = &v
	return s
}

func (s *GetComputeQuotaScheduleResponseBodyDataCondition) Validate() error {
	return dara.Validate(s)
}
