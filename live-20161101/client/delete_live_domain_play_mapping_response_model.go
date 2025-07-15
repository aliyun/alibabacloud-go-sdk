// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDomainPlayMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveDomainPlayMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveDomainPlayMappingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveDomainPlayMappingResponseBody) *DeleteLiveDomainPlayMappingResponse
	GetBody() *DeleteLiveDomainPlayMappingResponseBody
}

type DeleteLiveDomainPlayMappingResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveDomainPlayMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveDomainPlayMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDomainPlayMappingResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainPlayMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveDomainPlayMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveDomainPlayMappingResponse) GetBody() *DeleteLiveDomainPlayMappingResponseBody {
	return s.Body
}

func (s *DeleteLiveDomainPlayMappingResponse) SetHeaders(v map[string]*string) *DeleteLiveDomainPlayMappingResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveDomainPlayMappingResponse) SetStatusCode(v int32) *DeleteLiveDomainPlayMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveDomainPlayMappingResponse) SetBody(v *DeleteLiveDomainPlayMappingResponseBody) *DeleteLiveDomainPlayMappingResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveDomainPlayMappingResponse) Validate() error {
	return dara.Validate(s)
}
