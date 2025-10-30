// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmInfoMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWmInfoMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWmInfoMappingResponse
	GetStatusCode() *int32
	SetBody(v *CreateWmInfoMappingResponseBody) *CreateWmInfoMappingResponse
	GetBody() *CreateWmInfoMappingResponseBody
}

type CreateWmInfoMappingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWmInfoMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWmInfoMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWmInfoMappingResponse) GoString() string {
	return s.String()
}

func (s *CreateWmInfoMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWmInfoMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWmInfoMappingResponse) GetBody() *CreateWmInfoMappingResponseBody {
	return s.Body
}

func (s *CreateWmInfoMappingResponse) SetHeaders(v map[string]*string) *CreateWmInfoMappingResponse {
	s.Headers = v
	return s
}

func (s *CreateWmInfoMappingResponse) SetStatusCode(v int32) *CreateWmInfoMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWmInfoMappingResponse) SetBody(v *CreateWmInfoMappingResponseBody) *CreateWmInfoMappingResponse {
	s.Body = v
	return s
}

func (s *CreateWmInfoMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
