// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartStatisticsPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeSmartStatisticsPageListRequest
	GetCurrentPage() *string
	SetEndDate(v string) *DescribeSmartStatisticsPageListRequest
	GetEndDate() *string
	SetPageSize(v string) *DescribeSmartStatisticsPageListRequest
	GetPageSize() *string
	SetSceneId(v string) *DescribeSmartStatisticsPageListRequest
	GetSceneId() *string
	SetServiceCode(v string) *DescribeSmartStatisticsPageListRequest
	GetServiceCode() *string
	SetStartDate(v string) *DescribeSmartStatisticsPageListRequest
	GetStartDate() *string
}

type DescribeSmartStatisticsPageListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-16 23:59:59 +0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 36**01
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// cloudauthst
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-11-01 00:00:00 +0800
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeSmartStatisticsPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartStatisticsPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmartStatisticsPageListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSmartStatisticsPageListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeSmartStatisticsPageListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSmartStatisticsPageListRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeSmartStatisticsPageListRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeSmartStatisticsPageListRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeSmartStatisticsPageListRequest) SetCurrentPage(v string) *DescribeSmartStatisticsPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetEndDate(v string) *DescribeSmartStatisticsPageListRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetPageSize(v string) *DescribeSmartStatisticsPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetSceneId(v string) *DescribeSmartStatisticsPageListRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetServiceCode(v string) *DescribeSmartStatisticsPageListRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) SetStartDate(v string) *DescribeSmartStatisticsPageListRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeSmartStatisticsPageListRequest) Validate() error {
	return dara.Validate(s)
}
