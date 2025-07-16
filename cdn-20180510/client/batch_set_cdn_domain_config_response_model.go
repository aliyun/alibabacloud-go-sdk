// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetCdnDomainConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSetCdnDomainConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSetCdnDomainConfigResponse
	GetStatusCode() *int32
	SetBody(v *BatchSetCdnDomainConfigResponseBody) *BatchSetCdnDomainConfigResponse
	GetBody() *BatchSetCdnDomainConfigResponseBody
}

type BatchSetCdnDomainConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetCdnDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetCdnDomainConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSetCdnDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSetCdnDomainConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSetCdnDomainConfigResponse) GetBody() *BatchSetCdnDomainConfigResponseBody {
	return s.Body
}

func (s *BatchSetCdnDomainConfigResponse) SetHeaders(v map[string]*string) *BatchSetCdnDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *BatchSetCdnDomainConfigResponse) SetStatusCode(v int32) *BatchSetCdnDomainConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetCdnDomainConfigResponse) SetBody(v *BatchSetCdnDomainConfigResponseBody) *BatchSetCdnDomainConfigResponse {
	s.Body = v
	return s
}

func (s *BatchSetCdnDomainConfigResponse) Validate() error {
	return dara.Validate(s)
}
