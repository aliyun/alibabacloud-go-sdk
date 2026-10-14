// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ListDataAssetAccountsRequest
	GetAccountName() *string
	SetAuthRole(v string) *ListDataAssetAccountsRequest
	GetAuthRole() *string
	SetBizType(v string) *ListDataAssetAccountsRequest
	GetBizType() *string
	SetCurrentPage(v int32) *ListDataAssetAccountsRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *ListDataAssetAccountsRequest
	GetInstanceId() *string
	SetLang(v string) *ListDataAssetAccountsRequest
	GetLang() *string
	SetPageSize(v int32) *ListDataAssetAccountsRequest
	GetPageSize() *int32
	SetProductCode(v string) *ListDataAssetAccountsRequest
	GetProductCode() *string
	SetProductIds(v string) *ListDataAssetAccountsRequest
	GetProductIds() *string
}

type ListDataAssetAccountsRequest struct {
	// The database account name. Filters by exact account name.
	//
	// example:
	//
	// app_reader
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The column encryption access permissions. Valid values:
	//
	// - fullAccess: plaintext access.
	//
	// - restrictedAccess: restricted access.
	//
	// - noneAccess: decryption prohibited. noneAccess applies only to ApsaraDB RDS for MySQL instances whose current effective key pattern is client_key.
	//
	// example:
	//
	// fullAccess
	AuthRole *string `json:"AuthRole,omitempty" xml:"AuthRole,omitempty"`
	// The query business scenario. Use DataMasking for the column encryption scenario. This value is case-sensitive.
	//
	// example:
	//
	// DataMasking
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The page number of the current query. Pages start from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The database instance identity. You can obtain this value from Items[].InstanceId returned by ListDataMaskingInstances.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of records per page. Keep other filter conditions consistent when querying the next page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The cloud product code. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also specified, the two values must correspond to each other.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The list of cloud product type identifiers. Separate multiple values with commas (,). Valid values for the column encryption scenario: 5 (RDS), 8 (POLARDB), and 27 (PolarDBX2).
	//
	// example:
	//
	// 5
	ProductIds *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
}

func (s ListDataAssetAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListDataAssetAccountsRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ListDataAssetAccountsRequest) GetAuthRole() *string {
	return s.AuthRole
}

func (s *ListDataAssetAccountsRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListDataAssetAccountsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDataAssetAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataAssetAccountsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataAssetAccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAssetAccountsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDataAssetAccountsRequest) GetProductIds() *string {
	return s.ProductIds
}

func (s *ListDataAssetAccountsRequest) SetAccountName(v string) *ListDataAssetAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *ListDataAssetAccountsRequest) SetAuthRole(v string) *ListDataAssetAccountsRequest {
	s.AuthRole = &v
	return s
}

func (s *ListDataAssetAccountsRequest) SetBizType(v string) *ListDataAssetAccountsRequest {
	s.BizType = &v
	return s
}

func (s *ListDataAssetAccountsRequest) SetCurrentPage(v int32) *ListDataAssetAccountsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDataAssetAccountsRequest) SetInstanceId(v string) *ListDataAssetAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataAssetAccountsRequest) SetLang(v string) *ListDataAssetAccountsRequest {
	s.Lang = &v
	return s
}

func (s *ListDataAssetAccountsRequest) SetPageSize(v int32) *ListDataAssetAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAssetAccountsRequest) SetProductCode(v string) *ListDataAssetAccountsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDataAssetAccountsRequest) SetProductIds(v string) *ListDataAssetAccountsRequest {
	s.ProductIds = &v
	return s
}

func (s *ListDataAssetAccountsRequest) Validate() error {
	return dara.Validate(s)
}
