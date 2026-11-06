// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataProcessJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDataProcessJobsResponseBodyItems) *DescribeDataProcessJobsResponseBody
	GetItems() []*DescribeDataProcessJobsResponseBodyItems
	SetPageNumber(v string) *DescribeDataProcessJobsResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeDataProcessJobsResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeDataProcessJobsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeDataProcessJobsResponseBody
	GetTotalRecordCount() *string
}

type DescribeDataProcessJobsResponseBody struct {
	// The list of data processing jobs.
	Items []*DescribeDataProcessJobsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of jobs returned on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of jobs that match the query conditions.
	//
	// example:
	//
	// 1
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDataProcessJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobsResponseBody) GetItems() []*DescribeDataProcessJobsResponseBodyItems {
	return s.Items
}

func (s *DescribeDataProcessJobsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDataProcessJobsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeDataProcessJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataProcessJobsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeDataProcessJobsResponseBody) SetItems(v []*DescribeDataProcessJobsResponseBodyItems) *DescribeDataProcessJobsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataProcessJobsResponseBody) SetPageNumber(v string) *DescribeDataProcessJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBody) SetPageRecordCount(v string) *DescribeDataProcessJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBody) SetRequestId(v string) *DescribeDataProcessJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBody) SetTotalRecordCount(v string) *DescribeDataProcessJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataProcessJobsResponseBodyItems struct {
	// The time when the job was created, in RFC 3339 format.
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the data processing job.
	//
	// example:
	//
	// Image data processing job
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the destination PolarDB cluster.
	//
	// example:
	//
	// pc-test123456
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The name of the destination database.
	//
	// example:
	//
	// testdb
	DestinationPath *string `json:"DestinationPath,omitempty" xml:"DestinationPath,omitempty"`
	// The type of the destination. Currently, only PolarDB is supported.
	//
	// example:
	//
	// PolarDB
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The ID of the data processing job.
	//
	// example:
	//
	// job-123456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The data processing mode. Valid values: Full, Incremental, and Both.
	//
	// example:
	//
	// Full
	ProcessingMode *string `json:"ProcessingMode,omitempty" xml:"ProcessingMode,omitempty"`
	// The ID of the region where the data processing job resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the source PolarFS instance.
	//
	// example:
	//
	// pfs-123456
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source data path in PolarFS.
	//
	// example:
	//
	// /data/input
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// The type of the data source. Currently, only PolarFS is supported.
	//
	// example:
	//
	// PolarFS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The status of the data processing job.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the template used by the job.
	//
	// example:
	//
	// tpl-123456
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDataProcessJobsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetDestinationId() *string {
	return s.DestinationId
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetDestinationPath() *string {
	return s.DestinationPath
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetDestinationType() *string {
	return s.DestinationType
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetProcessingMode() *string {
	return s.ProcessingMode
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetSourceId() *string {
	return s.SourceId
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetSourcePath() *string {
	return s.SourcePath
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataProcessJobsResponseBodyItems) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetCreationTime(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.CreationTime = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetDescription(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetDestinationId(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.DestinationId = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetDestinationPath(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.DestinationPath = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetDestinationType(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.DestinationType = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetJobId(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.JobId = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetProcessingMode(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.ProcessingMode = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetRegionId(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetSourceId(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.SourceId = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetSourcePath(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.SourcePath = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetSourceType(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetStatus(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) SetTemplateId(v string) *DescribeDataProcessJobsResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *DescribeDataProcessJobsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
