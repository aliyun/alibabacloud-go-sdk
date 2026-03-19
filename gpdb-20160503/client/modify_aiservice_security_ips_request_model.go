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
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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
