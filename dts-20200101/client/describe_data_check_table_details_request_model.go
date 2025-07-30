// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCheckTableDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckType(v int32) *DescribeDataCheckTableDetailsRequest
	GetCheckType() *int32
	SetDtsJobId(v string) *DescribeDataCheckTableDetailsRequest
	GetDtsJobId() *string
	SetPageNumber(v int32) *DescribeDataCheckTableDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataCheckTableDetailsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeDataCheckTableDetailsRequest
	GetResourceGroupId() *string
	SetSchemaName(v string) *DescribeDataCheckTableDetailsRequest
	GetSchemaName() *string
	SetStatus(v string) *DescribeDataCheckTableDetailsRequest
	GetStatus() *string
	SetTableName(v string) *DescribeDataCheckTableDetailsRequest
	GetTableName() *string
}

type DescribeDataCheckTableDetailsRequest struct {
	// The data verification method. Valid values:
	//
	// 	- **1**: full data verification.
	//
	// 	- **2**: incremental data verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The ID of the data migration or data synchronization task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xd4e4xb419q****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than **0*	- and does not exceed the maximum value of the Integer data type. Default value:**1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the schema whose data is verified in the source database.
	//
	// example:
	//
	// dtstest
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The status of the data verification result. Valid values:
	//
	// 	- **-1*	- (default): All status.
	//
	// 	- **6**: Inconsistent data detected in the table.
	//
	// example:
	//
	// -1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the table whose data is verified in the source database.
	//
	// example:
	//
	// student
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeDataCheckTableDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckTableDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckTableDetailsRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *DescribeDataCheckTableDetailsRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDataCheckTableDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataCheckTableDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataCheckTableDetailsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDataCheckTableDetailsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeDataCheckTableDetailsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataCheckTableDetailsRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDataCheckTableDetailsRequest) SetCheckType(v int32) *DescribeDataCheckTableDetailsRequest {
	s.CheckType = &v
	return s
}

func (s *DescribeDataCheckTableDetailsRequest) SetDtsJobId(v string) *DescribeDataCheckTableDetailsRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDataCheckTableDetailsRequest) SetPageNumber(v int32) *DescribeDataCheckTableDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataCheckTableDetailsRequest) SetPageSize(v int32) *DescribeDataCheckTableDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataCheckTableDetailsRequest) SetResourceGroupId(v string) *DescribeDataCheckTableDetailsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataCheckTableDetailsRequest) SetSchemaName(v string) *DescribeDataCheckTableDetailsRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeDataCheckTableDetailsRequest) SetStatus(v string) *DescribeDataCheckTableDetailsRequest {
	s.Status = &v
	return s
}

func (s *DescribeDataCheckTableDetailsRequest) SetTableName(v string) *DescribeDataCheckTableDetailsRequest {
	s.TableName = &v
	return s
}

func (s *DescribeDataCheckTableDetailsRequest) Validate() error {
	return dara.Validate(s)
}
