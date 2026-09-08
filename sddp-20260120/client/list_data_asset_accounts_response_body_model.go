// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDataAssetAccountsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListDataAssetAccountsResponseBodyItems) *ListDataAssetAccountsResponseBody
	GetItems() []*ListDataAssetAccountsResponseBodyItems
	SetPageSize(v int32) *ListDataAssetAccountsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataAssetAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDataAssetAccountsResponseBody
	GetTotalCount() *int32
}

type ListDataAssetAccountsResponseBody struct {
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListDataAssetAccountsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataAssetAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAssetAccountsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDataAssetAccountsResponseBody) GetItems() []*ListDataAssetAccountsResponseBodyItems {
	return s.Items
}

func (s *ListDataAssetAccountsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAssetAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAssetAccountsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataAssetAccountsResponseBody) SetCurrentPage(v int32) *ListDataAssetAccountsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListDataAssetAccountsResponseBody) SetItems(v []*ListDataAssetAccountsResponseBodyItems) *ListDataAssetAccountsResponseBody {
	s.Items = v
	return s
}

func (s *ListDataAssetAccountsResponseBody) SetPageSize(v int32) *ListDataAssetAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataAssetAccountsResponseBody) SetRequestId(v string) *ListDataAssetAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAssetAccountsResponseBody) SetTotalCount(v int32) *ListDataAssetAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataAssetAccountsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataAssetAccountsResponseBodyItems struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AliUid      *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AuthRole    *string `json:"AuthRole,omitempty" xml:"AuthRole,omitempty"`
	// example:
	//
	// client_key
	EncryptionKeyMode *string `json:"EncryptionKeyMode,omitempty" xml:"EncryptionKeyMode,omitempty"`
	EngineType        *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// 2145953410000
	ExpireTime  *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductId   *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s ListDataAssetAccountsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetAccountsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListDataAssetAccountsResponseBodyItems) GetAccountName() *string {
	return s.AccountName
}

func (s *ListDataAssetAccountsResponseBodyItems) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListDataAssetAccountsResponseBodyItems) GetAuthRole() *string {
	return s.AuthRole
}

func (s *ListDataAssetAccountsResponseBodyItems) GetEncryptionKeyMode() *string {
	return s.EncryptionKeyMode
}

func (s *ListDataAssetAccountsResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *ListDataAssetAccountsResponseBodyItems) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListDataAssetAccountsResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *ListDataAssetAccountsResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataAssetAccountsResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDataAssetAccountsResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListDataAssetAccountsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataAssetAccountsResponseBodyItems) GetRegionName() *string {
	return s.RegionName
}

func (s *ListDataAssetAccountsResponseBodyItems) SetAccountName(v string) *ListDataAssetAccountsResponseBodyItems {
	s.AccountName = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetAliUid(v int64) *ListDataAssetAccountsResponseBodyItems {
	s.AliUid = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetAuthRole(v string) *ListDataAssetAccountsResponseBodyItems {
	s.AuthRole = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetEncryptionKeyMode(v string) *ListDataAssetAccountsResponseBodyItems {
	s.EncryptionKeyMode = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetEngineType(v string) *ListDataAssetAccountsResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetExpireTime(v int64) *ListDataAssetAccountsResponseBodyItems {
	s.ExpireTime = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetId(v int64) *ListDataAssetAccountsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetInstanceId(v string) *ListDataAssetAccountsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetProductCode(v string) *ListDataAssetAccountsResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetProductId(v int64) *ListDataAssetAccountsResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetRegionId(v string) *ListDataAssetAccountsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) SetRegionName(v string) *ListDataAssetAccountsResponseBodyItems {
	s.RegionName = &v
	return s
}

func (s *ListDataAssetAccountsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
