// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultResourceDiscoveryAssociationId(v string) *CreateIpamResponseBody
	GetDefaultResourceDiscoveryAssociationId() *string
	SetDefaultResourceDiscoveryId(v string) *CreateIpamResponseBody
	GetDefaultResourceDiscoveryId() *string
	SetIpamId(v string) *CreateIpamResponseBody
	GetIpamId() *string
	SetPrivateDefaultScopeId(v string) *CreateIpamResponseBody
	GetPrivateDefaultScopeId() *string
	SetPublicDefaultScopeId(v string) *CreateIpamResponseBody
	GetPublicDefaultScopeId() *string
	SetRequestId(v string) *CreateIpamResponseBody
	GetRequestId() *string
	SetResourceDiscoveryAssociationCount(v int32) *CreateIpamResponseBody
	GetResourceDiscoveryAssociationCount() *int32
}

type CreateIpamResponseBody struct {
	// The ID of the default resource discovery association.
	//
	// example:
	//
	// ipam-res-disco-assoc-jt5fac8twugdbbgip****
	DefaultResourceDiscoveryAssociationId *string `json:"DefaultResourceDiscoveryAssociationId,omitempty" xml:"DefaultResourceDiscoveryAssociationId,omitempty"`
	// The ID of the default resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	DefaultResourceDiscoveryId *string `json:"DefaultResourceDiscoveryId,omitempty" xml:"DefaultResourceDiscoveryId,omitempty"`
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The default private scope created by the system after the IPAM is created.
	//
	// example:
	//
	// ipam-scope-8wiontzmiy6cg0****
	PrivateDefaultScopeId *string `json:"PrivateDefaultScopeId,omitempty" xml:"PrivateDefaultScopeId,omitempty"`
	// The default public scope created by the system after the IPAM is created.
	//
	// example:
	//
	// ipam-scope-r5c5c7bmym1brc****
	PublicDefaultScopeId *string `json:"PublicDefaultScopeId,omitempty" xml:"PublicDefaultScopeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED39DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of discovered resources.
	//
	// example:
	//
	// 1
	ResourceDiscoveryAssociationCount *int32 `json:"ResourceDiscoveryAssociationCount,omitempty" xml:"ResourceDiscoveryAssociationCount,omitempty"`
}

func (s CreateIpamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamResponseBody) GetDefaultResourceDiscoveryAssociationId() *string {
	return s.DefaultResourceDiscoveryAssociationId
}

func (s *CreateIpamResponseBody) GetDefaultResourceDiscoveryId() *string {
	return s.DefaultResourceDiscoveryId
}

func (s *CreateIpamResponseBody) GetIpamId() *string {
	return s.IpamId
}

func (s *CreateIpamResponseBody) GetPrivateDefaultScopeId() *string {
	return s.PrivateDefaultScopeId
}

func (s *CreateIpamResponseBody) GetPublicDefaultScopeId() *string {
	return s.PublicDefaultScopeId
}

func (s *CreateIpamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpamResponseBody) GetResourceDiscoveryAssociationCount() *int32 {
	return s.ResourceDiscoveryAssociationCount
}

func (s *CreateIpamResponseBody) SetDefaultResourceDiscoveryAssociationId(v string) *CreateIpamResponseBody {
	s.DefaultResourceDiscoveryAssociationId = &v
	return s
}

func (s *CreateIpamResponseBody) SetDefaultResourceDiscoveryId(v string) *CreateIpamResponseBody {
	s.DefaultResourceDiscoveryId = &v
	return s
}

func (s *CreateIpamResponseBody) SetIpamId(v string) *CreateIpamResponseBody {
	s.IpamId = &v
	return s
}

func (s *CreateIpamResponseBody) SetPrivateDefaultScopeId(v string) *CreateIpamResponseBody {
	s.PrivateDefaultScopeId = &v
	return s
}

func (s *CreateIpamResponseBody) SetPublicDefaultScopeId(v string) *CreateIpamResponseBody {
	s.PublicDefaultScopeId = &v
	return s
}

func (s *CreateIpamResponseBody) SetRequestId(v string) *CreateIpamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpamResponseBody) SetResourceDiscoveryAssociationCount(v int32) *CreateIpamResponseBody {
	s.ResourceDiscoveryAssociationCount = &v
	return s
}

func (s *CreateIpamResponseBody) Validate() error {
	return dara.Validate(s)
}
