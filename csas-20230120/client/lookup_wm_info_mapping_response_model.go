// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupWmInfoMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LookupWmInfoMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LookupWmInfoMappingResponse
	GetStatusCode() *int32
	SetBody(v *LookupWmInfoMappingResponseBody) *LookupWmInfoMappingResponse
	GetBody() *LookupWmInfoMappingResponseBody
}

type LookupWmInfoMappingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LookupWmInfoMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LookupWmInfoMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s LookupWmInfoMappingResponse) GoString() string {
	return s.String()
}

func (s *LookupWmInfoMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LookupWmInfoMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LookupWmInfoMappingResponse) GetBody() *LookupWmInfoMappingResponseBody {
	return s.Body
}

func (s *LookupWmInfoMappingResponse) SetHeaders(v map[string]*string) *LookupWmInfoMappingResponse {
	s.Headers = v
	return s
}

func (s *LookupWmInfoMappingResponse) SetStatusCode(v int32) *LookupWmInfoMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *LookupWmInfoMappingResponse) SetBody(v *LookupWmInfoMappingResponseBody) *LookupWmInfoMappingResponse {
	s.Body = v
	return s
}

func (s *LookupWmInfoMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
