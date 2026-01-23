// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStandardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStandardResponse
	GetStatusCode() *int32
	SetBody(v *CreateStandardResponseBody) *CreateStandardResponse
	GetBody() *CreateStandardResponseBody
}

type CreateStandardResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStandardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStandardResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStandardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStandardResponse) GetBody() *CreateStandardResponseBody {
	return s.Body
}

func (s *CreateStandardResponse) SetHeaders(v map[string]*string) *CreateStandardResponse {
	s.Headers = v
	return s
}

func (s *CreateStandardResponse) SetStatusCode(v int32) *CreateStandardResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStandardResponse) SetBody(v *CreateStandardResponseBody) *CreateStandardResponse {
	s.Body = v
	return s
}

func (s *CreateStandardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
