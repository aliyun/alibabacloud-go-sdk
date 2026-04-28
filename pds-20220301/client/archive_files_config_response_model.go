// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveFilesConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *ArchiveFilesConfigResponse
	GetEnabled() *bool
	SetVersion(v string) *ArchiveFilesConfigResponse
	GetVersion() *string
}

type ArchiveFilesConfigResponse struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ArchiveFilesConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ArchiveFilesConfigResponse) GoString() string {
	return s.String()
}

func (s *ArchiveFilesConfigResponse) GetEnabled() *bool {
	return s.Enabled
}

func (s *ArchiveFilesConfigResponse) GetVersion() *string {
	return s.Version
}

func (s *ArchiveFilesConfigResponse) SetEnabled(v bool) *ArchiveFilesConfigResponse {
	s.Enabled = &v
	return s
}

func (s *ArchiveFilesConfigResponse) SetVersion(v string) *ArchiveFilesConfigResponse {
	s.Version = &v
	return s
}

func (s *ArchiveFilesConfigResponse) Validate() error {
	return dara.Validate(s)
}
