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
	// The page number of the returned data.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of records returned in the current query.
	Items []*ListDataAssetAccountsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page in the current query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID, which can be used to trace the current API call.
	//
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the current statistics or query conditions.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The database account name.
	//
	// example:
	//
	// app_reader
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the Alibaba Cloud account to which the asset belongs.
	//
	// example:
	//
	// 123456789012****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The authentication database of the MongoDB account. This parameter must be passed together with AccountName when modifying permissions.
	//
	// example:
	//
	// admin
	AuthDb *string `json:"AuthDb,omitempty" xml:"AuthDb,omitempty"`
	// The column encryption access permissions. Valid values:
	//
	// - fullAccess: plaintext access
	//
	// - restrictedAccess: restricted access
	//
	// - noneAccess: decryption prohibited. This value applies only to ApsaraDB RDS for MySQL instances whose current effective key mode is client_key.
	//
	// example:
	//
	// fullAccess
	AuthRole *string `json:"AuthRole,omitempty" xml:"AuthRole,omitempty"`
	// The current effective column encryption key mode. Valid values:
	//
	// - client_key: local key
	//
	// - kms_key: KMS key
	//
	// This parameter is not returned if no effective configuration exists.
	//
	// example:
	//
	// client_key
	EncryptionKeyMode *string `json:"EncryptionKeyMode,omitempty" xml:"EncryptionKeyMode,omitempty"`
	// The database engine type.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The expiration time of the plaintext access permissions. The value is a UNIX timestamp that represents the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. This parameter is not returned if no validity period is set.
	//
	// example:
	//
	// 2145953410000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The database account ID. This ID is used as AccountId when configuring permissions.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The cloud service code.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud service type ID. Valid values: 5 (RDS), 8 (POLARDB), and 27 (PolarDBX2).
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region name of the asset.
	//
	// example:
	//
	// China (Zhangjiakou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
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

func (s *ListDataAssetAccountsResponseBodyItems) GetAuthDb() *string {
	return s.AuthDb
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

func (s *ListDataAssetAccountsResponseBodyItems) SetAuthDb(v string) *ListDataAssetAccountsResponseBodyItems {
	s.AuthDb = &v
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
