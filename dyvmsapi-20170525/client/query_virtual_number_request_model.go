// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVirtualNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryVirtualNumberRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryVirtualNumberRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryVirtualNumberRequest
	GetPageSize() *int32
	SetProdCode(v string) *QueryVirtualNumberRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryVirtualNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryVirtualNumberRequest
	GetResourceOwnerId() *int64
	SetRouteType(v int32) *QueryVirtualNumberRequest
	GetRouteType() *int32
}

type QueryVirtualNumberRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The service name. Default value: **dyvms**.
	//
	// example:
	//
	// dyvms
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route type. Valid values:
	//
	// 	- **0**: number location first.
	//
	// 	- **1**: random.
	//
	// example:
	//
	// 0
	RouteType *int32 `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s QueryVirtualNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVirtualNumberRequest) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryVirtualNumberRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryVirtualNumberRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryVirtualNumberRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryVirtualNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryVirtualNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryVirtualNumberRequest) GetRouteType() *int32 {
	return s.RouteType
}

func (s *QueryVirtualNumberRequest) SetOwnerId(v int64) *QueryVirtualNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetPageNo(v int32) *QueryVirtualNumberRequest {
	s.PageNo = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetPageSize(v int32) *QueryVirtualNumberRequest {
	s.PageSize = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetProdCode(v string) *QueryVirtualNumberRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetResourceOwnerAccount(v string) *QueryVirtualNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetResourceOwnerId(v int64) *QueryVirtualNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetRouteType(v int32) *QueryVirtualNumberRequest {
	s.RouteType = &v
	return s
}

func (s *QueryVirtualNumberRequest) Validate() error {
	return dara.Validate(s)
}
