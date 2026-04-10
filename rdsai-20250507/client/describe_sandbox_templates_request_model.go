// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DescribeSandboxTemplatesRequest
	GetInstanceName() *string
	SetMaxResults(v int32) *DescribeSandboxTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSandboxTemplatesRequest
	GetNextToken() *string
	SetPageNumber(v int64) *DescribeSandboxTemplatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSandboxTemplatesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSandboxTemplatesRequest
	GetRegionId() *string
	SetTemplateName(v string) *DescribeSandboxTemplatesRequest
	GetTemplateName() *string
}

type DescribeSandboxTemplatesRequest struct {
	// This parameter is required.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// None
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// gAAAAABo-D-ze1Zog63nWMa8eDDMkqUoBB5-FDsHDUMiNIDSDZeP9g0LwJEozulOPG_LbsGwLRgmDFvTHZeSU90YsukT0pHtnA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// code-interpreter
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeSandboxTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxTemplatesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSandboxTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSandboxTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSandboxTemplatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSandboxTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSandboxTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSandboxTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeSandboxTemplatesRequest) SetInstanceName(v string) *DescribeSandboxTemplatesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeSandboxTemplatesRequest) SetMaxResults(v int32) *DescribeSandboxTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSandboxTemplatesRequest) SetNextToken(v string) *DescribeSandboxTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSandboxTemplatesRequest) SetPageNumber(v int64) *DescribeSandboxTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSandboxTemplatesRequest) SetPageSize(v int64) *DescribeSandboxTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSandboxTemplatesRequest) SetRegionId(v string) *DescribeSandboxTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSandboxTemplatesRequest) SetTemplateName(v string) *DescribeSandboxTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeSandboxTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
