// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStandardMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStandardMappingResponse
	GetStatusCode() *int32
	SetBody(v *CreateStandardMappingResponseBody) *CreateStandardMappingResponse
	GetBody() *CreateStandardMappingResponseBody
}

type CreateStandardMappingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStandardMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStandardMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardMappingResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStandardMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStandardMappingResponse) GetBody() *CreateStandardMappingResponseBody {
	return s.Body
}

func (s *CreateStandardMappingResponse) SetHeaders(v map[string]*string) *CreateStandardMappingResponse {
	s.Headers = v
	return s
}

func (s *CreateStandardMappingResponse) SetStatusCode(v int32) *CreateStandardMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStandardMappingResponse) SetBody(v *CreateStandardMappingResponseBody) *CreateStandardMappingResponse {
	s.Body = v
	return s
}

func (s *CreateStandardMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
