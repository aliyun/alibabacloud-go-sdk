// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVirtualNumberRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpNameList(v string) *AddVirtualNumberRelationRequest
	GetCorpNameList() *string
	SetNumberList(v string) *AddVirtualNumberRelationRequest
	GetNumberList() *string
	SetOwnerId(v int64) *AddVirtualNumberRelationRequest
	GetOwnerId() *int64
	SetPhoneNum(v string) *AddVirtualNumberRelationRequest
	GetPhoneNum() *string
	SetProdCode(v string) *AddVirtualNumberRelationRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *AddVirtualNumberRelationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddVirtualNumberRelationRequest
	GetResourceOwnerId() *int64
	SetRouteType(v int32) *AddVirtualNumberRelationRequest
	GetRouteType() *int32
}

type AddVirtualNumberRelationRequest struct {
	// The company names. Separate multiple company names with commas (,).
	//
	// example:
	//
	// Company 1
	CorpNameList *string `json:"CorpNameList,omitempty" xml:"CorpNameList,omitempty"`
	// The real numbers. Separate multiple real numbers with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1321111****,1322222****
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The virtual number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05516214****
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 0
	RouteType *int32 `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s AddVirtualNumberRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddVirtualNumberRelationRequest) GoString() string {
	return s.String()
}

func (s *AddVirtualNumberRelationRequest) GetCorpNameList() *string {
	return s.CorpNameList
}

func (s *AddVirtualNumberRelationRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *AddVirtualNumberRelationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddVirtualNumberRelationRequest) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *AddVirtualNumberRelationRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *AddVirtualNumberRelationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddVirtualNumberRelationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddVirtualNumberRelationRequest) GetRouteType() *int32 {
	return s.RouteType
}

func (s *AddVirtualNumberRelationRequest) SetCorpNameList(v string) *AddVirtualNumberRelationRequest {
	s.CorpNameList = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetNumberList(v string) *AddVirtualNumberRelationRequest {
	s.NumberList = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetOwnerId(v int64) *AddVirtualNumberRelationRequest {
	s.OwnerId = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetPhoneNum(v string) *AddVirtualNumberRelationRequest {
	s.PhoneNum = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetProdCode(v string) *AddVirtualNumberRelationRequest {
	s.ProdCode = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetResourceOwnerAccount(v string) *AddVirtualNumberRelationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetResourceOwnerId(v int64) *AddVirtualNumberRelationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetRouteType(v int32) *AddVirtualNumberRelationRequest {
	s.RouteType = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) Validate() error {
	return dara.Validate(s)
}
