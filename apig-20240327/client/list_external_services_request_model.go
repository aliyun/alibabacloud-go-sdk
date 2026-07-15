// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportableOnly(v bool) *ListExternalServicesRequest
	GetImportableOnly() *bool
	SetLimit(v int32) *ListExternalServicesRequest
	GetLimit() *int32
	SetNameLike(v string) *ListExternalServicesRequest
	GetNameLike() *string
	SetPaiWorkspaceId(v string) *ListExternalServicesRequest
	GetPaiWorkspaceId() *string
	SetSourceType(v string) *ListExternalServicesRequest
	GetSourceType() *string
}

type ListExternalServicesRequest struct {
	// Specifies whether to return only services that have not been imported.
	//
	// example:
	//
	// true
	ImportableOnly *bool `json:"importableOnly,omitempty" xml:"importableOnly,omitempty"`
	// The maximum number of entries to return. Valid range: (0, 100]. Default value: 10.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The API name for fuzzy match.
	//
	// example:
	//
	// imah
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The workspace ID of the PAI-EAS service.
	//
	// example:
	//
	// 667435
	PaiWorkspaceId *string `json:"paiWorkspaceId,omitempty" xml:"paiWorkspaceId,omitempty"`
	// The service source type used to filter results. Valid values:
	//
	// - MSE_NACOS: services from MSE Nacos.
	//
	// - K8S: services from a Kubernetes cluster in Container Service.
	//
	// - FC3: services from Function Compute.
	//
	// - VIP: services from a fixed address.
	//
	// - DNS: services from a domain name.
	//
	// example:
	//
	// FC3
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListExternalServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExternalServicesRequest) GoString() string {
	return s.String()
}

func (s *ListExternalServicesRequest) GetImportableOnly() *bool {
	return s.ImportableOnly
}

func (s *ListExternalServicesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListExternalServicesRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListExternalServicesRequest) GetPaiWorkspaceId() *string {
	return s.PaiWorkspaceId
}

func (s *ListExternalServicesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListExternalServicesRequest) SetImportableOnly(v bool) *ListExternalServicesRequest {
	s.ImportableOnly = &v
	return s
}

func (s *ListExternalServicesRequest) SetLimit(v int32) *ListExternalServicesRequest {
	s.Limit = &v
	return s
}

func (s *ListExternalServicesRequest) SetNameLike(v string) *ListExternalServicesRequest {
	s.NameLike = &v
	return s
}

func (s *ListExternalServicesRequest) SetPaiWorkspaceId(v string) *ListExternalServicesRequest {
	s.PaiWorkspaceId = &v
	return s
}

func (s *ListExternalServicesRequest) SetSourceType(v string) *ListExternalServicesRequest {
	s.SourceType = &v
	return s
}

func (s *ListExternalServicesRequest) Validate() error {
	return dara.Validate(s)
}
