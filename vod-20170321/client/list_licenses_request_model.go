// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLicensesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListLicensesRequest
	GetAppId() *string
	SetAppName(v string) *ListLicensesRequest
	GetAppName() *string
	SetAppType(v int32) *ListLicensesRequest
	GetAppType() *int32
	SetBusinessType(v string) *ListLicensesRequest
	GetBusinessType() *string
	SetNeedTotalCount(v bool) *ListLicensesRequest
	GetNeedTotalCount() *bool
	SetOffset(v int64) *ListLicensesRequest
	GetOffset() *int64
	SetOrders(v []*ListLicensesRequestOrders) *ListLicensesRequest
	GetOrders() []*ListLicensesRequestOrders
	SetPageNo(v int64) *ListLicensesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListLicensesRequest
	GetPageSize() *int64
	SetPkgName(v string) *ListLicensesRequest
	GetPkgName() *string
	SetPlatformType(v int64) *ListLicensesRequest
	GetPlatformType() *int64
	SetProduct(v string) *ListLicensesRequest
	GetProduct() *string
	SetUserId(v int64) *ListLicensesRequest
	GetUserId() *int64
}

type ListLicensesRequest struct {
	AppId          *string                      `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string                      `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType        *int32                       `json:"AppType,omitempty" xml:"AppType,omitempty"`
	BusinessType   *string                      `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	NeedTotalCount *bool                        `json:"NeedTotalCount,omitempty" xml:"NeedTotalCount,omitempty"`
	Offset         *int64                       `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Orders         []*ListLicensesRequestOrders `json:"Orders,omitempty" xml:"Orders,omitempty" type:"Repeated"`
	PageNo         *int64                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PkgName        *string                      `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	PlatformType   *int64                       `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	Product        *string                      `json:"Product,omitempty" xml:"Product,omitempty"`
	UserId         *int64                       `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListLicensesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLicensesRequest) GoString() string {
	return s.String()
}

func (s *ListLicensesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListLicensesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListLicensesRequest) GetAppType() *int32 {
	return s.AppType
}

func (s *ListLicensesRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListLicensesRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ListLicensesRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ListLicensesRequest) GetOrders() []*ListLicensesRequestOrders {
	return s.Orders
}

func (s *ListLicensesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLicensesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLicensesRequest) GetPkgName() *string {
	return s.PkgName
}

func (s *ListLicensesRequest) GetPlatformType() *int64 {
	return s.PlatformType
}

func (s *ListLicensesRequest) GetProduct() *string {
	return s.Product
}

func (s *ListLicensesRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *ListLicensesRequest) SetAppId(v string) *ListLicensesRequest {
	s.AppId = &v
	return s
}

func (s *ListLicensesRequest) SetAppName(v string) *ListLicensesRequest {
	s.AppName = &v
	return s
}

func (s *ListLicensesRequest) SetAppType(v int32) *ListLicensesRequest {
	s.AppType = &v
	return s
}

func (s *ListLicensesRequest) SetBusinessType(v string) *ListLicensesRequest {
	s.BusinessType = &v
	return s
}

func (s *ListLicensesRequest) SetNeedTotalCount(v bool) *ListLicensesRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ListLicensesRequest) SetOffset(v int64) *ListLicensesRequest {
	s.Offset = &v
	return s
}

func (s *ListLicensesRequest) SetOrders(v []*ListLicensesRequestOrders) *ListLicensesRequest {
	s.Orders = v
	return s
}

func (s *ListLicensesRequest) SetPageNo(v int64) *ListLicensesRequest {
	s.PageNo = &v
	return s
}

func (s *ListLicensesRequest) SetPageSize(v int64) *ListLicensesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLicensesRequest) SetPkgName(v string) *ListLicensesRequest {
	s.PkgName = &v
	return s
}

func (s *ListLicensesRequest) SetPlatformType(v int64) *ListLicensesRequest {
	s.PlatformType = &v
	return s
}

func (s *ListLicensesRequest) SetProduct(v string) *ListLicensesRequest {
	s.Product = &v
	return s
}

func (s *ListLicensesRequest) SetUserId(v int64) *ListLicensesRequest {
	s.UserId = &v
	return s
}

func (s *ListLicensesRequest) Validate() error {
	return dara.Validate(s)
}

type ListLicensesRequestOrders struct {
	OrderBy        *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
}

func (s ListLicensesRequestOrders) String() string {
	return dara.Prettify(s)
}

func (s ListLicensesRequestOrders) GoString() string {
	return s.String()
}

func (s *ListLicensesRequestOrders) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListLicensesRequestOrders) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListLicensesRequestOrders) SetOrderBy(v string) *ListLicensesRequestOrders {
	s.OrderBy = &v
	return s
}

func (s *ListLicensesRequestOrders) SetOrderDirection(v string) *ListLicensesRequestOrders {
	s.OrderDirection = &v
	return s
}

func (s *ListLicensesRequestOrders) Validate() error {
	return dara.Validate(s)
}
