// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsDatasourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApsDatasources(v []*DescribeApsDatasourcesResponseBodyApsDatasources) *DescribeApsDatasourcesResponseBody
	GetApsDatasources() []*DescribeApsDatasourcesResponseBodyApsDatasources
	SetPageNumber(v string) *DescribeApsDatasourcesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeApsDatasourcesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeApsDatasourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeApsDatasourcesResponseBody
	GetTotalCount() *string
}

type DescribeApsDatasourcesResponseBody struct {
	// The queried APS data sources.
	//
	// example:
	//
	// -
	ApsDatasources []*DescribeApsDatasourcesResponseBodyApsDatasources `json:"ApsDatasources,omitempty" xml:"ApsDatasources,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-5213-******-B608-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApsDatasourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourcesResponseBody) GetApsDatasources() []*DescribeApsDatasourcesResponseBodyApsDatasources {
	return s.ApsDatasources
}

func (s *DescribeApsDatasourcesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeApsDatasourcesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeApsDatasourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApsDatasourcesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeApsDatasourcesResponseBody) SetApsDatasources(v []*DescribeApsDatasourcesResponseBodyApsDatasources) *DescribeApsDatasourcesResponseBody {
	s.ApsDatasources = v
	return s
}

func (s *DescribeApsDatasourcesResponseBody) SetPageNumber(v string) *DescribeApsDatasourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBody) SetPageSize(v string) *DescribeApsDatasourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBody) SetRequestId(v string) *DescribeApsDatasourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBody) SetTotalCount(v string) *DescribeApsDatasourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBody) Validate() error {
	if s.ApsDatasources != nil {
		for _, item := range s.ApsDatasources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApsDatasourcesResponseBodyApsDatasources struct {
	// The time when the data source was created.
	//
	// example:
	//
	// 2024-01-10 14:44:33
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// test
	DatasourceDescription *string `json:"DatasourceDescription,omitempty" xml:"DatasourceDescription,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// SLS
	DatasourceType *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// Indicates whether a job is using the data source.
	//
	// example:
	//
	// false
	HasJob *bool `json:"HasJob,omitempty" xml:"HasJob,omitempty"`
}

func (s DescribeApsDatasourcesResponseBodyApsDatasources) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourcesResponseBodyApsDatasources) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) GetDatasourceDescription() *string {
	return s.DatasourceDescription
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) GetHasJob() *bool {
	return s.HasJob
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) SetCreateTime(v string) *DescribeApsDatasourcesResponseBodyApsDatasources {
	s.CreateTime = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) SetDatasourceDescription(v string) *DescribeApsDatasourcesResponseBodyApsDatasources {
	s.DatasourceDescription = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) SetDatasourceId(v int64) *DescribeApsDatasourcesResponseBodyApsDatasources {
	s.DatasourceId = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) SetDatasourceName(v string) *DescribeApsDatasourcesResponseBodyApsDatasources {
	s.DatasourceName = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) SetDatasourceType(v string) *DescribeApsDatasourcesResponseBodyApsDatasources {
	s.DatasourceType = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) SetHasJob(v bool) *DescribeApsDatasourcesResponseBodyApsDatasources {
	s.HasJob = &v
	return s
}

func (s *DescribeApsDatasourcesResponseBodyApsDatasources) Validate() error {
	return dara.Validate(s)
}
