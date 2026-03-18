// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetQuotaScheduleResponseBodyData) *GetQuotaScheduleResponseBody
	GetData() []*GetQuotaScheduleResponseBodyData
	SetErrorCode(v string) *GetQuotaScheduleResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetQuotaScheduleResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetQuotaScheduleResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetQuotaScheduleResponseBody
	GetRequestId() *string
}

type GetQuotaScheduleResponseBody struct {
	Data      []*GetQuotaScheduleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorCode *string                             `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                             `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                              `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetQuotaScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponseBody) GetData() []*GetQuotaScheduleResponseBodyData {
	return s.Data
}

func (s *GetQuotaScheduleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetQuotaScheduleResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetQuotaScheduleResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetQuotaScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaScheduleResponseBody) SetData(v []*GetQuotaScheduleResponseBodyData) *GetQuotaScheduleResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetErrorCode(v string) *GetQuotaScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetErrorMsg(v string) *GetQuotaScheduleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetHttpCode(v int32) *GetQuotaScheduleResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetRequestId(v string) *GetQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) Validate() error {
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

type GetQuotaScheduleResponseBodyData struct {
	Condition *GetQuotaScheduleResponseBodyDataCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
	Id        *string                                    `json:"id,omitempty" xml:"id,omitempty"`
	Operator  *string                                    `json:"operator,omitempty" xml:"operator,omitempty"`
	Plan      *string                                    `json:"plan,omitempty" xml:"plan,omitempty"`
	Timezone  *string                                    `json:"timezone,omitempty" xml:"timezone,omitempty"`
	Type      *string                                    `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetQuotaScheduleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponseBodyData) GetCondition() *GetQuotaScheduleResponseBodyDataCondition {
	return s.Condition
}

func (s *GetQuotaScheduleResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetQuotaScheduleResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *GetQuotaScheduleResponseBodyData) GetPlan() *string {
	return s.Plan
}

func (s *GetQuotaScheduleResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *GetQuotaScheduleResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetQuotaScheduleResponseBodyData) SetCondition(v *GetQuotaScheduleResponseBodyDataCondition) *GetQuotaScheduleResponseBodyData {
	s.Condition = v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetId(v string) *GetQuotaScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetOperator(v string) *GetQuotaScheduleResponseBodyData {
	s.Operator = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetPlan(v string) *GetQuotaScheduleResponseBodyData {
	s.Plan = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetTimezone(v string) *GetQuotaScheduleResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetType(v string) *GetQuotaScheduleResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQuotaScheduleResponseBodyDataCondition struct {
	After *string `json:"after,omitempty" xml:"after,omitempty"`
	At    *string `json:"at,omitempty" xml:"at,omitempty"`
}

func (s GetQuotaScheduleResponseBodyDataCondition) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaScheduleResponseBodyDataCondition) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponseBodyDataCondition) GetAfter() *string {
	return s.After
}

func (s *GetQuotaScheduleResponseBodyDataCondition) GetAt() *string {
	return s.At
}

func (s *GetQuotaScheduleResponseBodyDataCondition) SetAfter(v string) *GetQuotaScheduleResponseBodyDataCondition {
	s.After = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyDataCondition) SetAt(v string) *GetQuotaScheduleResponseBodyDataCondition {
	s.At = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyDataCondition) Validate() error {
	return dara.Validate(s)
}
