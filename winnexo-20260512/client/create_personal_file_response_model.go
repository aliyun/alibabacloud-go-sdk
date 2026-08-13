// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalFileResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalFileResponseBody) *CreatePersonalFileResponse
	GetBody() *CreatePersonalFileResponseBody
}

type CreatePersonalFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFileResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalFileResponse) GetBody() *CreatePersonalFileResponseBody {
	return s.Body
}

func (s *CreatePersonalFileResponse) SetHeaders(v map[string]*string) *CreatePersonalFileResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalFileResponse) SetStatusCode(v int32) *CreatePersonalFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalFileResponse) SetBody(v *CreatePersonalFileResponseBody) *CreatePersonalFileResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
