// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVirtualNumberRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryVirtualNumberRelationRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryVirtualNumberRelationRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryVirtualNumberRelationRequest
	GetPageSize() *int32
	SetPhoneNum(v string) *QueryVirtualNumberRelationRequest
	GetPhoneNum() *string
	SetProdCode(v string) *QueryVirtualNumberRelationRequest
	GetProdCode() *string
	SetQualificationId(v int64) *QueryVirtualNumberRelationRequest
	GetQualificationId() *int64
	SetRegionNameCity(v string) *QueryVirtualNumberRelationRequest
	GetRegionNameCity() *string
	SetRelatedNum(v string) *QueryVirtualNumberRelationRequest
	GetRelatedNum() *string
	SetResourceOwnerAccount(v string) *QueryVirtualNumberRelationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryVirtualNumberRelationRequest
	GetResourceOwnerId() *int64
	SetRouteType(v int32) *QueryVirtualNumberRelationRequest
	GetRouteType() *int32
	SetSpecId(v int64) *QueryVirtualNumberRelationRequest
	GetSpecId() *int64
}

type QueryVirtualNumberRelationRequest struct {
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
	// The virtual number.
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
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The qualification ID.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Qualifications\\&Communication Scripts > Qualification Management**, and then click **Details*	- in the Actions column to view the qualification ID.
	//
	// example:
	//
	// 1000000542****
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The city to which the virtual number belongs.
	//
	// example:
	//
	// hangzhou
	RegionNameCity *string `json:"RegionNameCity,omitempty" xml:"RegionNameCity,omitempty"`
	// The real number.
	//
	// example:
	//
	// 1705559****
	RelatedNum           *string `json:"RelatedNum,omitempty" xml:"RelatedNum,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route type. Valid values:
	//
	// **0**: number location first. **1**: random.
	//
	// example:
	//
	// 0
	RouteType *int32 `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The number type. Valid values:
	//
	// 	- **1**: the number provided by a virtual network operator, in the 05710000\\*\\*\\*\\	- format.
	//
	// 	- **2**: the number provided by an Internet service provider (ISP).
	//
	// 	- **3**: a 5-digit phone number that starts with 95.
	//
	// example:
	//
	// 1
	SpecId *int64 `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s QueryVirtualNumberRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVirtualNumberRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberRelationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryVirtualNumberRelationRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryVirtualNumberRelationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryVirtualNumberRelationRequest) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *QueryVirtualNumberRelationRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryVirtualNumberRelationRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *QueryVirtualNumberRelationRequest) GetRegionNameCity() *string {
	return s.RegionNameCity
}

func (s *QueryVirtualNumberRelationRequest) GetRelatedNum() *string {
	return s.RelatedNum
}

func (s *QueryVirtualNumberRelationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryVirtualNumberRelationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryVirtualNumberRelationRequest) GetRouteType() *int32 {
	return s.RouteType
}

func (s *QueryVirtualNumberRelationRequest) GetSpecId() *int64 {
	return s.SpecId
}

func (s *QueryVirtualNumberRelationRequest) SetOwnerId(v int64) *QueryVirtualNumberRelationRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetPageNo(v int32) *QueryVirtualNumberRelationRequest {
	s.PageNo = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetPageSize(v int32) *QueryVirtualNumberRelationRequest {
	s.PageSize = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetPhoneNum(v string) *QueryVirtualNumberRelationRequest {
	s.PhoneNum = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetProdCode(v string) *QueryVirtualNumberRelationRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetQualificationId(v int64) *QueryVirtualNumberRelationRequest {
	s.QualificationId = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetRegionNameCity(v string) *QueryVirtualNumberRelationRequest {
	s.RegionNameCity = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetRelatedNum(v string) *QueryVirtualNumberRelationRequest {
	s.RelatedNum = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetResourceOwnerAccount(v string) *QueryVirtualNumberRelationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetResourceOwnerId(v int64) *QueryVirtualNumberRelationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetRouteType(v int32) *QueryVirtualNumberRelationRequest {
	s.RouteType = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetSpecId(v int64) *QueryVirtualNumberRelationRequest {
	s.SpecId = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) Validate() error {
	return dara.Validate(s)
}
