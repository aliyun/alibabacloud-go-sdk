// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCheckReportUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckType(v int32) *DescribeDataCheckReportUrlRequest
	GetCheckType() *int32
	SetDbName(v string) *DescribeDataCheckReportUrlRequest
	GetDbName() *string
	SetDtsJobId(v string) *DescribeDataCheckReportUrlRequest
	GetDtsJobId() *string
	SetResourceGroupId(v string) *DescribeDataCheckReportUrlRequest
	GetResourceGroupId() *string
	SetTbName(v string) *DescribeDataCheckReportUrlRequest
	GetTbName() *string
}

type DescribeDataCheckReportUrlRequest struct {
	// The data verification method. Currently, only the value **1*	- is supported, which indicates full verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The name of the source database to be verified.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsdb
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DTS task. You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f4612nr2182****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the table in the source database to be verified.
	//
	// This parameter is required.
	//
	// example:
	//
	// student
	TbName *string `json:"TbName,omitempty" xml:"TbName,omitempty"`
}

func (s DescribeDataCheckReportUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckReportUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckReportUrlRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *DescribeDataCheckReportUrlRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDataCheckReportUrlRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDataCheckReportUrlRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDataCheckReportUrlRequest) GetTbName() *string {
	return s.TbName
}

func (s *DescribeDataCheckReportUrlRequest) SetCheckType(v int32) *DescribeDataCheckReportUrlRequest {
	s.CheckType = &v
	return s
}

func (s *DescribeDataCheckReportUrlRequest) SetDbName(v string) *DescribeDataCheckReportUrlRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDataCheckReportUrlRequest) SetDtsJobId(v string) *DescribeDataCheckReportUrlRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDataCheckReportUrlRequest) SetResourceGroupId(v string) *DescribeDataCheckReportUrlRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataCheckReportUrlRequest) SetTbName(v string) *DescribeDataCheckReportUrlRequest {
	s.TbName = &v
	return s
}

func (s *DescribeDataCheckReportUrlRequest) Validate() error {
	return dara.Validate(s)
}
