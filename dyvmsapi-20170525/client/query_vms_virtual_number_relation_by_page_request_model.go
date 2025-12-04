// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVmsVirtualNumberRelationByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumberCity(v string) *QueryVmsVirtualNumberRelationByPageRequest
	GetNumberCity() *string
	SetNumberProvince(v string) *QueryVmsVirtualNumberRelationByPageRequest
	GetNumberProvince() *string
	SetOwnerId(v int64) *QueryVmsVirtualNumberRelationByPageRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QueryVmsVirtualNumberRelationByPageRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryVmsVirtualNumberRelationByPageRequest
	GetPageSize() *int64
	SetRealNumber(v string) *QueryVmsVirtualNumberRelationByPageRequest
	GetRealNumber() *string
	SetResourceOwnerAccount(v string) *QueryVmsVirtualNumberRelationByPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryVmsVirtualNumberRelationByPageRequest
	GetResourceOwnerId() *int64
	SetState(v int64) *QueryVmsVirtualNumberRelationByPageRequest
	GetState() *int64
	SetVirtualNumber(v string) *QueryVmsVirtualNumberRelationByPageRequest
	GetVirtualNumber() *string
}

type QueryVmsVirtualNumberRelationByPageRequest struct {
	// 号码归属地--城市
	//
	// example:
	//
	// 示例值示例值
	NumberCity *string `json:"NumberCity,omitempty" xml:"NumberCity,omitempty"`
	// 号码归属地--省
	//
	// example:
	//
	// 示例值示例值
	NumberProvince *string `json:"NumberProvince,omitempty" xml:"NumberProvince,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 74
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 81
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 真实号码
	//
	// example:
	//
	// 131***1234
	RealNumber           *string `json:"RealNumber,omitempty" xml:"RealNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 状态 1:有效；0:无效；-1:注销
	//
	// example:
	//
	// 1
	State *int64 `json:"State,omitempty" xml:"State,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// 0571****1234
	VirtualNumber *string `json:"VirtualNumber,omitempty" xml:"VirtualNumber,omitempty"`
}

func (s QueryVmsVirtualNumberRelationByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVmsVirtualNumberRelationByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetNumberCity() *string {
	return s.NumberCity
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetNumberProvince() *string {
	return s.NumberProvince
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetRealNumber() *string {
	return s.RealNumber
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetState() *int64 {
	return s.State
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) GetVirtualNumber() *string {
	return s.VirtualNumber
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetNumberCity(v string) *QueryVmsVirtualNumberRelationByPageRequest {
	s.NumberCity = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetNumberProvince(v string) *QueryVmsVirtualNumberRelationByPageRequest {
	s.NumberProvince = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetOwnerId(v int64) *QueryVmsVirtualNumberRelationByPageRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetPageNo(v int64) *QueryVmsVirtualNumberRelationByPageRequest {
	s.PageNo = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetPageSize(v int64) *QueryVmsVirtualNumberRelationByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetRealNumber(v string) *QueryVmsVirtualNumberRelationByPageRequest {
	s.RealNumber = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetResourceOwnerAccount(v string) *QueryVmsVirtualNumberRelationByPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetResourceOwnerId(v int64) *QueryVmsVirtualNumberRelationByPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetState(v int64) *QueryVmsVirtualNumberRelationByPageRequest {
	s.State = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) SetVirtualNumber(v string) *QueryVmsVirtualNumberRelationByPageRequest {
	s.VirtualNumber = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageRequest) Validate() error {
	return dara.Validate(s)
}
