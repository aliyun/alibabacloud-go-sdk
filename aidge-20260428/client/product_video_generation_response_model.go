// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductVideoGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProductVideoGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProductVideoGenerationResponse
	GetStatusCode() *int32
	SetBody(v *ProductVideoGenerationResponseBody) *ProductVideoGenerationResponse
	GetBody() *ProductVideoGenerationResponseBody
}

type ProductVideoGenerationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProductVideoGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProductVideoGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s ProductVideoGenerationResponse) GoString() string {
	return s.String()
}

func (s *ProductVideoGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProductVideoGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProductVideoGenerationResponse) GetBody() *ProductVideoGenerationResponseBody {
	return s.Body
}

func (s *ProductVideoGenerationResponse) SetHeaders(v map[string]*string) *ProductVideoGenerationResponse {
	s.Headers = v
	return s
}

func (s *ProductVideoGenerationResponse) SetStatusCode(v int32) *ProductVideoGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ProductVideoGenerationResponse) SetBody(v *ProductVideoGenerationResponseBody) *ProductVideoGenerationResponse {
	s.Body = v
	return s
}

func (s *ProductVideoGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
