// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDBRequest
	GetComment() *string
	SetDBInstanceId(v string) *CreateDBRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CreateDBRequest
	GetDBName() *string
	SetRegionId(v string) *CreateDBRequest
	GetRegionId() *string
}

type CreateDBRequest struct {
	// Database remark information.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name. The name must meet the following requirements:
	//
	// 	- The name can contain lowercase letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or digit.
	//
	// 	- The name can be up to 64 characters in length.
	//
	// >  An underscore (_) is counted as two characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb001
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDBRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBRequest) GoString() string {
	return s.String()
}

func (s *CreateDBRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDBRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBRequest) SetComment(v string) *CreateDBRequest {
	s.Comment = &v
	return s
}

func (s *CreateDBRequest) SetDBInstanceId(v string) *CreateDBRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBRequest) SetDBName(v string) *CreateDBRequest {
	s.DBName = &v
	return s
}

func (s *CreateDBRequest) SetRegionId(v string) *CreateDBRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBRequest) Validate() error {
	return dara.Validate(s)
}
