// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsByDependencyIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDependencyProjectId(v int64) *ListProjectsByDependencyIdRequest
	GetDependencyProjectId() *int64
	SetJwtToken(v string) *ListProjectsByDependencyIdRequest
	GetJwtToken() *string
	SetMapUuid(v string) *ListProjectsByDependencyIdRequest
	GetMapUuid() *string
	SetWithDataset(v bool) *ListProjectsByDependencyIdRequest
	GetWithDataset() *bool
	SetWithSource(v bool) *ListProjectsByDependencyIdRequest
	GetWithSource() *bool
}

type ListProjectsByDependencyIdRequest struct {
	// This parameter is required.
	DependencyProjectId *int64  `json:"DependencyProjectId,omitempty" xml:"DependencyProjectId,omitempty"`
	JwtToken            *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	MapUuid             *string `json:"MapUuid,omitempty" xml:"MapUuid,omitempty"`
	// This parameter is required.
	WithDataset *bool `json:"WithDataset,omitempty" xml:"WithDataset,omitempty"`
	// This parameter is required.
	WithSource *bool `json:"WithSource,omitempty" xml:"WithSource,omitempty"`
}

func (s ListProjectsByDependencyIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsByDependencyIdRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsByDependencyIdRequest) GetDependencyProjectId() *int64 {
	return s.DependencyProjectId
}

func (s *ListProjectsByDependencyIdRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ListProjectsByDependencyIdRequest) GetMapUuid() *string {
	return s.MapUuid
}

func (s *ListProjectsByDependencyIdRequest) GetWithDataset() *bool {
	return s.WithDataset
}

func (s *ListProjectsByDependencyIdRequest) GetWithSource() *bool {
	return s.WithSource
}

func (s *ListProjectsByDependencyIdRequest) SetDependencyProjectId(v int64) *ListProjectsByDependencyIdRequest {
	s.DependencyProjectId = &v
	return s
}

func (s *ListProjectsByDependencyIdRequest) SetJwtToken(v string) *ListProjectsByDependencyIdRequest {
	s.JwtToken = &v
	return s
}

func (s *ListProjectsByDependencyIdRequest) SetMapUuid(v string) *ListProjectsByDependencyIdRequest {
	s.MapUuid = &v
	return s
}

func (s *ListProjectsByDependencyIdRequest) SetWithDataset(v bool) *ListProjectsByDependencyIdRequest {
	s.WithDataset = &v
	return s
}

func (s *ListProjectsByDependencyIdRequest) SetWithSource(v bool) *ListProjectsByDependencyIdRequest {
	s.WithSource = &v
	return s
}

func (s *ListProjectsByDependencyIdRequest) Validate() error {
	return dara.Validate(s)
}
