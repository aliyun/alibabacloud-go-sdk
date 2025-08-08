// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *ListInstancesRequest
	GetDbType() *string
	SetEnvType(v string) *ListInstancesRequest
	GetEnvType() *string
	SetInstanceSource(v string) *ListInstancesRequest
	GetInstanceSource() *string
	SetInstanceState(v string) *ListInstancesRequest
	GetInstanceState() *string
	SetNetType(v string) *ListInstancesRequest
	GetNetType() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetRealLoginUserUid(v string) *ListInstancesRequest
	GetRealLoginUserUid() *string
	SetRegion(v string) *ListInstancesRequest
	GetRegion() *string
	SetSearchKey(v string) *ListInstancesRequest
	GetSearchKey() *string
	SetTid(v int64) *ListInstancesRequest
	GetTid() *int64
}

type ListInstancesRequest struct {
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database instance belongs. Valid values:
	//
	// 	- **product:*	- production environment
	//
	// 	- **dev**: development environment
	//
	// 	- **pre**: pre-release environment
	//
	// 	- **test**: test environment
	//
	// 	- **sit**: system integration testing (SIT) environment
	//
	// 	- **uat**: user acceptance testing (UAT) environment
	//
	// 	- **pet**: stress testing environment
	//
	// 	- **stag:*	- staging environment
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The source of the database instance. Valid values:
	//
	// 	- **PUBLIC_OWN**: a self-managed database instance that is deployed on the Internet
	//
	// 	- **RDS**: an ApsaraDB RDS instance
	//
	// 	- **ECS_OWN**: a self-managed database that is deployed on an Elastic Compute Service (ECS) instance
	//
	// 	- **VPC_IDC**: a self-managed database instance that is deployed in a data center connected over a virtual private cloud (VPC)
	//
	// example:
	//
	// RDS
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The status of the database instance. Valid values:
	//
	// 	- **NORMAL**
	//
	// 	- **DISABLE**
	//
	// example:
	//
	// NORMAL
	InstanceState *string `json:"InstanceState,omitempty" xml:"InstanceState,omitempty"`
	// The network type of the database instance. Valid values:
	//
	// 	- **CLASSIC:*	- classic network
	//
	// 	- **VPC:*	- VPC
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The number cannot exceed 100.
	//
	// example:
	//
	// 50
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The keyword that is used to search for database instances.
	//
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetDbType() *string {
	return s.DbType
}

func (s *ListInstancesRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListInstancesRequest) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *ListInstancesRequest) GetInstanceState() *string {
	return s.InstanceState
}

func (s *ListInstancesRequest) GetNetType() *string {
	return s.NetType
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *ListInstancesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListInstancesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListInstancesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListInstancesRequest) SetDbType(v string) *ListInstancesRequest {
	s.DbType = &v
	return s
}

func (s *ListInstancesRequest) SetEnvType(v string) *ListInstancesRequest {
	s.EnvType = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceSource(v string) *ListInstancesRequest {
	s.InstanceSource = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceState(v string) *ListInstancesRequest {
	s.InstanceState = &v
	return s
}

func (s *ListInstancesRequest) SetNetType(v string) *ListInstancesRequest {
	s.NetType = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetRealLoginUserUid(v string) *ListInstancesRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *ListInstancesRequest) SetRegion(v string) *ListInstancesRequest {
	s.Region = &v
	return s
}

func (s *ListInstancesRequest) SetSearchKey(v string) *ListInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *ListInstancesRequest) SetTid(v int64) *ListInstancesRequest {
	s.Tid = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
