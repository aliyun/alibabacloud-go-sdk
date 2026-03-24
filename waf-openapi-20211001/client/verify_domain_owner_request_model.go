// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyDomainOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *VerifyDomainOwnerRequest
	GetDomainName() *string
	SetInstanceId(v string) *VerifyDomainOwnerRequest
	GetInstanceId() *string
	SetProtocol(v string) *VerifyDomainOwnerRequest
	GetProtocol() *string
	SetVerifyType(v string) *VerifyDomainOwnerRequest
	GetVerifyType() *string
}

type VerifyDomainOwnerRequest struct {
	// The domain name whose ownership you want to verify.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to obtain the WAF instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The protocol type used for file verification. This parameter is required only when you set VerifyType to fileCheck. Valid values:
	//
	// - **HTTP**: the HTTP protocol.
	//
	// - **HTTPS**: the HTTPS protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The verification method. Valid values:
	//
	// - **dnsCheck**: DNS verification.
	//
	// - **fileCheck**: File verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// dnsCheck
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyDomainOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *VerifyDomainOwnerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *VerifyDomainOwnerRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *VerifyDomainOwnerRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *VerifyDomainOwnerRequest) SetDomainName(v string) *VerifyDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyDomainOwnerRequest) SetInstanceId(v string) *VerifyDomainOwnerRequest {
	s.InstanceId = &v
	return s
}

func (s *VerifyDomainOwnerRequest) SetProtocol(v string) *VerifyDomainOwnerRequest {
	s.Protocol = &v
	return s
}

func (s *VerifyDomainOwnerRequest) SetVerifyType(v string) *VerifyDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

func (s *VerifyDomainOwnerRequest) Validate() error {
	return dara.Validate(s)
}
