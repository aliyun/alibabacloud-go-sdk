// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTablesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeTablesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTablesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeTablesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTablesRequest
	GetResourceOwnerId() *int64
	SetSchemaName(v string) *DescribeTablesRequest
	GetSchemaName() *string
}

type DescribeTablesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTablesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTablesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTablesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTablesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTablesRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeTablesRequest) SetDBClusterId(v string) *DescribeTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesRequest) SetOwnerAccount(v string) *DescribeTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTablesRequest) SetOwnerId(v int64) *DescribeTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTablesRequest) SetResourceOwnerAccount(v string) *DescribeTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTablesRequest) SetResourceOwnerId(v int64) *DescribeTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTablesRequest) SetSchemaName(v string) *DescribeTablesRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablesRequest) Validate() error {
	return dara.Validate(s)
}
