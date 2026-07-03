// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIncidentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIncidents(v []*ListIncidentsResponseBodyIncidents) *ListIncidentsResponseBody
	GetIncidents() []*ListIncidentsResponseBodyIncidents
	SetMaxResults(v int32) *ListIncidentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListIncidentsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListIncidentsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIncidentsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListIncidentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListIncidentsResponseBody
	GetTotalCount() *int32
}

type ListIncidentsResponseBody struct {
	// The list of incidents.
	Incidents []*ListIncidentsResponseBodyIncidents `json:"Incidents,omitempty" xml:"Incidents,omitempty" type:"Repeated"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIncidentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIncidentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIncidentsResponseBody) GetIncidents() []*ListIncidentsResponseBodyIncidents {
	return s.Incidents
}

func (s *ListIncidentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIncidentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIncidentsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIncidentsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIncidentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIncidentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIncidentsResponseBody) SetIncidents(v []*ListIncidentsResponseBodyIncidents) *ListIncidentsResponseBody {
	s.Incidents = v
	return s
}

func (s *ListIncidentsResponseBody) SetMaxResults(v int32) *ListIncidentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIncidentsResponseBody) SetNextToken(v string) *ListIncidentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIncidentsResponseBody) SetPageNumber(v int32) *ListIncidentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIncidentsResponseBody) SetPageSize(v int32) *ListIncidentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIncidentsResponseBody) SetRequestId(v string) *ListIncidentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIncidentsResponseBody) SetTotalCount(v int32) *ListIncidentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIncidentsResponseBody) Validate() error {
	if s.Incidents != nil {
		for _, item := range s.Incidents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIncidentsResponseBodyIncidents struct {
	// The creation time.
	//
	// example:
	//
	// 1603248483000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the detection rule.
	//
	// example:
	//
	// dr-qo5ww6ux0uc28*****
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// The name of the incident.
	//
	// example:
	//
	// ECS unusual log in
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// The remarks of the incident.
	//
	// example:
	//
	// remark
	IncidentRemark *string `json:"IncidentRemark,omitempty" xml:"IncidentRemark,omitempty"`
	// The status of the incident. Valid values:
	//
	// - 0: unhandled.
	//
	// - 1: handling.
	//
	// - 5: handling failed.
	//
	// - 10: handled.
	//
	// example:
	//
	// 0
	IncidentStatus *int32 `json:"IncidentStatus,omitempty" xml:"IncidentStatus,omitempty"`
	// The tags of the incident.
	//
	// example:
	//
	// ["sys:data_source:siem","sys:trigger_type:auto"]
	IncidentTags *string `json:"IncidentTags,omitempty" xml:"IncidentTags,omitempty"`
	// The UUID of the incident.
	//
	// example:
	//
	// dbb1d7211c9285c862aa89385098****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The UID of the account that owns the incident.
	//
	// example:
	//
	// 1234567890xxxxxx
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The number of alerts associated with the incident.
	//
	// example:
	//
	// 3
	RelateAlertCount *int32 `json:"RelateAlertCount,omitempty" xml:"RelateAlertCount,omitempty"`
	// The number of assets associated with the incident.
	//
	// example:
	//
	// 4
	RelateAssetCount *int32 `json:"RelateAssetCount,omitempty" xml:"RelateAssetCount,omitempty"`
	// The response time, in milliseconds (ms).
	//
	// example:
	//
	// 1603248483000
	ResponseTime *int64 `json:"ResponseTime,omitempty" xml:"ResponseTime,omitempty"`
	// The threat level. Valid values:
	//
	// - 5: critical.
	//
	// - 4: high.
	//
	// - 3: medium.
	//
	// - 2: low.
	//
	// - 1: informational.
	//
	// example:
	//
	// 2
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1603248483000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListIncidentsResponseBodyIncidents) String() string {
	return dara.Prettify(s)
}

func (s ListIncidentsResponseBodyIncidents) GoString() string {
	return s.String()
}

func (s *ListIncidentsResponseBodyIncidents) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListIncidentsResponseBodyIncidents) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *ListIncidentsResponseBodyIncidents) GetIncidentName() *string {
	return s.IncidentName
}

func (s *ListIncidentsResponseBodyIncidents) GetIncidentRemark() *string {
	return s.IncidentRemark
}

func (s *ListIncidentsResponseBodyIncidents) GetIncidentStatus() *int32 {
	return s.IncidentStatus
}

func (s *ListIncidentsResponseBodyIncidents) GetIncidentTags() *string {
	return s.IncidentTags
}

func (s *ListIncidentsResponseBodyIncidents) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *ListIncidentsResponseBodyIncidents) GetOwner() *string {
	return s.Owner
}

func (s *ListIncidentsResponseBodyIncidents) GetRelateAlertCount() *int32 {
	return s.RelateAlertCount
}

func (s *ListIncidentsResponseBodyIncidents) GetRelateAssetCount() *int32 {
	return s.RelateAssetCount
}

func (s *ListIncidentsResponseBodyIncidents) GetResponseTime() *int64 {
	return s.ResponseTime
}

func (s *ListIncidentsResponseBodyIncidents) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *ListIncidentsResponseBodyIncidents) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListIncidentsResponseBodyIncidents) SetCreateTime(v int64) *ListIncidentsResponseBodyIncidents {
	s.CreateTime = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetDetectionRuleId(v string) *ListIncidentsResponseBodyIncidents {
	s.DetectionRuleId = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetIncidentName(v string) *ListIncidentsResponseBodyIncidents {
	s.IncidentName = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetIncidentRemark(v string) *ListIncidentsResponseBodyIncidents {
	s.IncidentRemark = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetIncidentStatus(v int32) *ListIncidentsResponseBodyIncidents {
	s.IncidentStatus = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetIncidentTags(v string) *ListIncidentsResponseBodyIncidents {
	s.IncidentTags = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetIncidentUuid(v string) *ListIncidentsResponseBodyIncidents {
	s.IncidentUuid = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetOwner(v string) *ListIncidentsResponseBodyIncidents {
	s.Owner = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetRelateAlertCount(v int32) *ListIncidentsResponseBodyIncidents {
	s.RelateAlertCount = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetRelateAssetCount(v int32) *ListIncidentsResponseBodyIncidents {
	s.RelateAssetCount = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetResponseTime(v int64) *ListIncidentsResponseBodyIncidents {
	s.ResponseTime = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetThreatLevel(v string) *ListIncidentsResponseBodyIncidents {
	s.ThreatLevel = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) SetUpdateTime(v int64) *ListIncidentsResponseBodyIncidents {
	s.UpdateTime = &v
	return s
}

func (s *ListIncidentsResponseBodyIncidents) Validate() error {
	return dara.Validate(s)
}
