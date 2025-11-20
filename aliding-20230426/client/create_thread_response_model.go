// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThreadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateThreadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateThreadResponse
	GetStatusCode() *int32
	SetBody(v *CreateThreadResponseBody) *CreateThreadResponse
	GetBody() *CreateThreadResponseBody
}

type CreateThreadResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateThreadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateThreadResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadResponse) GoString() string {
	return s.String()
}

func (s *CreateThreadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateThreadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateThreadResponse) GetBody() *CreateThreadResponseBody {
	return s.Body
}

func (s *CreateThreadResponse) SetHeaders(v map[string]*string) *CreateThreadResponse {
	s.Headers = v
	return s
}

func (s *CreateThreadResponse) SetStatusCode(v int32) *CreateThreadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateThreadResponse) SetBody(v *CreateThreadResponseBody) *CreateThreadResponse {
	s.Body = v
	return s
}

func (s *CreateThreadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
