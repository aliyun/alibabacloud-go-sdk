// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContacterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContacterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContacterResponse
	GetStatusCode() *int32
	SetBody(v *CreateContacterResponseBody) *CreateContacterResponse
	GetBody() *CreateContacterResponseBody
}

type CreateContacterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContacterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContacterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContacterResponse) GoString() string {
	return s.String()
}

func (s *CreateContacterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContacterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContacterResponse) GetBody() *CreateContacterResponseBody {
	return s.Body
}

func (s *CreateContacterResponse) SetHeaders(v map[string]*string) *CreateContacterResponse {
	s.Headers = v
	return s
}

func (s *CreateContacterResponse) SetStatusCode(v int32) *CreateContacterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContacterResponse) SetBody(v *CreateContacterResponseBody) *CreateContacterResponse {
	s.Body = v
	return s
}

func (s *CreateContacterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
