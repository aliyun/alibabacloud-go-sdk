// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageImageRegistryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PageImageRegistryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PageImageRegistryResponse
	GetStatusCode() *int32
	SetBody(v *PageImageRegistryResponseBody) *PageImageRegistryResponse
	GetBody() *PageImageRegistryResponseBody
}

type PageImageRegistryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageImageRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageImageRegistryResponse) String() string {
	return dara.Prettify(s)
}

func (s PageImageRegistryResponse) GoString() string {
	return s.String()
}

func (s *PageImageRegistryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PageImageRegistryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageImageRegistryResponse) GetBody() *PageImageRegistryResponseBody {
	return s.Body
}

func (s *PageImageRegistryResponse) SetHeaders(v map[string]*string) *PageImageRegistryResponse {
	s.Headers = v
	return s
}

func (s *PageImageRegistryResponse) SetStatusCode(v int32) *PageImageRegistryResponse {
	s.StatusCode = &v
	return s
}

func (s *PageImageRegistryResponse) SetBody(v *PageImageRegistryResponseBody) *PageImageRegistryResponse {
	s.Body = v
	return s
}

func (s *PageImageRegistryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
