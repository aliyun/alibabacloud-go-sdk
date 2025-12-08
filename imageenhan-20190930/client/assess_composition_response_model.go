// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssessCompositionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssessCompositionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssessCompositionResponse
	GetStatusCode() *int32
	SetBody(v *AssessCompositionResponseBody) *AssessCompositionResponse
	GetBody() *AssessCompositionResponseBody
}

type AssessCompositionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssessCompositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssessCompositionResponse) String() string {
	return dara.Prettify(s)
}

func (s AssessCompositionResponse) GoString() string {
	return s.String()
}

func (s *AssessCompositionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssessCompositionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssessCompositionResponse) GetBody() *AssessCompositionResponseBody {
	return s.Body
}

func (s *AssessCompositionResponse) SetHeaders(v map[string]*string) *AssessCompositionResponse {
	s.Headers = v
	return s
}

func (s *AssessCompositionResponse) SetStatusCode(v int32) *AssessCompositionResponse {
	s.StatusCode = &v
	return s
}

func (s *AssessCompositionResponse) SetBody(v *AssessCompositionResponseBody) *AssessCompositionResponse {
	s.Body = v
	return s
}

func (s *AssessCompositionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
