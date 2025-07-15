// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDirectorySsoStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SetDirectorySsoStatusRequest
	GetDirectoryId() *string
	SetEnableSso(v bool) *SetDirectorySsoStatusRequest
	GetEnableSso() *bool
	SetRegionId(v string) *SetDirectorySsoStatusRequest
	GetRegionId() *string
}

type SetDirectorySsoStatusRequest struct {
	// The AD directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-h95efs1mbukd9****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Specifies whether to enable SSO. Valid values:
	//
	// 	- true: enables SSO.
	//
	// 	- false: disables SSO.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableSso *bool `json:"EnableSso,omitempty" xml:"EnableSso,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDirectorySsoStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDirectorySsoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetDirectorySsoStatusRequest) GetEnableSso() *bool {
	return s.EnableSso
}

func (s *SetDirectorySsoStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDirectorySsoStatusRequest) SetDirectoryId(v string) *SetDirectorySsoStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetEnableSso(v bool) *SetDirectorySsoStatusRequest {
	s.EnableSso = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetRegionId(v string) *SetDirectorySsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) Validate() error {
	return dara.Validate(s)
}
