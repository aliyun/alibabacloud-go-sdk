// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskListCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeRiskListCheckResultRequest
	GetCurrentPage() *int32
	SetInstanceIds(v []*string) *DescribeRiskListCheckResultRequest
	GetInstanceIds() []*string
	SetLang(v string) *DescribeRiskListCheckResultRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeRiskListCheckResultRequest
	GetPageSize() *int32
	SetResourceOwnerId(v int64) *DescribeRiskListCheckResultRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeRiskListCheckResultRequest
	GetSourceIp() *string
}

type DescribeRiskListCheckResultRequest struct {
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance IDs of the cloud services that you want to query. Separate multiple IDs with commas (,).
	//
	// > If you do not specify this parameter, an empty list is returned.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 59.57.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeRiskListCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskListCheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskListCheckResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRiskListCheckResultRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeRiskListCheckResultRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskListCheckResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRiskListCheckResultRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRiskListCheckResultRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskListCheckResultRequest) SetCurrentPage(v int32) *DescribeRiskListCheckResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetInstanceIds(v []*string) *DescribeRiskListCheckResultRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetLang(v string) *DescribeRiskListCheckResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetPageSize(v int32) *DescribeRiskListCheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetResourceOwnerId(v int64) *DescribeRiskListCheckResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetSourceIp(v string) *DescribeRiskListCheckResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
