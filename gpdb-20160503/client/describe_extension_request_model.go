// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeExtensionRequest
	GetDBInstanceId() *string
	SetDatabaseName(v string) *DescribeExtensionRequest
	GetDatabaseName() *string
	SetExtensionName(v string) *DescribeExtensionRequest
	GetExtensionName() *string
}

type DescribeExtensionRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) Interface to query the details of all AnalyticDB PostgreSQL Instances in the target region, including Instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Database name.
	//
	// 	- Only contain letters, digits, and underscores (_).
	//
	// 	- Must start with a letter.
	//
	// 	- Up to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The extension name.
	//
	// This parameter is required.
	//
	// example:
	//
	// zhparser
	ExtensionName *string `json:"ExtensionName,omitempty" xml:"ExtensionName,omitempty"`
}

func (s DescribeExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExtensionRequest) GoString() string {
	return s.String()
}

func (s *DescribeExtensionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeExtensionRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeExtensionRequest) GetExtensionName() *string {
	return s.ExtensionName
}

func (s *DescribeExtensionRequest) SetDBInstanceId(v string) *DescribeExtensionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeExtensionRequest) SetDatabaseName(v string) *DescribeExtensionRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeExtensionRequest) SetExtensionName(v string) *DescribeExtensionRequest {
	s.ExtensionName = &v
	return s
}

func (s *DescribeExtensionRequest) Validate() error {
	return dara.Validate(s)
}
