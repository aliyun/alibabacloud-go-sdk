// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDomainMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveDomainMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveDomainMappingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveDomainMappingResponseBody) *DeleteLiveDomainMappingResponse
	GetBody() *DeleteLiveDomainMappingResponseBody
}

type DeleteLiveDomainMappingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveDomainMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveDomainMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDomainMappingResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveDomainMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveDomainMappingResponse) GetBody() *DeleteLiveDomainMappingResponseBody {
	return s.Body
}

func (s *DeleteLiveDomainMappingResponse) SetHeaders(v map[string]*string) *DeleteLiveDomainMappingResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveDomainMappingResponse) SetStatusCode(v int32) *DeleteLiveDomainMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveDomainMappingResponse) SetBody(v *DeleteLiveDomainMappingResponseBody) *DeleteLiveDomainMappingResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveDomainMappingResponse) Validate() error {
	return dara.Validate(s)
}
