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
	DataModule []*DescribeCustomScenePoliciesResponseBodyDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
	PageNumber *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Quota      *int32                                               `json:"Quota,omitempty" xml:"Quota,omitempty"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	EndTime   *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Name      *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Objects   []*string `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	PolicyId  *int64    `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SiteIds   *string   `json:"SiteIds,omitempty" xml:"SiteIds,omitempty"`
	StartTime *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Template  *string   `json:"Template,omitempty" xml:"Template,omitempty"`
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
