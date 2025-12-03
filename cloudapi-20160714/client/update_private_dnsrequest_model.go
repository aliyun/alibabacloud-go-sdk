// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateDNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntranetDomain(v string) *UpdatePrivateDNSRequest
	GetIntranetDomain() *string
	SetRecords(v []*UpdatePrivateDNSRequestRecords) *UpdatePrivateDNSRequest
	GetRecords() []*UpdatePrivateDNSRequestRecords
	SetSecurityToken(v string) *UpdatePrivateDNSRequest
	GetSecurityToken() *string
	SetType(v string) *UpdatePrivateDNSRequest
	GetType() *string
}

type UpdatePrivateDNSRequest struct {
	// The internal domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// api.demo.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// The resolution records. This parameter is valid only when Type is set to A.
	//
	// This parameter is required.
	Records       []*UpdatePrivateDNSRequestRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	SecurityToken *string                           `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The internal domain name resolution type. Valid values:
	//
	// 	- VPC: resolution for virtual private cloud (VPC) access authorizations. A resolution of this type can be bound only to traditional dedicated instances.
	//
	// 	- A: resolution that supports A records. A resolution of this type can be bound only to VPC integration dedicated instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrivateDNSRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateDNSRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateDNSRequest) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *UpdatePrivateDNSRequest) GetRecords() []*UpdatePrivateDNSRequestRecords {
	return s.Records
}

func (s *UpdatePrivateDNSRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdatePrivateDNSRequest) GetType() *string {
	return s.Type
}

func (s *UpdatePrivateDNSRequest) SetIntranetDomain(v string) *UpdatePrivateDNSRequest {
	s.IntranetDomain = &v
	return s
}

func (s *UpdatePrivateDNSRequest) SetRecords(v []*UpdatePrivateDNSRequestRecords) *UpdatePrivateDNSRequest {
	s.Records = v
	return s
}

func (s *UpdatePrivateDNSRequest) SetSecurityToken(v string) *UpdatePrivateDNSRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdatePrivateDNSRequest) SetType(v string) *UpdatePrivateDNSRequest {
	s.Type = &v
	return s
}

func (s *UpdatePrivateDNSRequest) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePrivateDNSRequestRecords struct {
	// The resolution record.
	//
	// example:
	//
	// 192.168.0.2
	Record *string `json:"Record,omitempty" xml:"Record,omitempty"`
	// The weight of the record.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdatePrivateDNSRequestRecords) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateDNSRequestRecords) GoString() string {
	return s.String()
}

func (s *UpdatePrivateDNSRequestRecords) GetRecord() *string {
	return s.Record
}

func (s *UpdatePrivateDNSRequestRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdatePrivateDNSRequestRecords) SetRecord(v string) *UpdatePrivateDNSRequestRecords {
	s.Record = &v
	return s
}

func (s *UpdatePrivateDNSRequestRecords) SetWeight(v int32) *UpdatePrivateDNSRequestRecords {
	s.Weight = &v
	return s
}

func (s *UpdatePrivateDNSRequestRecords) Validate() error {
	return dara.Validate(s)
}
