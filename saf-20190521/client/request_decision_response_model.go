// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestDecisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RequestDecisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RequestDecisionResponse
	GetStatusCode() *int32
	SetBody(v *RequestDecisionResponseBody) *RequestDecisionResponse
	GetBody() *RequestDecisionResponseBody
}

type RequestDecisionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestDecisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestDecisionResponse) String() string {
	return dara.Prettify(s)
}

func (s RequestDecisionResponse) GoString() string {
	return s.String()
}

func (s *RequestDecisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RequestDecisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RequestDecisionResponse) GetBody() *RequestDecisionResponseBody {
	return s.Body
}

func (s *RequestDecisionResponse) SetHeaders(v map[string]*string) *RequestDecisionResponse {
	s.Headers = v
	return s
}

func (s *RequestDecisionResponse) SetStatusCode(v int32) *RequestDecisionResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestDecisionResponse) SetBody(v *RequestDecisionResponseBody) *RequestDecisionResponse {
	s.Body = v
	return s
}

func (s *RequestDecisionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
