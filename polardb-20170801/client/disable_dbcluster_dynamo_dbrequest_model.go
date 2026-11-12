// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterDynamoDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DisableDBClusterDynamoDBRequest
	GetDBClusterId() *string
}

type DisableDBClusterDynamoDBRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query all clusters in the target region, including cluster IDs. Usage notes: Status prerequisite: The DBClusterStatus of the cluster must be Running. After you call EnableDBClusterDynamoDB, the cluster enters the ConfigSwitching transitional state. You must poll DescribeDBClusterAttribute and wait for the cluster to recover. Usage notes: You must first call EnableDBClusterDynamoDB to enable the DynamoDB compatibility mode. Only PostgreSQL clusters are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DisableDBClusterDynamoDBRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterDynamoDBRequest) GoString() string {
	return s.String()
}

func (s *DisableDBClusterDynamoDBRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DisableDBClusterDynamoDBRequest) SetDBClusterId(v string) *DisableDBClusterDynamoDBRequest {
	s.DBClusterId = &v
	return s
}

func (s *DisableDBClusterDynamoDBRequest) Validate() error {
	return dara.Validate(s)
}
