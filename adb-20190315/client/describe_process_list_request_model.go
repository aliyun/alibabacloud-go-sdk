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
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The keyword in an SQL statement, which is used to filter queries. Set the value to **SELECT**.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The order in which queries are sorted based on the specified fields. Specify this parameter as an ordered JSON array in the `[{"Field":"Time","Type":"Desc" },{ "Field":"User", "Type":"Asc" }]` format.
	//
	// 	- **Field*	- specifies the field used to sort queries. Valid values: Time, User, Host, and DB.
	//
	// 	- **Type*	- specifies the sorting sequence. Valid values: **Desc*	- and **Asc**.
	//
	// example:
	//
	// [ { "Field":"Time","Type":"Desc" },  { "Field":"User", "Type":"Asc" }]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The execution duration used to filter queries. Queries that take a longer time than the specified execution duration are displayed. Unit: seconds.
	//
	// example:
	//
	// 5
	RunningTime *int32 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// Specifies whether to show a complete SQL statement. Valid values:
	//
	// 	- **True**: shows a complete SQL statement.
	//
	// 	- **False**: shows only the first 100 characters of an SQL statement.
	//
	// >  The default value is False.
	//
	// example:
	//
	// True
	ShowFull *bool `json:"ShowFull,omitempty" xml:"ShowFull,omitempty"`
	// The name of the user used to filter queries.
	//
	// example:
	//
	// test
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
