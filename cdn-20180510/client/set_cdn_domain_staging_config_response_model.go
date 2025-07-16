// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnDomainStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCdnDomainStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCdnDomainStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetCdnDomainStagingConfigResponseBody) *SetCdnDomainStagingConfigResponse
	GetBody() *SetCdnDomainStagingConfigResponseBody
}

type SetCdnDomainStagingConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCdnDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCdnDomainStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCdnDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *SetCdnDomainStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCdnDomainStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCdnDomainStagingConfigResponse) GetBody() *SetCdnDomainStagingConfigResponseBody {
	return s.Body
}

func (s *SetCdnDomainStagingConfigResponse) SetHeaders(v map[string]*string) *SetCdnDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *SetCdnDomainStagingConfigResponse) SetStatusCode(v int32) *SetCdnDomainStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCdnDomainStagingConfigResponse) SetBody(v *SetCdnDomainStagingConfigResponseBody) *SetCdnDomainStagingConfigResponse {
	s.Body = v
	return s
}

func (s *SetCdnDomainStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
