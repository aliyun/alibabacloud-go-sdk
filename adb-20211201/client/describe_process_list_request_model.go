// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeProcessListRequest
	GetDBClusterId() *string
	SetKeyword(v string) *DescribeProcessListRequest
	GetKeyword() *string
	SetOrder(v string) *DescribeProcessListRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeProcessListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeProcessListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeProcessListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeProcessListRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeProcessListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeProcessListRequest
	GetResourceOwnerId() *int64
	SetRunningTime(v int32) *DescribeProcessListRequest
	GetRunningTime() *int32
	SetShowFull(v bool) *DescribeProcessListRequest
	GetShowFull() *bool
	SetUser(v string) *DescribeProcessListRequest
	GetUser() *string
}

type DescribeProcessListRequest struct {
	// The cluster ID of the AnalyticDB for MySQL Data Lakehouse Edition.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to view details of all AnalyticDB for MySQL Data Lakehouse Edition clusters in the destination region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-xxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Filter keyword. Currently, only **SELECT*	- is supported.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Sort by the specified field in JSON format, `[{"Field":"Time","Type":"Desc" },{ "Field":"User", "Type":"Asc" }]`. Values:
	//
	// - **Field**: The name of the field to sort by. Supports Time, User, Host, and DB fields.
	//
	// - **Type**: Sort type. **Desc*	- for descending order, **Asc*	- for ascending order.
	//
	// example:
	//
	// [ { "Field":"Time","Type":"Desc" },  { "Field":"User", "Type":"Asc" }]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Page number. Must be a positive integer. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of records per page. Values:
	//
	// - **30*	- (Default value)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Filter by running time. Displays queries that have run longer than the specified time. Unit: seconds.
	//
	// example:
	//
	// 20
	RunningTime *int32 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// Specifies whether to display the full SQL statement. Values:
	//
	// - **True**: Displays the full SQL statement.
	//
	// - **False**: Displays only the first 100 characters of the SQL statement.
	//
	// > Default value: False.
	//
	// example:
	//
	// True
	ShowFull *bool `json:"ShowFull,omitempty" xml:"ShowFull,omitempty"`
	// The database account.
	//
	// example:
	//
	// admin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeProcessListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeProcessListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeProcessListRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeProcessListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeProcessListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeProcessListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeProcessListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProcessListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeProcessListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeProcessListRequest) GetRunningTime() *int32 {
	return s.RunningTime
}

func (s *DescribeProcessListRequest) GetShowFull() *bool {
	return s.ShowFull
}

func (s *DescribeProcessListRequest) GetUser() *string {
	return s.User
}

func (s *DescribeProcessListRequest) SetDBClusterId(v string) *DescribeProcessListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeProcessListRequest) SetKeyword(v string) *DescribeProcessListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeProcessListRequest) SetOrder(v string) *DescribeProcessListRequest {
	s.Order = &v
	return s
}

func (s *DescribeProcessListRequest) SetOwnerAccount(v string) *DescribeProcessListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeProcessListRequest) SetOwnerId(v int64) *DescribeProcessListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageNumber(v int32) *DescribeProcessListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProcessListRequest) SetPageSize(v int32) *DescribeProcessListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProcessListRequest) SetResourceOwnerAccount(v string) *DescribeProcessListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeProcessListRequest) SetResourceOwnerId(v int64) *DescribeProcessListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeProcessListRequest) SetRunningTime(v int32) *DescribeProcessListRequest {
	s.RunningTime = &v
	return s
}

func (s *DescribeProcessListRequest) SetShowFull(v bool) *DescribeProcessListRequest {
	s.ShowFull = &v
	return s
}

func (s *DescribeProcessListRequest) SetUser(v string) *DescribeProcessListRequest {
	s.User = &v
	return s
}

func (s *DescribeProcessListRequest) Validate() error {
	return dara.Validate(s)
}
