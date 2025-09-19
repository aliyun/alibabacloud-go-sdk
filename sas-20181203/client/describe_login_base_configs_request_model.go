// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoginBaseConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeLoginBaseConfigsRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeLoginBaseConfigsRequest
	GetPageSize() *int32
	SetTarget(v string) *DescribeLoginBaseConfigsRequest
	GetTarget() *string
	SetType(v string) *DescribeLoginBaseConfigsRequest
	GetType() *string
}

type DescribeLoginBaseConfigsRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: **5**.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The server to which the configuration is applied. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **Target**: the UUID or group ID of the server to add or delete.
	//
	// > If targetType is set to uuid, the value of Target is the UUID of the server. If targetType is set to groupId, the value of Target is the group ID of the server. If targetType is set to global, the value of Target is a hyphen (-).
	//
	// 	- **targetType**: the type of the server to which the configuration is applied. Valid values:
	//
	//     	- **uuid**: a server
	//
	//     	- **groupId**: a server group
	//
	//     	- **global**: all servers
	//
	// example:
	//
	// [ {"target": "inet-7c676676-06fa-442e-90fb-b802e5d6****", "targetType": "uuid" } ]
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The logon type of the configuration to query. Valid values:
	//
	// 	- **login_common_location**: common logon location
	//
	// 	- **login_common_ip**: common logon IP address
	//
	// 	- **login_common_time**: common logon time range
	//
	// 	- **login_common_account**: common logon account
	//
	// This parameter is required.
	//
	// example:
	//
	// login_common_location
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLoginBaseConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginBaseConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoginBaseConfigsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeLoginBaseConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoginBaseConfigsRequest) GetTarget() *string {
	return s.Target
}

func (s *DescribeLoginBaseConfigsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeLoginBaseConfigsRequest) SetCurrentPage(v int32) *DescribeLoginBaseConfigsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeLoginBaseConfigsRequest) SetPageSize(v int32) *DescribeLoginBaseConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoginBaseConfigsRequest) SetTarget(v string) *DescribeLoginBaseConfigsRequest {
	s.Target = &v
	return s
}

func (s *DescribeLoginBaseConfigsRequest) SetType(v string) *DescribeLoginBaseConfigsRequest {
	s.Type = &v
	return s
}

func (s *DescribeLoginBaseConfigsRequest) Validate() error {
	return dara.Validate(s)
}
