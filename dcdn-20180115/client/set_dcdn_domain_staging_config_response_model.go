// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDcdnDomainStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDcdnDomainStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetDcdnDomainStagingConfigResponseBody) *SetDcdnDomainStagingConfigResponse
	GetBody() *SetDcdnDomainStagingConfigResponseBody
}

type SetDcdnDomainStagingConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDcdnDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDcdnDomainStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDcdnDomainStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDcdnDomainStagingConfigResponse) GetBody() *SetDcdnDomainStagingConfigResponseBody {
	return s.Body
}

func (s *SetDcdnDomainStagingConfigResponse) SetHeaders(v map[string]*string) *SetDcdnDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnDomainStagingConfigResponse) SetStatusCode(v int32) *SetDcdnDomainStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDcdnDomainStagingConfigResponse) SetBody(v *SetDcdnDomainStagingConfigResponseBody) *SetDcdnDomainStagingConfigResponse {
	s.Body = v
	return s
}

func (s *SetDcdnDomainStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
