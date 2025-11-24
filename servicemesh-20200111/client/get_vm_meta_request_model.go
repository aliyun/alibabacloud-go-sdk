// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *GetVmMetaRequest
	GetNamespace() *string
	SetServiceAccount(v string) *GetVmMetaRequest
	GetServiceAccount() *string
	SetServiceMeshId(v string) *GetVmMetaRequest
	GetServiceMeshId() *string
	SetTrustDomain(v string) *GetVmMetaRequest
	GetTrustDomain() *string
}

type GetVmMetaRequest struct {
	// The name of the namespace. This parameter is valid only after you set the Namespace and the ServiceAccount parameters.
	//
	// example:
	//
	// hello
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The service account. This parameter is valid only after you set the Namespace and the ServiceAccount parameters.
	//
	// example:
	//
	// http-sa
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce51a7de4a5144db88a864ed91****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The trusted domain. Default value: cluster.local. This parameter is valid only after you set the Namespace and the ServiceAccount parameters.
	//
	// example:
	//
	// cluster.local
	TrustDomain *string `json:"TrustDomain,omitempty" xml:"TrustDomain,omitempty"`
}

func (s GetVmMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVmMetaRequest) GoString() string {
	return s.String()
}

func (s *GetVmMetaRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetVmMetaRequest) GetServiceAccount() *string {
	return s.ServiceAccount
}

func (s *GetVmMetaRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetVmMetaRequest) GetTrustDomain() *string {
	return s.TrustDomain
}

func (s *GetVmMetaRequest) SetNamespace(v string) *GetVmMetaRequest {
	s.Namespace = &v
	return s
}

func (s *GetVmMetaRequest) SetServiceAccount(v string) *GetVmMetaRequest {
	s.ServiceAccount = &v
	return s
}

func (s *GetVmMetaRequest) SetServiceMeshId(v string) *GetVmMetaRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetVmMetaRequest) SetTrustDomain(v string) *GetVmMetaRequest {
	s.TrustDomain = &v
	return s
}

func (s *GetVmMetaRequest) Validate() error {
	return dara.Validate(s)
}
