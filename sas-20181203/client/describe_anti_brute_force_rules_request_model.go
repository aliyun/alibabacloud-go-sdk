// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAntiBruteForceRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeAntiBruteForceRulesRequest
	GetCurrentPage() *int32
	SetId(v int64) *DescribeAntiBruteForceRulesRequest
	GetId() *int64
	SetName(v string) *DescribeAntiBruteForceRulesRequest
	GetName() *string
	SetPageSize(v string) *DescribeAntiBruteForceRulesRequest
	GetPageSize() *string
	SetResourceOwnerId(v int64) *DescribeAntiBruteForceRulesRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeAntiBruteForceRulesRequest
	GetSourceIp() *string
}

type DescribeAntiBruteForceRulesRequest struct {
	// Set which page of the returned results to start displaying the query results. The default value is **1**, indicating that the display starts from the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the anti-brute force rule.
	//
	// > You can obtain this parameter by calling the [DescribeAntiBruteForceRules](~~DescribeAntiBruteForceRules~~) interface.
	//
	// example:
	//
	// 1141****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the brute force rule.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The maximum number of data entries displayed per page during a paginated query.
	//
	// example:
	//
	// 10
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 121.69.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAntiBruteForceRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAntiBruteForceRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAntiBruteForceRulesRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeAntiBruteForceRulesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeAntiBruteForceRulesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAntiBruteForceRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAntiBruteForceRulesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAntiBruteForceRulesRequest) SetCurrentPage(v int32) *DescribeAntiBruteForceRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAntiBruteForceRulesRequest) SetId(v int64) *DescribeAntiBruteForceRulesRequest {
	s.Id = &v
	return s
}

func (s *DescribeAntiBruteForceRulesRequest) SetName(v string) *DescribeAntiBruteForceRulesRequest {
	s.Name = &v
	return s
}

func (s *DescribeAntiBruteForceRulesRequest) SetPageSize(v string) *DescribeAntiBruteForceRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAntiBruteForceRulesRequest) SetResourceOwnerId(v int64) *DescribeAntiBruteForceRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAntiBruteForceRulesRequest) SetSourceIp(v string) *DescribeAntiBruteForceRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAntiBruteForceRulesRequest) Validate() error {
	return dara.Validate(s)
}
