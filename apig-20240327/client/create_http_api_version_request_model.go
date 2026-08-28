// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersionConfig(v *HttpApiVersionConfig) *CreateHttpApiVersionRequest
	GetVersionConfig() *HttpApiVersionConfig
}

type CreateHttpApiVersionRequest struct {
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
}

func (s CreateHttpApiVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpApiVersionRequest) GetVersionConfig() *HttpApiVersionConfig {
	return s.VersionConfig
}

func (s *CreateHttpApiVersionRequest) SetVersionConfig(v *HttpApiVersionConfig) *CreateHttpApiVersionRequest {
	s.VersionConfig = v
	return s
}

func (s *CreateHttpApiVersionRequest) Validate() error {
	if s.VersionConfig != nil {
		if err := s.VersionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
