// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckResourceList(v []*DescribeNisInspectionTaskResponseBodyCheckResourceList) *DescribeNisInspectionTaskResponseBody
	GetCheckResourceList() []*DescribeNisInspectionTaskResponseBodyCheckResourceList
	SetCreateTime(v string) *DescribeNisInspectionTaskResponseBody
	GetCreateTime() *string
	SetInspectionInterval(v string) *DescribeNisInspectionTaskResponseBody
	GetInspectionInterval() *string
	SetInspectionName(v string) *DescribeNisInspectionTaskResponseBody
	GetInspectionName() *string
	SetInspectionProject(v string) *DescribeNisInspectionTaskResponseBody
	GetInspectionProject() *string
	SetInspectionTaskId(v string) *DescribeNisInspectionTaskResponseBody
	GetInspectionTaskId() *string
	SetInspectionTriggerTime(v string) *DescribeNisInspectionTaskResponseBody
	GetInspectionTriggerTime() *string
	SetLastUpdateReportId(v string) *DescribeNisInspectionTaskResponseBody
	GetLastUpdateReportId() *string
	SetLastUpdateTime(v string) *DescribeNisInspectionTaskResponseBody
	GetLastUpdateTime() *string
	SetRequestId(v string) *DescribeNisInspectionTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeNisInspectionTaskResponseBody
	GetStatus() *string
}

type DescribeNisInspectionTaskResponseBody struct {
	CheckResourceList []*DescribeNisInspectionTaskResponseBodyCheckResourceList `json:"CheckResourceList,omitempty" xml:"CheckResourceList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-07-01 10:00:57
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	InspectionInterval *string `json:"InspectionInterval,omitempty" xml:"InspectionInterval,omitempty"`
	// example:
	//
	// Default
	InspectionName *string `json:"InspectionName,omitempty" xml:"InspectionName,omitempty"`
	// example:
	//
	// basic
	InspectionProject *string `json:"InspectionProject,omitempty" xml:"InspectionProject,omitempty"`
	// example:
	//
	// ni-8svmpe0yso2bhzr7fh79
	InspectionTaskId *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
	// example:
	//
	// 2024-07-01 15:15:57
	InspectionTriggerTime *string `json:"InspectionTriggerTime,omitempty" xml:"InspectionTriggerTime,omitempty"`
	// example:
	//
	// nir-7c3dd178738a429abe6d
	LastUpdateReportId *string `json:"LastUpdateReportId,omitempty" xml:"LastUpdateReportId,omitempty"`
	// example:
	//
	// 2024-07-01 10:00:59
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNisInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionTaskResponseBody) GetCheckResourceList() []*DescribeNisInspectionTaskResponseBodyCheckResourceList {
	return s.CheckResourceList
}

func (s *DescribeNisInspectionTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeNisInspectionTaskResponseBody) GetInspectionInterval() *string {
	return s.InspectionInterval
}

func (s *DescribeNisInspectionTaskResponseBody) GetInspectionName() *string {
	return s.InspectionName
}

func (s *DescribeNisInspectionTaskResponseBody) GetInspectionProject() *string {
	return s.InspectionProject
}

func (s *DescribeNisInspectionTaskResponseBody) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *DescribeNisInspectionTaskResponseBody) GetInspectionTriggerTime() *string {
	return s.InspectionTriggerTime
}

func (s *DescribeNisInspectionTaskResponseBody) GetLastUpdateReportId() *string {
	return s.LastUpdateReportId
}

func (s *DescribeNisInspectionTaskResponseBody) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *DescribeNisInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNisInspectionTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNisInspectionTaskResponseBody) SetCheckResourceList(v []*DescribeNisInspectionTaskResponseBodyCheckResourceList) *DescribeNisInspectionTaskResponseBody {
	s.CheckResourceList = v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetCreateTime(v string) *DescribeNisInspectionTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetInspectionInterval(v string) *DescribeNisInspectionTaskResponseBody {
	s.InspectionInterval = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetInspectionName(v string) *DescribeNisInspectionTaskResponseBody {
	s.InspectionName = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetInspectionProject(v string) *DescribeNisInspectionTaskResponseBody {
	s.InspectionProject = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetInspectionTaskId(v string) *DescribeNisInspectionTaskResponseBody {
	s.InspectionTaskId = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetInspectionTriggerTime(v string) *DescribeNisInspectionTaskResponseBody {
	s.InspectionTriggerTime = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetLastUpdateReportId(v string) *DescribeNisInspectionTaskResponseBody {
	s.LastUpdateReportId = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetLastUpdateTime(v string) *DescribeNisInspectionTaskResponseBody {
	s.LastUpdateTime = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetRequestId(v string) *DescribeNisInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) SetStatus(v string) *DescribeNisInspectionTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBody) Validate() error {
	if s.CheckResourceList != nil {
		for _, item := range s.CheckResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNisInspectionTaskResponseBodyCheckResourceList struct {
	// example:
	//
	// CheckAll
	CheckScope *string `json:"CheckScope,omitempty" xml:"CheckScope,omitempty"`
	// example:
	//
	// EIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeNisInspectionTaskResponseBodyCheckResourceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionTaskResponseBodyCheckResourceList) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionTaskResponseBodyCheckResourceList) GetCheckScope() *string {
	return s.CheckScope
}

func (s *DescribeNisInspectionTaskResponseBodyCheckResourceList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeNisInspectionTaskResponseBodyCheckResourceList) SetCheckScope(v string) *DescribeNisInspectionTaskResponseBodyCheckResourceList {
	s.CheckScope = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBodyCheckResourceList) SetResourceType(v string) *DescribeNisInspectionTaskResponseBodyCheckResourceList {
	s.ResourceType = &v
	return s
}

func (s *DescribeNisInspectionTaskResponseBodyCheckResourceList) Validate() error {
	return dara.Validate(s)
}
