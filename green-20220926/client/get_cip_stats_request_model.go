// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCipStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetByMonth(v bool) *GetCipStatsRequest
	GetByMonth() *bool
	SetEndDate(v string) *GetCipStatsRequest
	GetEndDate() *string
	SetLabel(v string) *GetCipStatsRequest
	GetLabel() *string
	SetQuery(v string) *GetCipStatsRequest
	GetQuery() *string
	SetRegionId(v string) *GetCipStatsRequest
	GetRegionId() *string
	SetResourceType(v string) *GetCipStatsRequest
	GetResourceType() *string
	SetServiceCode(v string) *GetCipStatsRequest
	GetServiceCode() *string
	SetStartDate(v string) *GetCipStatsRequest
	GetStartDate() *string
	SetSubUid(v string) *GetCipStatsRequest
	GetSubUid() *string
	SetType(v string) *GetCipStatsRequest
	GetType() *string
}

type GetCipStatsRequest struct {
	// Specifies whether monthly indexing is supported. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
	//
	// example:
	//
	// true
	ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
	// The end time of the query. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2024-03-11 10:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The label.
	//
	// example:
	//
	// xx
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The query condition.
	//
	// example:
	//
	// {}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service code.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The start time of the query. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2024-03-10 10:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The UID of the RAM user.
	//
	// example:
	//
	// 253552244990701265
	SubUid *string `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
	// The type. Valid values:
	//
	// - **cip**: Content Moderation invocation usage statistics.
	//
	// - **risk_level**: Content Moderation risk level statistics.
	//
	// - **content_moderation**: AI safety guardrail content compliance risk level and tag statistics.
	//
	// - **sensitive_data**: AI safety guardrail sensitive data risk level and tag statistics.
	//
	// - **prompt_attack**: AI safety guardrail prompt risk level and tag statistics.
	//
	// example:
	//
	// risk_level
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCipStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCipStatsRequest) GoString() string {
	return s.String()
}

func (s *GetCipStatsRequest) GetByMonth() *bool {
	return s.ByMonth
}

func (s *GetCipStatsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetCipStatsRequest) GetLabel() *string {
	return s.Label
}

func (s *GetCipStatsRequest) GetQuery() *string {
	return s.Query
}

func (s *GetCipStatsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCipStatsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetCipStatsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetCipStatsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetCipStatsRequest) GetSubUid() *string {
	return s.SubUid
}

func (s *GetCipStatsRequest) GetType() *string {
	return s.Type
}

func (s *GetCipStatsRequest) SetByMonth(v bool) *GetCipStatsRequest {
	s.ByMonth = &v
	return s
}

func (s *GetCipStatsRequest) SetEndDate(v string) *GetCipStatsRequest {
	s.EndDate = &v
	return s
}

func (s *GetCipStatsRequest) SetLabel(v string) *GetCipStatsRequest {
	s.Label = &v
	return s
}

func (s *GetCipStatsRequest) SetQuery(v string) *GetCipStatsRequest {
	s.Query = &v
	return s
}

func (s *GetCipStatsRequest) SetRegionId(v string) *GetCipStatsRequest {
	s.RegionId = &v
	return s
}

func (s *GetCipStatsRequest) SetResourceType(v string) *GetCipStatsRequest {
	s.ResourceType = &v
	return s
}

func (s *GetCipStatsRequest) SetServiceCode(v string) *GetCipStatsRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetCipStatsRequest) SetStartDate(v string) *GetCipStatsRequest {
	s.StartDate = &v
	return s
}

func (s *GetCipStatsRequest) SetSubUid(v string) *GetCipStatsRequest {
	s.SubUid = &v
	return s
}

func (s *GetCipStatsRequest) SetType(v string) *GetCipStatsRequest {
	s.Type = &v
	return s
}

func (s *GetCipStatsRequest) Validate() error {
	return dara.Validate(s)
}
