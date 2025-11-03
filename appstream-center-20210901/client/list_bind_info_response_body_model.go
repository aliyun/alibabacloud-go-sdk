// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBindInfoModels(v []*ListBindInfoResponseBodyBindInfoModels) *ListBindInfoResponseBody
	GetBindInfoModels() []*ListBindInfoResponseBodyBindInfoModels
	SetPageNumber(v int32) *ListBindInfoResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBindInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBindInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBindInfoResponseBody
	GetTotalCount() *int32
}

type ListBindInfoResponseBody struct {
	// The bindings.
	BindInfoModels []*ListBindInfoResponseBodyBindInfoModels `json:"BindInfoModels,omitempty" xml:"BindInfoModels,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AD2D0761-1FE5-549D-B169-D3F8D19C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBindInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBindInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindInfoResponseBody) GetBindInfoModels() []*ListBindInfoResponseBodyBindInfoModels {
	return s.BindInfoModels
}

func (s *ListBindInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBindInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBindInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBindInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBindInfoResponseBody) SetBindInfoModels(v []*ListBindInfoResponseBodyBindInfoModels) *ListBindInfoResponseBody {
	s.BindInfoModels = v
	return s
}

func (s *ListBindInfoResponseBody) SetPageNumber(v int32) *ListBindInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBindInfoResponseBody) SetPageSize(v int32) *ListBindInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBindInfoResponseBody) SetRequestId(v string) *ListBindInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindInfoResponseBody) SetTotalCount(v int32) *ListBindInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBindInfoResponseBody) Validate() error {
	if s.BindInfoModels != nil {
		for _, item := range s.BindInfoModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBindInfoResponseBodyBindInfoModels struct {
	// The account type.
	//
	// Valid values:
	//
	// 	- ad: Active Directory (AD) account
	//
	// 	- simple: convenience account
	//
	// example:
	//
	// simple
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The app ID.
	//
	// example:
	//
	// ca-fq738or6vd854****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-0abxhr6ce35w8****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the app instance.
	//
	// example:
	//
	// ai-83oe276fre4l3****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The app version.
	//
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The product type.
	//
	// Valid values:
	//
	// 	- CloudApp: App Streaming
	//
	// 	- CloudBrowser: Cloud-based Browser
	//
	// 	- AndroidCloud: Cloud Phone
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// Alice
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the Alibaba Cloud Workspace user.
	//
	// example:
	//
	// 2ca6f5a93536****
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s ListBindInfoResponseBodyBindInfoModels) String() string {
	return dara.Prettify(s)
}

func (s ListBindInfoResponseBodyBindInfoModels) GoString() string {
	return s.String()
}

func (s *ListBindInfoResponseBodyBindInfoModels) GetAccountType() *string {
	return s.AccountType
}

func (s *ListBindInfoResponseBodyBindInfoModels) GetAppId() *string {
	return s.AppId
}

func (s *ListBindInfoResponseBodyBindInfoModels) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListBindInfoResponseBodyBindInfoModels) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *ListBindInfoResponseBodyBindInfoModels) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListBindInfoResponseBodyBindInfoModels) GetProductType() *string {
	return s.ProductType
}

func (s *ListBindInfoResponseBodyBindInfoModels) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBindInfoResponseBodyBindInfoModels) GetUserId() *string {
	return s.UserId
}

func (s *ListBindInfoResponseBodyBindInfoModels) GetWyId() *string {
	return s.WyId
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAccountType(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AccountType = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAppId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AppId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAppInstanceGroupId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAppInstanceId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AppInstanceId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAppVersion(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AppVersion = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetProductType(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.ProductType = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetRegionId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.RegionId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetUserId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.UserId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetWyId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.WyId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) Validate() error {
	return dara.Validate(s)
}
