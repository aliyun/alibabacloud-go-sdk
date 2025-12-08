// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePcaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *UpdatePcaCertificateRequest
	GetAliasName() *string
	SetClientToken(v string) *UpdatePcaCertificateRequest
	GetClientToken() *string
	SetIdentifier(v string) *UpdatePcaCertificateRequest
	GetIdentifier() *string
	SetResourceGroupId(v string) *UpdatePcaCertificateRequest
	GetResourceGroupId() *string
	SetTags(v []*UpdatePcaCertificateRequestTags) *UpdatePcaCertificateRequest
	GetTags() []*UpdatePcaCertificateRequestTags
}

type UpdatePcaCertificateRequest struct {
	// example:
	//
	// cert-name
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// XXX
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*UpdatePcaCertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdatePcaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePcaCertificateRequest) GoString() string {
	return s.String()
}

func (s *UpdatePcaCertificateRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *UpdatePcaCertificateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePcaCertificateRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdatePcaCertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdatePcaCertificateRequest) GetTags() []*UpdatePcaCertificateRequestTags {
	return s.Tags
}

func (s *UpdatePcaCertificateRequest) SetAliasName(v string) *UpdatePcaCertificateRequest {
	s.AliasName = &v
	return s
}

func (s *UpdatePcaCertificateRequest) SetClientToken(v string) *UpdatePcaCertificateRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePcaCertificateRequest) SetIdentifier(v string) *UpdatePcaCertificateRequest {
	s.Identifier = &v
	return s
}

func (s *UpdatePcaCertificateRequest) SetResourceGroupId(v string) *UpdatePcaCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePcaCertificateRequest) SetTags(v []*UpdatePcaCertificateRequestTags) *UpdatePcaCertificateRequest {
	s.Tags = v
	return s
}

func (s *UpdatePcaCertificateRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePcaCertificateRequestTags struct {
	// example:
	//
	// runtime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePcaCertificateRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdatePcaCertificateRequestTags) GoString() string {
	return s.String()
}

func (s *UpdatePcaCertificateRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdatePcaCertificateRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdatePcaCertificateRequestTags) SetKey(v string) *UpdatePcaCertificateRequestTags {
	s.Key = &v
	return s
}

func (s *UpdatePcaCertificateRequestTags) SetValue(v string) *UpdatePcaCertificateRequestTags {
	s.Value = &v
	return s
}

func (s *UpdatePcaCertificateRequestTags) Validate() error {
	return dara.Validate(s)
}
