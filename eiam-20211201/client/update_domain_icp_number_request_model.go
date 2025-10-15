// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainIcpNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *UpdateDomainIcpNumberRequest
	GetDomainId() *string
	SetIcpNumber(v string) *UpdateDomainIcpNumberRequest
	GetIcpNumber() *string
	SetInstanceId(v string) *UpdateDomainIcpNumberRequest
	GetInstanceId() *string
}

type UpdateDomainIcpNumberRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// 域名关联的备案号，长度最大限制64。
	//
	// This parameter is required.
	//
	// example:
	//
	// 浙xx-xxxxxx
	IcpNumber *string `json:"IcpNumber,omitempty" xml:"IcpNumber,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDomainIcpNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainIcpNumberRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainIcpNumberRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *UpdateDomainIcpNumberRequest) GetIcpNumber() *string {
	return s.IcpNumber
}

func (s *UpdateDomainIcpNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDomainIcpNumberRequest) SetDomainId(v string) *UpdateDomainIcpNumberRequest {
	s.DomainId = &v
	return s
}

func (s *UpdateDomainIcpNumberRequest) SetIcpNumber(v string) *UpdateDomainIcpNumberRequest {
	s.IcpNumber = &v
	return s
}

func (s *UpdateDomainIcpNumberRequest) SetInstanceId(v string) *UpdateDomainIcpNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDomainIcpNumberRequest) Validate() error {
	return dara.Validate(s)
}
