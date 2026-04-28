// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUncompressConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UncompressConfigResponse
	GetEnabled() *bool
	SetVersion(v string) *UncompressConfigResponse
	GetVersion() *string
}

type UncompressConfigResponse struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UncompressConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UncompressConfigResponse) GoString() string {
	return s.String()
}

func (s *UncompressConfigResponse) GetEnabled() *bool {
	return s.Enabled
}

func (s *UncompressConfigResponse) GetVersion() *string {
	return s.Version
}

func (s *UncompressConfigResponse) SetEnabled(v bool) *UncompressConfigResponse {
	s.Enabled = &v
	return s
}

func (s *UncompressConfigResponse) SetVersion(v string) *UncompressConfigResponse {
	s.Version = &v
	return s
}

func (s *UncompressConfigResponse) Validate() error {
	return dara.Validate(s)
}
