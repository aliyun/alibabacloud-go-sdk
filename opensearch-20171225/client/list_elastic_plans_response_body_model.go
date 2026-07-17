// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListElasticPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListElasticPlansResponseBody
	GetCode() *string
	SetHttpCode(v int64) *ListElasticPlansResponseBody
	GetHttpCode() *int64
	SetMaxResults(v int32) *ListElasticPlansResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListElasticPlansResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListElasticPlansResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListElasticPlansResponseBody
	GetRequestId() *string
	SetResult(v []*ListElasticPlansResponseBodyResult) *ListElasticPlansResponseBody
	GetResult() []*ListElasticPlansResponseBodyResult
	SetTotalCount(v int32) *ListElasticPlansResponseBody
	GetTotalCount() *int32
}

type ListElasticPlansResponseBody struct {
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
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// Elastic plan not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 20
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListElasticPlansResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListElasticPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListElasticPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListElasticPlansResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListElasticPlansResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *ListElasticPlansResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListElasticPlansResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListElasticPlansResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListElasticPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListElasticPlansResponseBody) GetResult() []*ListElasticPlansResponseBodyResult {
	return s.Result
}

func (s *ListElasticPlansResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListElasticPlansResponseBody) SetCode(v string) *ListElasticPlansResponseBody {
	s.Code = &v
	return s
}

func (s *ListElasticPlansResponseBody) SetHttpCode(v int64) *ListElasticPlansResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListElasticPlansResponseBody) SetMaxResults(v int32) *ListElasticPlansResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListElasticPlansResponseBody) SetMessage(v string) *ListElasticPlansResponseBody {
	s.Message = &v
	return s
}

func (s *ListElasticPlansResponseBody) SetNextToken(v string) *ListElasticPlansResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListElasticPlansResponseBody) SetRequestId(v string) *ListElasticPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListElasticPlansResponseBody) SetResult(v []*ListElasticPlansResponseBodyResult) *ListElasticPlansResponseBody {
	s.Result = v
	return s
}

func (s *ListElasticPlansResponseBody) SetTotalCount(v int32) *ListElasticPlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListElasticPlansResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListElasticPlansResponseBodyResult struct {
	// example:
	//
	// 120123456
	AppGroupId *string `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// example:
	//
	// 1588839490
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
	// 16
	EndHour *int32 `json:"endHour,omitempty" xml:"endHour,omitempty"`
	// example:
	//
	// 134
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// elastic plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// WEEK
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// example:
	//
	// 11
	StartHour *int32 `json:"startHour,omitempty" xml:"startHour,omitempty"`
	// example:
	//
	// 1588839490
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListElasticPlansResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListElasticPlansResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListElasticPlansResponseBodyResult) GetAppGroupId() *string {
	return s.AppGroupId
}

func (s *ListElasticPlansResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *ListElasticPlansResponseBodyResult) GetCustomDates() []*string {
	return s.CustomDates
}

func (s *ListElasticPlansResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListElasticPlansResponseBodyResult) GetElasticLcu() *int32 {
	return s.ElasticLcu
}

func (s *ListElasticPlansResponseBodyResult) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListElasticPlansResponseBodyResult) GetEndHour() *int32 {
	return s.EndHour
}

func (s *ListElasticPlansResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListElasticPlansResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListElasticPlansResponseBodyResult) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListElasticPlansResponseBodyResult) GetStartHour() *int32 {
	return s.StartHour
}

func (s *ListElasticPlansResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *ListElasticPlansResponseBodyResult) SetAppGroupId(v string) *ListElasticPlansResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetCreated(v int64) *ListElasticPlansResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetCustomDates(v []*string) *ListElasticPlansResponseBodyResult {
	s.CustomDates = v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetDescription(v string) *ListElasticPlansResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetElasticLcu(v int32) *ListElasticPlansResponseBodyResult {
	s.ElasticLcu = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetEnabled(v bool) *ListElasticPlansResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetEndHour(v int32) *ListElasticPlansResponseBodyResult {
	s.EndHour = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetId(v int64) *ListElasticPlansResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetName(v string) *ListElasticPlansResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetScheduleType(v string) *ListElasticPlansResponseBodyResult {
	s.ScheduleType = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetStartHour(v int32) *ListElasticPlansResponseBodyResult {
	s.StartHour = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) SetUpdated(v int64) *ListElasticPlansResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListElasticPlansResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
