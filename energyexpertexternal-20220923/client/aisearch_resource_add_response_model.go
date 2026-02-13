// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceAddResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AISearchResourceAddResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AISearchResourceAddResponse
	GetStatusCode() *int32
	SetBody(v *AISearchResourceAddResponseBody) *AISearchResourceAddResponse
	GetBody() *AISearchResourceAddResponseBody
}

type AISearchResourceAddResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AISearchResourceAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AISearchResourceAddResponse) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceAddResponse) GoString() string {
	return s.String()
}

func (s *AISearchResourceAddResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AISearchResourceAddResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AISearchResourceAddResponse) GetBody() *AISearchResourceAddResponseBody {
	return s.Body
}

func (s *AISearchResourceAddResponse) SetHeaders(v map[string]*string) *AISearchResourceAddResponse {
	s.Headers = v
	return s
}

func (s *AISearchResourceAddResponse) SetStatusCode(v int32) *AISearchResourceAddResponse {
	s.StatusCode = &v
	return s
}

func (s *AISearchResourceAddResponse) SetBody(v *AISearchResourceAddResponseBody) *AISearchResourceAddResponse {
	s.Body = v
	return s
}

func (s *AISearchResourceAddResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
