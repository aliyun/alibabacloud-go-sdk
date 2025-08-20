// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnabledPrivilegesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeEnabledPrivilegesRequest
	GetAccountName() *string
	SetDBClusterId(v string) *DescribeEnabledPrivilegesRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeEnabledPrivilegesRequest
	GetRegionId() *string
}

type DescribeEnabledPrivilegesRequest struct {
	// The name of the database account.
	//
	// >  You can call the [DescribeAccounts](https://help.aliyun.com/document_detail/612430.html) operation to query the information about database accounts for a cluster, including the account name.
	//
	// example:
	//
	// test_accout
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp14t95lun0w****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEnabledPrivilegesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnabledPrivilegesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnabledPrivilegesRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeEnabledPrivilegesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeEnabledPrivilegesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnabledPrivilegesRequest) SetAccountName(v string) *DescribeEnabledPrivilegesRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeEnabledPrivilegesRequest) SetDBClusterId(v string) *DescribeEnabledPrivilegesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEnabledPrivilegesRequest) SetRegionId(v string) *DescribeEnabledPrivilegesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnabledPrivilegesRequest) Validate() error {
	return dara.Validate(s)
}
