// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAllPrivilegesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAccountAllPrivilegesRequest
	GetAccountName() *string
	SetDBClusterId(v string) *DescribeAccountAllPrivilegesRequest
	GetDBClusterId() *string
	SetMarker(v string) *DescribeAccountAllPrivilegesRequest
	GetMarker() *string
	SetRegionId(v string) *DescribeAccountAllPrivilegesRequest
	GetRegionId() *string
}

type DescribeAccountAllPrivilegesRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// account1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp14t95lun0w****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies the start position marker from which to return results. If you receive a response indicating that the results are truncated, set this parameter to the value of the `Marker` parameter in the response that you received.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAccountAllPrivilegesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAllPrivilegesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountAllPrivilegesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAccountAllPrivilegesRequest) GetMarker() *string {
	return s.Marker
}

func (s *DescribeAccountAllPrivilegesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountAllPrivilegesRequest) SetAccountName(v string) *DescribeAccountAllPrivilegesRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountAllPrivilegesRequest) SetDBClusterId(v string) *DescribeAccountAllPrivilegesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountAllPrivilegesRequest) SetMarker(v string) *DescribeAccountAllPrivilegesRequest {
	s.Marker = &v
	return s
}

func (s *DescribeAccountAllPrivilegesRequest) SetRegionId(v string) *DescribeAccountAllPrivilegesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountAllPrivilegesRequest) Validate() error {
	return dara.Validate(s)
}
