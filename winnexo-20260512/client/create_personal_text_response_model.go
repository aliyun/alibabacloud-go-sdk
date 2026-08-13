// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalTextResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalTextResponseBody) *CreatePersonalTextResponse
	GetBody() *CreatePersonalTextResponseBody
}

type CreatePersonalTextResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalTextResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTextResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalTextResponse) GetBody() *CreatePersonalTextResponseBody {
	return s.Body
}

func (s *CreatePersonalTextResponse) SetHeaders(v map[string]*string) *CreatePersonalTextResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalTextResponse) SetStatusCode(v int32) *CreatePersonalTextResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalTextResponse) SetBody(v *CreatePersonalTextResponseBody) *CreatePersonalTextResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
