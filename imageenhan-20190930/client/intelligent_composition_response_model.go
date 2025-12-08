// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntelligentCompositionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntelligentCompositionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntelligentCompositionResponse
	GetStatusCode() *int32
	SetBody(v *IntelligentCompositionResponseBody) *IntelligentCompositionResponse
	GetBody() *IntelligentCompositionResponseBody
}

type IntelligentCompositionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntelligentCompositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntelligentCompositionResponse) String() string {
	return dara.Prettify(s)
}

func (s IntelligentCompositionResponse) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntelligentCompositionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntelligentCompositionResponse) GetBody() *IntelligentCompositionResponseBody {
	return s.Body
}

func (s *IntelligentCompositionResponse) SetHeaders(v map[string]*string) *IntelligentCompositionResponse {
	s.Headers = v
	return s
}

func (s *IntelligentCompositionResponse) SetStatusCode(v int32) *IntelligentCompositionResponse {
	s.StatusCode = &v
	return s
}

func (s *IntelligentCompositionResponse) SetBody(v *IntelligentCompositionResponseBody) *IntelligentCompositionResponse {
	s.Body = v
	return s
}

func (s *IntelligentCompositionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
