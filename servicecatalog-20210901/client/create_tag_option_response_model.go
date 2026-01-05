// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTagOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTagOptionResponse
	GetStatusCode() *int32
	SetBody(v *CreateTagOptionResponseBody) *CreateTagOptionResponse
	GetBody() *CreateTagOptionResponseBody
}

type CreateTagOptionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTagOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTagOptionResponse) GoString() string {
	return s.String()
}

func (s *CreateTagOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTagOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTagOptionResponse) GetBody() *CreateTagOptionResponseBody {
	return s.Body
}

func (s *CreateTagOptionResponse) SetHeaders(v map[string]*string) *CreateTagOptionResponse {
	s.Headers = v
	return s
}

func (s *CreateTagOptionResponse) SetStatusCode(v int32) *CreateTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTagOptionResponse) SetBody(v *CreateTagOptionResponseBody) *CreateTagOptionResponse {
	s.Body = v
	return s
}

func (s *CreateTagOptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
