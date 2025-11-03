// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskCheckItemResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeRiskCheckItemResultRequest
	GetCurrentPage() *int32
	SetItemId(v int64) *DescribeRiskCheckItemResultRequest
	GetItemId() *int64
	SetLang(v string) *DescribeRiskCheckItemResultRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeRiskCheckItemResultRequest
	GetPageSize() *int32
	SetResourceOwnerId(v int64) *DescribeRiskCheckItemResultRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeRiskCheckItemResultRequest
	GetSourceIp() *string
}

type DescribeRiskCheckItemResultRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the check item.
	//
	// >  For more information about the IDs and details of the check items that can be used in configuration assessment, see [DescribeRiskCheckResult](https://help.aliyun.com/document_detail/113520.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 173.128.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeRiskCheckItemResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckItemResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckItemResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRiskCheckItemResultRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *DescribeRiskCheckItemResultRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskCheckItemResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRiskCheckItemResultRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRiskCheckItemResultRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskCheckItemResultRequest) SetCurrentPage(v int32) *DescribeRiskCheckItemResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetItemId(v int64) *DescribeRiskCheckItemResultRequest {
	s.ItemId = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetLang(v string) *DescribeRiskCheckItemResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetPageSize(v int32) *DescribeRiskCheckItemResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetResourceOwnerId(v int64) *DescribeRiskCheckItemResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetSourceIp(v string) *DescribeRiskCheckItemResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) Validate() error {
	return dara.Validate(s)
}
