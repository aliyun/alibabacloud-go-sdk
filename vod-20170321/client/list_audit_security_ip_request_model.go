// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditSecurityIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupName(v string) *ListAuditSecurityIpRequest
	GetSecurityGroupName() *string
}

type ListAuditSecurityIpRequest struct {
	// The name of the review security group in which you want to query IP addresses. If you do not specify this parameter, IP addresses in all review security groups are queried.
	//
	// example:
	//
	// Default
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s ListAuditSecurityIpRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuditSecurityIpRequest) GoString() string {
	return s.String()
}

func (s *ListAuditSecurityIpRequest) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *ListAuditSecurityIpRequest) SetSecurityGroupName(v string) *ListAuditSecurityIpRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *ListAuditSecurityIpRequest) Validate() error {
	return dara.Validate(s)
}
