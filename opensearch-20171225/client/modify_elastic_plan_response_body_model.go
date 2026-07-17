// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyElasticPlanResponseBody
	GetCode() *string
	SetHttpCode(v int64) *ModifyElasticPlanResponseBody
	GetHttpCode() *int64
	SetMessage(v string) *ModifyElasticPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyElasticPlanResponseBody
	GetRequestId() *string
	SetResult(v *ModifyElasticPlanResponseBodyResult) *ModifyElasticPlanResponseBody
	GetResult() *ModifyElasticPlanResponseBodyResult
}

type ModifyElasticPlanResponseBody struct {
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
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *ModifyElasticPlanResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyElasticPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyElasticPlanResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *ModifyElasticPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyElasticPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElasticPlanResponseBody) GetResult() *ModifyElasticPlanResponseBodyResult {
	return s.Result
}

func (s *ModifyElasticPlanResponseBody) SetCode(v string) *ModifyElasticPlanResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyElasticPlanResponseBody) SetHttpCode(v int64) *ModifyElasticPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ModifyElasticPlanResponseBody) SetMessage(v string) *ModifyElasticPlanResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyElasticPlanResponseBody) SetRequestId(v string) *ModifyElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElasticPlanResponseBody) SetResult(v *ModifyElasticPlanResponseBodyResult) *ModifyElasticPlanResponseBody {
	s.Result = v
	return s
}

func (s *ModifyElasticPlanResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyElasticPlanResponseBodyResult struct {
	// example:
	//
	// 120123456
	AppGroupId *int64 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// example:
	//
	// 1588839490
	Created     *int64    `json:"created,omitempty" xml:"created,omitempty"`
	CustomDates []*string `json:"customDates,omitempty" xml:"customDates,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 100
	ElasticLcu *int32 `json:"elasticLcu,omitempty" xml:"elasticLcu,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 10
	EndHour *int32 `json:"endHour,omitempty" xml:"endHour,omitempty"`
	// example:
	//
	// 286
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// plan name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// WEEK
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// example:
	//
	// 6
	StartHour *int32 `json:"startHour,omitempty" xml:"startHour,omitempty"`
	// example:
	//
	// 1539158313
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyElasticPlanResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticPlanResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponseBodyResult) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *ModifyElasticPlanResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *ModifyElasticPlanResponseBodyResult) GetCustomDates() []*string {
	return s.CustomDates
}

func (s *ModifyElasticPlanResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ModifyElasticPlanResponseBodyResult) GetElasticLcu() *int32 {
	return s.ElasticLcu
}

func (s *ModifyElasticPlanResponseBodyResult) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyElasticPlanResponseBodyResult) GetEndHour() *int32 {
	return s.EndHour
}

func (s *ModifyElasticPlanResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ModifyElasticPlanResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ModifyElasticPlanResponseBodyResult) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ModifyElasticPlanResponseBodyResult) GetStartHour() *int32 {
	return s.StartHour
}

func (s *ModifyElasticPlanResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *ModifyElasticPlanResponseBodyResult) SetAppGroupId(v int64) *ModifyElasticPlanResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetCreated(v int64) *ModifyElasticPlanResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetCustomDates(v []*string) *ModifyElasticPlanResponseBodyResult {
	s.CustomDates = v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetDescription(v string) *ModifyElasticPlanResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetElasticLcu(v int32) *ModifyElasticPlanResponseBodyResult {
	s.ElasticLcu = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetEnabled(v bool) *ModifyElasticPlanResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetEndHour(v int32) *ModifyElasticPlanResponseBodyResult {
	s.EndHour = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetId(v int64) *ModifyElasticPlanResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetName(v string) *ModifyElasticPlanResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetScheduleType(v string) *ModifyElasticPlanResponseBodyResult {
	s.ScheduleType = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetStartHour(v int32) *ModifyElasticPlanResponseBodyResult {
	s.StartHour = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) SetUpdated(v int64) *ModifyElasticPlanResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ModifyElasticPlanResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
