// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveDomainStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveDomainStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveDomainStagingConfigResponseBody) *SetLiveDomainStagingConfigResponse
	GetBody() *SetLiveDomainStagingConfigResponseBody
}

type SetLiveDomainStagingConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveDomainStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *SetLiveDomainStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveDomainStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveDomainStagingConfigResponse) GetBody() *SetLiveDomainStagingConfigResponseBody {
	return s.Body
}

func (s *SetLiveDomainStagingConfigResponse) SetHeaders(v map[string]*string) *SetLiveDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *SetLiveDomainStagingConfigResponse) SetStatusCode(v int32) *SetLiveDomainStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveDomainStagingConfigResponse) SetBody(v *SetLiveDomainStagingConfigResponseBody) *SetLiveDomainStagingConfigResponse {
	s.Body = v
	return s
}

func (s *SetLiveDomainStagingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
