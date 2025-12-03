// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateDNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntranetDomain(v string) *CreatePrivateDNSRequest
	GetIntranetDomain() *string
	SetRecords(v []*CreatePrivateDNSRequestRecords) *CreatePrivateDNSRequest
	GetRecords() []*CreatePrivateDNSRequestRecords
	SetSecurityToken(v string) *CreatePrivateDNSRequest
	GetSecurityToken() *string
	SetType(v string) *CreatePrivateDNSRequest
	GetType() *string
}

type CreatePrivateDNSRequest struct {
	// The internal domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// api.demo.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// The resolution records. This parameter is valid only when Type is set to A.
	Records       []*CreatePrivateDNSRequestRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	SecurityToken *string                           `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The internal domain name resolution type. Valid values:
	//
	// 	- VPC: resolution for VPC access authorizations. A resolution of this type can be bound only to traditional dedicated instances.
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

func (s CreatePrivateDNSRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateDNSRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateDNSRequest) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *CreatePrivateDNSRequest) GetRecords() []*CreatePrivateDNSRequestRecords {
	return s.Records
}

func (s *CreatePrivateDNSRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreatePrivateDNSRequest) GetType() *string {
	return s.Type
}

func (s *CreatePrivateDNSRequest) SetIntranetDomain(v string) *CreatePrivateDNSRequest {
	s.IntranetDomain = &v
	return s
}

func (s *CreatePrivateDNSRequest) SetRecords(v []*CreatePrivateDNSRequestRecords) *CreatePrivateDNSRequest {
	s.Records = v
	return s
}

func (s *CreatePrivateDNSRequest) SetSecurityToken(v string) *CreatePrivateDNSRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreatePrivateDNSRequest) SetType(v string) *CreatePrivateDNSRequest {
	s.Type = &v
	return s
}

func (s *CreatePrivateDNSRequest) Validate() error {
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

type CreatePrivateDNSRequestRecords struct {
	// The resolution record.
	//
	// example:
	//
	// 192.168.0.1
	Record *string `json:"Record,omitempty" xml:"Record,omitempty"`
	// The weight of the record.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreatePrivateDNSRequestRecords) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateDNSRequestRecords) GoString() string {
	return s.String()
}

func (s *CreatePrivateDNSRequestRecords) GetRecord() *string {
	return s.Record
}

func (s *CreatePrivateDNSRequestRecords) GetWeight() *int32 {
	return s.Weight
}

func (s *CreatePrivateDNSRequestRecords) SetRecord(v string) *CreatePrivateDNSRequestRecords {
	s.Record = &v
	return s
}

func (s *CreatePrivateDNSRequestRecords) SetWeight(v int32) *CreatePrivateDNSRequestRecords {
	s.Weight = &v
	return s
}

func (s *CreatePrivateDNSRequestRecords) Validate() error {
	return dara.Validate(s)
}
