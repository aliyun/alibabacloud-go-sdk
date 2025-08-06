// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVodDomainStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetVodDomainStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetVodDomainStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetVodDomainStagingConfigResponseBody) *SetVodDomainStagingConfigResponse
	GetBody() *SetVodDomainStagingConfigResponseBody
}

type SetVodDomainStagingConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetVodDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetVodDomainStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *SetVodDomainStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetVodDomainStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetVodDomainStagingConfigResponse) GetBody() *SetVodDomainStagingConfigResponseBody {
	return s.Body
}

func (s *SetVodDomainStagingConfigResponse) SetHeaders(v map[string]*string) *SetVodDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *SetVodDomainStagingConfigResponse) SetStatusCode(v int32) *SetVodDomainStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVodDomainStagingConfigResponse) SetBody(v *SetVodDomainStagingConfigResponseBody) *SetVodDomainStagingConfigResponse {
	s.Body = v
	return s
}

func (s *SetVodDomainStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
