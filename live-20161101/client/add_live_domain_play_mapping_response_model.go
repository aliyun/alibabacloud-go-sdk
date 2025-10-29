// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDomainPlayMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveDomainPlayMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveDomainPlayMappingResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveDomainPlayMappingResponseBody) *AddLiveDomainPlayMappingResponse
	GetBody() *AddLiveDomainPlayMappingResponseBody
}

type AddLiveDomainPlayMappingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveDomainPlayMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveDomainPlayMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainPlayMappingResponse) GoString() string {
	return s.String()
}

func (s *AddLiveDomainPlayMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveDomainPlayMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveDomainPlayMappingResponse) GetBody() *AddLiveDomainPlayMappingResponseBody {
	return s.Body
}

func (s *AddLiveDomainPlayMappingResponse) SetHeaders(v map[string]*string) *AddLiveDomainPlayMappingResponse {
	s.Headers = v
	return s
}

func (s *AddLiveDomainPlayMappingResponse) SetStatusCode(v int32) *AddLiveDomainPlayMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveDomainPlayMappingResponse) SetBody(v *AddLiveDomainPlayMappingResponseBody) *AddLiveDomainPlayMappingResponse {
	s.Body = v
	return s
}

func (s *AddLiveDomainPlayMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
