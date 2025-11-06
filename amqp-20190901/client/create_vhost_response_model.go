// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVhostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVhostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVhostResponse
	GetStatusCode() *int32
	SetBody(v *CreateVhostResponseBody) *CreateVhostResponse
	GetBody() *CreateVhostResponseBody
}

type CreateVhostResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVhostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVhostResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVhostResponse) GoString() string {
	return s.String()
}

func (s *CreateVhostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVhostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVhostResponse) GetBody() *CreateVhostResponseBody {
	return s.Body
}

func (s *CreateVhostResponse) SetHeaders(v map[string]*string) *CreateVhostResponse {
	s.Headers = v
	return s
}

func (s *CreateVhostResponse) SetStatusCode(v int32) *CreateVhostResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVhostResponse) SetBody(v *CreateVhostResponseBody) *CreateVhostResponse {
	s.Body = v
	return s
}

func (s *CreateVhostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
