// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *AllocateInstancePublicConnectionResponseBody
	GetConnectionString() *string
	SetDbInstanceName(v string) *AllocateInstancePublicConnectionResponseBody
	GetDbInstanceName() *string
	SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody
	GetRequestId() *string
}

type AllocateInstancePublicConnectionResponseBody struct {
	// The endpoint that is used to connect to the database instance.
	//
	// example:
	//
	// test*****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-bp1*****
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateInstancePublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponseBody) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *AllocateInstancePublicConnectionResponseBody) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *AllocateInstancePublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateInstancePublicConnectionResponseBody) SetConnectionString(v string) *AllocateInstancePublicConnectionResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetDbInstanceName(v string) *AllocateInstancePublicConnectionResponseBody {
	s.DbInstanceName = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
