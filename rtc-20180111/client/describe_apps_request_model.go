// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppsRequest
	GetAppId() *string
	SetAppVersion(v string) *DescribeAppsRequest
	GetAppVersion() *string
	SetOrder(v string) *DescribeAppsRequest
	GetOrder() *string
	SetOwnerId(v int64) *DescribeAppsRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeAppsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppsRequest
	GetPageSize() *int32
	SetStatus(v string) *DescribeAppsRequest
	GetStatus() *string
}

type DescribeAppsRequest struct {
	// The application ID. You can create and query application IDs in the console. Only a single ID is supported. If this parameter is left empty, all application IDs are queried.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application version.
	//
	// example:
	//
	// 3.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The sort order. Valid values:
	//
	// - **asc**: Ascending.
	//
	//
	//
	// - **desc**: Descending. Default value.
	//
	// example:
	//
	// asc
	Order   *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The application status. If this parameter is left empty, all statuses are queried. Valid values:
	//
	// - **1**: Available.
	//
	// - **2**: Disabled.
	//
	// - **3**: Disabled due to overdue payments.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppsRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *DescribeAppsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeAppsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAppsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppsRequest) SetAppId(v string) *DescribeAppsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppsRequest) SetAppVersion(v string) *DescribeAppsRequest {
	s.AppVersion = &v
	return s
}

func (s *DescribeAppsRequest) SetOrder(v string) *DescribeAppsRequest {
	s.Order = &v
	return s
}

func (s *DescribeAppsRequest) SetOwnerId(v int64) *DescribeAppsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAppsRequest) SetPageNum(v int32) *DescribeAppsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppsRequest) SetPageSize(v int32) *DescribeAppsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsRequest) SetStatus(v string) *DescribeAppsRequest {
	s.Status = &v
	return s
}

func (s *DescribeAppsRequest) Validate() error {
	return dara.Validate(s)
}
