// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedInstanceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeExposedInstanceDetailRequest
	GetLang() *string
	SetResourceDirectoryAccountId(v int64) *DescribeExposedInstanceDetailRequest
	GetResourceDirectoryAccountId() *int64
	SetUuid(v string) *DescribeExposedInstanceDetailRequest
	GetUuid() *string
}

type DescribeExposedInstanceDetailRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the account ID.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The UUID of the server or the instance ID of the database that is exposed on the Internet.
	//
	// >  You can call the [DescribeExposedInstanceList](~~DescribeExposedInstanceList~~) operation to query the UUIDs of servers or instance IDs of databases.
	//
	// This parameter is required.
	//
	// example:
	//
	// fc82b966-4d70-4e01-bf4f-aa4076a5****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExposedInstanceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExposedInstanceDetailRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeExposedInstanceDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeExposedInstanceDetailRequest) SetLang(v string) *DescribeExposedInstanceDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExposedInstanceDetailRequest) SetResourceDirectoryAccountId(v int64) *DescribeExposedInstanceDetailRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeExposedInstanceDetailRequest) SetUuid(v string) *DescribeExposedInstanceDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceDetailRequest) Validate() error {
	return dara.Validate(s)
}
