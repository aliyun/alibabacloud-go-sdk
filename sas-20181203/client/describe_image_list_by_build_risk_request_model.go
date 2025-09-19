// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListByBuildRiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeImageListByBuildRiskRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeImageListByBuildRiskRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageListByBuildRiskRequest
	GetPageSize() *int32
	SetRiskKey(v string) *DescribeImageListByBuildRiskRequest
	GetRiskKey() *string
	SetRiskLevel(v string) *DescribeImageListByBuildRiskRequest
	GetRiskLevel() *string
	SetStatus(v int32) *DescribeImageListByBuildRiskRequest
	GetStatus() *int32
}

type DescribeImageListByBuildRiskRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
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
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The key of the risk. You can call the [DescribeImageBuildRiskList](~~~~) operation to obtain the value of RiskKey.
	//
	// example:
	//
	// no_user
	RiskKey *string `json:"RiskKey,omitempty" xml:"RiskKey,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the alert event. Valid values:
	//
	// 	- **0**: unhandled.
	//
	// 	- **1**: ignored.
	//
	// 	- **2**: false positive.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImageListByBuildRiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListByBuildRiskRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageListByBuildRiskRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageListByBuildRiskRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageListByBuildRiskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageListByBuildRiskRequest) GetRiskKey() *string {
	return s.RiskKey
}

func (s *DescribeImageListByBuildRiskRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageListByBuildRiskRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageListByBuildRiskRequest) SetCurrentPage(v int32) *DescribeImageListByBuildRiskRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetLang(v string) *DescribeImageListByBuildRiskRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetPageSize(v int32) *DescribeImageListByBuildRiskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetRiskKey(v string) *DescribeImageListByBuildRiskRequest {
	s.RiskKey = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetRiskLevel(v string) *DescribeImageListByBuildRiskRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetStatus(v int32) *DescribeImageListByBuildRiskRequest {
	s.Status = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) Validate() error {
	return dara.Validate(s)
}
