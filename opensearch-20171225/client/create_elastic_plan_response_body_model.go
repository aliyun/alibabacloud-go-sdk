// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateElasticPlanResponseBody
	GetCode() *string
	SetHttpCode(v int64) *CreateElasticPlanResponseBody
	GetHttpCode() *int64
	SetMessage(v string) *CreateElasticPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateElasticPlanResponseBody
	GetRequestId() *string
	SetResult(v *CreateElasticPlanResponseBodyResult) *CreateElasticPlanResponseBody
	GetResult() *CreateElasticPlanResponseBodyResult
}

type CreateElasticPlanResponseBody struct {
	// example:
	//
	// ElasticPlan.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// Elastic plan not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0FFF39C5-ED93-5234-806D-0824B967E6A3
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateElasticPlanResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateElasticPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateElasticPlanResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *CreateElasticPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateElasticPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateElasticPlanResponseBody) GetResult() *CreateElasticPlanResponseBodyResult {
	return s.Result
}

func (s *CreateElasticPlanResponseBody) SetCode(v string) *CreateElasticPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateElasticPlanResponseBody) SetHttpCode(v int64) *CreateElasticPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateElasticPlanResponseBody) SetMessage(v string) *CreateElasticPlanResponseBody {
	s.Message = &v
	return s
}

func (s *CreateElasticPlanResponseBody) SetRequestId(v string) *CreateElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateElasticPlanResponseBody) SetResult(v *CreateElasticPlanResponseBodyResult) *CreateElasticPlanResponseBody {
	s.Result = v
	return s
}

func (s *CreateElasticPlanResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateElasticPlanResponseBodyResult struct {
	// example:
	//
	// 100298370
	AppGroupId *string `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// example:
	//
	// 1588836130
	Created     *int64    `json:"created,omitempty" xml:"created,omitempty"`
	CustomDates []*string `json:"customDates,omitempty" xml:"customDates,omitempty" type:"Repeated"`
	// example:
	//
	// desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 100
	ElasticLcu *int32 `json:"elasticLcu,omitempty" xml:"elasticLcu,omitempty"`
	// example:
	//
	// True
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 15
	EndHour *int32 `json:"endHour,omitempty" xml:"endHour,omitempty"`
	// example:
	//
	// 89047
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// WEEK
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// example:
	//
	// 9
	StartHour *int32 `json:"startHour,omitempty" xml:"startHour,omitempty"`
	// example:
	//
	// 1588839490
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateElasticPlanResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticPlanResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanResponseBodyResult) GetAppGroupId() *string {
	return s.AppGroupId
}

func (s *CreateElasticPlanResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *CreateElasticPlanResponseBodyResult) GetCustomDates() []*string {
	return s.CustomDates
}

func (s *CreateElasticPlanResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *CreateElasticPlanResponseBodyResult) GetElasticLcu() *int32 {
	return s.ElasticLcu
}

func (s *CreateElasticPlanResponseBodyResult) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateElasticPlanResponseBodyResult) GetEndHour() *int32 {
	return s.EndHour
}

func (s *CreateElasticPlanResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateElasticPlanResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateElasticPlanResponseBodyResult) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateElasticPlanResponseBodyResult) GetStartHour() *int32 {
	return s.StartHour
}

func (s *CreateElasticPlanResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *CreateElasticPlanResponseBodyResult) SetAppGroupId(v string) *CreateElasticPlanResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetCreated(v int64) *CreateElasticPlanResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetCustomDates(v []*string) *CreateElasticPlanResponseBodyResult {
	s.CustomDates = v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetDescription(v string) *CreateElasticPlanResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetElasticLcu(v int32) *CreateElasticPlanResponseBodyResult {
	s.ElasticLcu = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetEnabled(v bool) *CreateElasticPlanResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetEndHour(v int32) *CreateElasticPlanResponseBodyResult {
	s.EndHour = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetId(v int64) *CreateElasticPlanResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetName(v string) *CreateElasticPlanResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetScheduleType(v string) *CreateElasticPlanResponseBodyResult {
	s.ScheduleType = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetStartHour(v int32) *CreateElasticPlanResponseBodyResult {
	s.StartHour = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) SetUpdated(v int64) *CreateElasticPlanResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateElasticPlanResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
