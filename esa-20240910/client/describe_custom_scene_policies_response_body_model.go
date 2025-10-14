// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomScenePoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataModule(v []*DescribeCustomScenePoliciesResponseBodyDataModule) *DescribeCustomScenePoliciesResponseBody
	GetDataModule() []*DescribeCustomScenePoliciesResponseBodyDataModule
	SetPageNumber(v int32) *DescribeCustomScenePoliciesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCustomScenePoliciesResponseBody
	GetPageSize() *int32
	SetQuota(v int32) *DescribeCustomScenePoliciesResponseBody
	GetQuota() *int32
	SetRequestId(v string) *DescribeCustomScenePoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCustomScenePoliciesResponseBody
	GetTotalCount() *int32
}

type DescribeCustomScenePoliciesResponseBody struct {
	// The scenario-specific policies.
	DataModule []*DescribeCustomScenePoliciesResponseBodyDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
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
	// The policy quota.
	//
	// example:
	//
	// 10
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 85H66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCustomScenePoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomScenePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomScenePoliciesResponseBody) GetDataModule() []*DescribeCustomScenePoliciesResponseBodyDataModule {
	return s.DataModule
}

func (s *DescribeCustomScenePoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCustomScenePoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomScenePoliciesResponseBody) GetQuota() *int32 {
	return s.Quota
}

func (s *DescribeCustomScenePoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomScenePoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCustomScenePoliciesResponseBody) SetDataModule(v []*DescribeCustomScenePoliciesResponseBodyDataModule) *DescribeCustomScenePoliciesResponseBody {
	s.DataModule = v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetPageNumber(v int32) *DescribeCustomScenePoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetPageSize(v int32) *DescribeCustomScenePoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetQuota(v int32) *DescribeCustomScenePoliciesResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetRequestId(v string) *DescribeCustomScenePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) SetTotalCount(v int32) *DescribeCustomScenePoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBody) Validate() error {
	if s.DataModule != nil {
		for _, item := range s.DataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomScenePoliciesResponseBodyDataModule struct {
	// The time when the policy expires.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-06T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the scenario-specific policy.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of websites that are associated with the policy.
	Objects []*string `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	// The policy ID.
	//
	// example:
	//
	// 1234****
	PolicyId *int64  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SiteIds  *string `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	// The time when the policy takes effect.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-04T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- **Disabled**
	//
	// 	- **Pending**
	//
	// 	- **Running**
	//
	// 	- **Expired**
	//
	// example:
	//
	// Expired
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the policy template. Valid value:
	//
	// 	- **promotion**: major events.
	//
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s DescribeCustomScenePoliciesResponseBodyDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomScenePoliciesResponseBodyDataModule) GoString() string {
	return s.String()
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) GetName() *string {
	return s.Name
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) GetObjects() []*string {
	return s.Objects
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) GetSiteIds() *string {
	return s.SiteIds
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) GetStatus() *string {
	return s.Status
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) GetTemplate() *string {
	return s.Template
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetEndTime(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetName(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.Name = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetObjects(v []*string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.Objects = v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetPolicyId(v int64) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.PolicyId = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetSiteIds(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.SiteIds = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetStartTime(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetStatus(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.Status = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) SetTemplate(v string) *DescribeCustomScenePoliciesResponseBodyDataModule {
	s.Template = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponseBodyDataModule) Validate() error {
	return dara.Validate(s)
}
