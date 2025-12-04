// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateErResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateErResponse
	GetStatusCode() *int32
	SetBody(v *CreateErResponseBody) *CreateErResponse
	GetBody() *CreateErResponseBody
}

type CreateErResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateErResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateErResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateErResponse) GoString() string {
	return s.String()
}

func (s *CreateErResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateErResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateErResponse) GetBody() *CreateErResponseBody {
	return s.Body
}

func (s *CreateErResponse) SetHeaders(v map[string]*string) *CreateErResponse {
	s.Headers = v
	return s
}

func (s *CreateErResponse) SetStatusCode(v int32) *CreateErResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateErResponse) SetBody(v *CreateErResponseBody) *CreateErResponse {
	s.Body = v
	return s
}

func (s *CreateErResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
