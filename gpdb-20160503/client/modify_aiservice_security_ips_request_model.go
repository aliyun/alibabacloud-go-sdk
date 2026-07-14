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
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all instances in a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The list of IP addresses in the IP address whitelist group. You can add up to 1,000 IP addresses, separated by commas (,). The value 127.0.0.1 indicates that no external IP addresses are allowed to access the instance. The following formats are supported:
	//
	// - 10.23.12.24 (IP address)
	//
	// - 10.23.12.24/24 (CIDR pattern, Classless Inter-Domain Routing. /24 specifies the prefix length, which ranges from 1 to 32.)
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service type. Currently, only drama is supported.
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
