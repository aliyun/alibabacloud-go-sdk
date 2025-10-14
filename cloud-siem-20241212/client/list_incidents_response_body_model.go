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
	Incidents []*ListIncidentsResponseBodyIncidents `json:"Incidents,omitempty" xml:"Incidents,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1603248483000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ECS unusual log in
	IncidentName *string `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	// example:
	//
	// remark
	IncidentRemark *string `json:"IncidentRemark,omitempty" xml:"IncidentRemark,omitempty"`
	// example:
	//
	// 0
	IncidentStatus *int32 `json:"IncidentStatus,omitempty" xml:"IncidentStatus,omitempty"`
	// example:
	//
	// dbb1d7211c9285c862aa89385098****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// example:
	//
	// 3
	RelateAlertCount *int32 `json:"RelateAlertCount,omitempty" xml:"RelateAlertCount,omitempty"`
	// example:
	//
	// 4
	RelateAssetCount *int32 `json:"RelateAssetCount,omitempty" xml:"RelateAssetCount,omitempty"`
	// example:
	//
	// 2
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
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

func (s *ListIncidentsResponseBodyIncidents) GetIncidentName() *string {
	return s.IncidentName
}

func (s *ListIncidentsResponseBodyIncidents) GetIncidentRemark() *string {
	return s.IncidentRemark
}

func (s *ListIncidentsResponseBodyIncidents) GetIncidentStatus() *int32 {
	return s.IncidentStatus
}

func (s *ListIncidentsResponseBodyIncidents) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *ListIncidentsResponseBodyIncidents) GetRelateAlertCount() *int32 {
	return s.RelateAlertCount
}

func (s *ListIncidentsResponseBodyIncidents) GetRelateAssetCount() *int32 {
	return s.RelateAssetCount
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

func (s *ListIncidentsResponseBodyIncidents) SetIncidentUuid(v string) *ListIncidentsResponseBodyIncidents {
	s.IncidentUuid = &v
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
