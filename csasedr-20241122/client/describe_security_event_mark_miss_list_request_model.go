// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventMarkMissListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSecurityEventMarkMissListRequest
	GetCurrentPage() *int32
	SetDisposalWay(v string) *DescribeSecurityEventMarkMissListRequest
	GetDisposalWay() *string
	SetEventName(v string) *DescribeSecurityEventMarkMissListRequest
	GetEventName() *string
	SetMaxResults(v int32) *DescribeSecurityEventMarkMissListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSecurityEventMarkMissListRequest
	GetNextToken() *string
	SetPageSize(v int32) *DescribeSecurityEventMarkMissListRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribeSecurityEventMarkMissListRequest
	GetRemark() *string
}

type DescribeSecurityEventMarkMissListRequest struct {
	// Specifies the page number of the returned page. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The disposal method of the rule. The current query does not filter by this field, and you do not need to set it. Valid values:
	//
	// - **auto_add_white**: Automatic whitelisting.
	//
	// - **defense_not_notification**: Defense only without notification.
	//
	// example:
	//
	// auto_add_white
	DisposalWay *string `json:"DisposalWay,omitempty" xml:"DisposalWay,omitempty"`
	// The name of the alert event (subtype).
	//
	// example:
	//
	// Untrusted process startup
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The maximum number of records to return in a single query. This field is reserved for compatibility. The backend currently uses CurrentPage and PageSize for pagination.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// ****==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of whitelist rules to return on each page in a paging query. Default value: **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// **[Deprecated]*	- This field is deprecated.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeSecurityEventMarkMissListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventMarkMissListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventMarkMissListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSecurityEventMarkMissListRequest) GetDisposalWay() *string {
	return s.DisposalWay
}

func (s *DescribeSecurityEventMarkMissListRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSecurityEventMarkMissListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSecurityEventMarkMissListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSecurityEventMarkMissListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSecurityEventMarkMissListRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeSecurityEventMarkMissListRequest) SetCurrentPage(v int32) *DescribeSecurityEventMarkMissListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetDisposalWay(v string) *DescribeSecurityEventMarkMissListRequest {
	s.DisposalWay = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetEventName(v string) *DescribeSecurityEventMarkMissListRequest {
	s.EventName = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetMaxResults(v int32) *DescribeSecurityEventMarkMissListRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetNextToken(v string) *DescribeSecurityEventMarkMissListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetPageSize(v int32) *DescribeSecurityEventMarkMissListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) SetRemark(v string) *DescribeSecurityEventMarkMissListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListRequest) Validate() error {
	return dara.Validate(s)
}
