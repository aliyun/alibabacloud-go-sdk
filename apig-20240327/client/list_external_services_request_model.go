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
	ImportableOnly *bool   `json:"importableOnly,omitempty" xml:"importableOnly,omitempty"`
	Limit          *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NameLike       *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	PaiWorkspaceId *string `json:"paiWorkspaceId,omitempty" xml:"paiWorkspaceId,omitempty"`
	SourceType     *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
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
