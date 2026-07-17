// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeElasticPlanResponseBody
	GetCode() *string
	SetHttpCode(v int64) *DescribeElasticPlanResponseBody
	GetHttpCode() *int64
	SetMessage(v string) *DescribeElasticPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeElasticPlanResponseBody
	GetRequestId() *string
	SetResult(v *DescribeElasticPlanResponseBodyResult) *DescribeElasticPlanResponseBody
	GetResult() *DescribeElasticPlanResponseBodyResult
}

type DescribeElasticPlanResponseBody struct {
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
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DescribeElasticPlanResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeElasticPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeElasticPlanResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *DescribeElasticPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeElasticPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticPlanResponseBody) GetResult() *DescribeElasticPlanResponseBodyResult {
	return s.Result
}

func (s *DescribeElasticPlanResponseBody) SetCode(v string) *DescribeElasticPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeElasticPlanResponseBody) SetHttpCode(v int64) *DescribeElasticPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DescribeElasticPlanResponseBody) SetMessage(v string) *DescribeElasticPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeElasticPlanResponseBody) SetRequestId(v string) *DescribeElasticPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlanResponseBody) SetResult(v *DescribeElasticPlanResponseBodyResult) *DescribeElasticPlanResponseBody {
	s.Result = v
	return s
}

func (s *DescribeElasticPlanResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeElasticPlanResponseBodyResult struct {
	// example:
	//
	// 120123456
	AppGroupId *int64 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// example:
	//
	// 1590139542
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
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 13
	EndHour *int32 `json:"endHour,omitempty" xml:"endHour,omitempty"`
	// example:
	//
	// 11
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
	// 9
	StartHour *int32 `json:"startHour,omitempty" xml:"startHour,omitempty"`
	// example:
	//
	// 1581065904
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeElasticPlanResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanResponseBodyResult) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *DescribeElasticPlanResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *DescribeElasticPlanResponseBodyResult) GetCustomDates() []*string {
	return s.CustomDates
}

func (s *DescribeElasticPlanResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeElasticPlanResponseBodyResult) GetElasticLcu() *int32 {
	return s.ElasticLcu
}

func (s *DescribeElasticPlanResponseBodyResult) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeElasticPlanResponseBodyResult) GetEndHour() *int32 {
	return s.EndHour
}

func (s *DescribeElasticPlanResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *DescribeElasticPlanResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeElasticPlanResponseBodyResult) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *DescribeElasticPlanResponseBodyResult) GetStartHour() *int32 {
	return s.StartHour
}

func (s *DescribeElasticPlanResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *DescribeElasticPlanResponseBodyResult) SetAppGroupId(v int64) *DescribeElasticPlanResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetCreated(v int64) *DescribeElasticPlanResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetCustomDates(v []*string) *DescribeElasticPlanResponseBodyResult {
	s.CustomDates = v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetDescription(v string) *DescribeElasticPlanResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetElasticLcu(v int32) *DescribeElasticPlanResponseBodyResult {
	s.ElasticLcu = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetEnabled(v bool) *DescribeElasticPlanResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetEndHour(v int32) *DescribeElasticPlanResponseBodyResult {
	s.EndHour = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetId(v int64) *DescribeElasticPlanResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetName(v string) *DescribeElasticPlanResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetScheduleType(v string) *DescribeElasticPlanResponseBodyResult {
	s.ScheduleType = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetStartHour(v int32) *DescribeElasticPlanResponseBodyResult {
	s.StartHour = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) SetUpdated(v int64) *DescribeElasticPlanResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeElasticPlanResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
