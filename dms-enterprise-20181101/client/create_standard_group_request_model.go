// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *CreateStandardGroupRequest
	GetDbType() *string
	SetDescription(v string) *CreateStandardGroupRequest
	GetDescription() *string
	SetGroupName(v string) *CreateStandardGroupRequest
	GetGroupName() *string
	SetTid(v int64) *CreateStandardGroupRequest
	GetTid() *int64
}

type CreateStandardGroupRequest struct {
	// The type of the database engine. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The description of the security rule set.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the security rule set.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 3000
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateStandardGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardGroupRequest) GetDbType() *string {
	return s.DbType
}

func (s *CreateStandardGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateStandardGroupRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateStandardGroupRequest) SetDbType(v string) *CreateStandardGroupRequest {
	s.DbType = &v
	return s
}

func (s *CreateStandardGroupRequest) SetDescription(v string) *CreateStandardGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateStandardGroupRequest) SetGroupName(v string) *CreateStandardGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateStandardGroupRequest) SetTid(v int64) *CreateStandardGroupRequest {
	s.Tid = &v
	return s
}

func (s *CreateStandardGroupRequest) Validate() error {
	return dara.Validate(s)
}
