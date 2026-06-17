// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIServiceSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyAIServiceSecurityIpsRequest
	GetDBInstanceId() *string
	SetSecurityIPList(v string) *ModifyAIServiceSecurityIpsRequest
	GetSecurityIPList() *string
	SetServiceId(v string) *ModifyAIServiceSecurityIpsRequest
	GetServiceId() *string
	SetType(v string) *ModifyAIServiceSecurityIpsRequest
	GetType() *string
}

type ModifyAIServiceSecurityIpsRequest struct {
	// The ID of the instance.
	//
	// > To view details of all instances in a destination region, including their IDs, call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A comma-separated list of IP addresses or CIDR blocks in the IP address whitelist group. You can specify up to 1000 entries. To block all external IP addresses, set this parameter to 127.0.0.1. Valid formats include the following:
	//
	// - 10.23.12.24 (an IPv4 address)
	//
	// - 10.23.12.24/24 (a CIDR block. The number after the slash indicates the prefix length and must be between 1 and 32.)
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The ID of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service type. Only drama is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// drama
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyAIServiceSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIServiceSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyAIServiceSecurityIpsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAIServiceSecurityIpsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifyAIServiceSecurityIpsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ModifyAIServiceSecurityIpsRequest) GetType() *string {
	return s.Type
}

func (s *ModifyAIServiceSecurityIpsRequest) SetDBInstanceId(v string) *ModifyAIServiceSecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAIServiceSecurityIpsRequest) SetSecurityIPList(v string) *ModifyAIServiceSecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifyAIServiceSecurityIpsRequest) SetServiceId(v string) *ModifyAIServiceSecurityIpsRequest {
	s.ServiceId = &v
	return s
}

func (s *ModifyAIServiceSecurityIpsRequest) SetType(v string) *ModifyAIServiceSecurityIpsRequest {
	s.Type = &v
	return s
}

func (s *ModifyAIServiceSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
