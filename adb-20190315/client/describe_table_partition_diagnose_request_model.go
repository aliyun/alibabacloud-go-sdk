// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablePartitionDiagnoseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTablePartitionDiagnoseRequest
	GetDBClusterId() *string
	SetLang(v string) *DescribeTablePartitionDiagnoseRequest
	GetLang() *string
	SetOrder(v string) *DescribeTablePartitionDiagnoseRequest
	GetOrder() *string
	SetOwnerAccount(v string) *DescribeTablePartitionDiagnoseRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTablePartitionDiagnoseRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTablePartitionDiagnoseRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTablePartitionDiagnoseRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTablePartitionDiagnoseRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTablePartitionDiagnoseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTablePartitionDiagnoseRequest
	GetResourceOwnerId() *int64
}

type DescribeTablePartitionDiagnoseRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order by which to sort query results. Specify the parameter value in the JSON string format. Example: `[{"Field":"TotalSize","Type":"Desc"}]`.
	//
	// 	- `Field` specifies the field by which to sort the query results. Valid values:
	//
	//     	- `SchemaName`: the name of the database to which the table belongs.
	//
	//     	- `TableName`: the name of the table.
	//
	//     	- `TotalSize`: the total data size of the table.
	//
	//     	- `SpaceRatio`: the storage percentage of the table.
	//
	// 	- `Type` specifies the sorting order. Valid values:
	//
	//     	- `Asc`: ascending order.
	//
	//     	- `Desc`: descending order.
	//
	// >  If you do not specify this parameter, the query results are sorted by the TotalSize field in descending order.
	//
	// example:
	//
	// [{\\"Field\\":\\"TotalSize\\",\\"Type\\":\\"Desc\\"}]
	Order        *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- 30
	//
	// 	- 50
	//
	// 	- 100
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeTablePartitionDiagnoseRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTablePartitionDiagnoseRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTablePartitionDiagnoseRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeTablePartitionDiagnoseRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTablePartitionDiagnoseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTablePartitionDiagnoseRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTablePartitionDiagnoseRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTablePartitionDiagnoseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTablePartitionDiagnoseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTablePartitionDiagnoseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTablePartitionDiagnoseRequest) SetDBClusterId(v string) *DescribeTablePartitionDiagnoseRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetLang(v string) *DescribeTablePartitionDiagnoseRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetOrder(v string) *DescribeTablePartitionDiagnoseRequest {
	s.Order = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetOwnerAccount(v string) *DescribeTablePartitionDiagnoseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetOwnerId(v int64) *DescribeTablePartitionDiagnoseRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetPageNumber(v int32) *DescribeTablePartitionDiagnoseRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetPageSize(v int32) *DescribeTablePartitionDiagnoseRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetRegionId(v string) *DescribeTablePartitionDiagnoseRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetResourceOwnerAccount(v string) *DescribeTablePartitionDiagnoseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) SetResourceOwnerId(v int64) *DescribeTablePartitionDiagnoseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseRequest) Validate() error {
	return dara.Validate(s)
}
