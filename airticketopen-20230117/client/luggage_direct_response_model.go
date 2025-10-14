// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLuggageDirectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LuggageDirectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LuggageDirectResponse
	GetStatusCode() *int32
	SetBody(v *LuggageDirectResponseBody) *LuggageDirectResponse
	GetBody() *LuggageDirectResponseBody
}

type LuggageDirectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LuggageDirectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LuggageDirectResponse) String() string {
	return dara.Prettify(s)
}

func (s LuggageDirectResponse) GoString() string {
	return s.String()
}

func (s *LuggageDirectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LuggageDirectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LuggageDirectResponse) GetBody() *LuggageDirectResponseBody {
	return s.Body
}

func (s *LuggageDirectResponse) SetHeaders(v map[string]*string) *LuggageDirectResponse {
	s.Headers = v
	return s
}

func (s *LuggageDirectResponse) SetStatusCode(v int32) *LuggageDirectResponse {
	s.StatusCode = &v
	return s
}

func (s *LuggageDirectResponse) SetBody(v *LuggageDirectResponseBody) *LuggageDirectResponse {
	s.Body = v
	return s
}

func (s *LuggageDirectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
