// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSQLPlanRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeSQLPlanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLPlanRequest
	GetOwnerId() *int64
	SetProcessId(v string) *DescribeSQLPlanRequest
	GetProcessId() *string
	SetResourceOwnerAccount(v string) *DescribeSQLPlanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLPlanRequest
	GetResourceOwnerId() *int64
}

type DescribeSQLPlanRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-****************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The query ID.
	//
	// >  You can call the [DescribeProcessList](https://help.aliyun.com/document_detail/612277.html) operation to query the IDs of queries that are being executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202105271604431720161662490345*******
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSQLPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSQLPlanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLPlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLPlanRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeSQLPlanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLPlanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLPlanRequest) SetDBClusterId(v string) *DescribeSQLPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetOwnerAccount(v string) *DescribeSQLPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetOwnerId(v int64) *DescribeSQLPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetProcessId(v string) *DescribeSQLPlanRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetResourceOwnerAccount(v string) *DescribeSQLPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanRequest) SetResourceOwnerId(v int64) *DescribeSQLPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLPlanRequest) Validate() error {
	return dara.Validate(s)
}
