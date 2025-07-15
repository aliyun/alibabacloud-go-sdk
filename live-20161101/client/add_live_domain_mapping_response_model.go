// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDomainMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveDomainMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveDomainMappingResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveDomainMappingResponseBody) *AddLiveDomainMappingResponse
	GetBody() *AddLiveDomainMappingResponseBody
}

type AddLiveDomainMappingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveDomainMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveDomainMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainMappingResponse) GoString() string {
	return s.String()
}

func (s *AddLiveDomainMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveDomainMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveDomainMappingResponse) GetBody() *AddLiveDomainMappingResponseBody {
	return s.Body
}

func (s *AddLiveDomainMappingResponse) SetHeaders(v map[string]*string) *AddLiveDomainMappingResponse {
	s.Headers = v
	return s
}

func (s *AddLiveDomainMappingResponse) SetStatusCode(v int32) *AddLiveDomainMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveDomainMappingResponse) SetBody(v *AddLiveDomainMappingResponseBody) *AddLiveDomainMappingResponse {
	s.Body = v
	return s
}

func (s *AddLiveDomainMappingResponse) Validate() error {
	return dara.Validate(s)
}
