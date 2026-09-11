// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGraphResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGraphResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGraphResponse
	GetStatusCode() *int32
	SetBody(v *CreateGraphResponseBody) *CreateGraphResponse
	GetBody() *CreateGraphResponseBody
}

type CreateGraphResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGraphResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGraphResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGraphResponse) GoString() string {
	return s.String()
}

func (s *CreateGraphResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGraphResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGraphResponse) GetBody() *CreateGraphResponseBody {
	return s.Body
}

func (s *CreateGraphResponse) SetHeaders(v map[string]*string) *CreateGraphResponse {
	s.Headers = v
	return s
}

func (s *CreateGraphResponse) SetStatusCode(v int32) *CreateGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGraphResponse) SetBody(v *CreateGraphResponseBody) *CreateGraphResponse {
	s.Body = v
	return s
}

func (s *CreateGraphResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
