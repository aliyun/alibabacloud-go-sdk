// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateZnodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateZnodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateZnodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateZnodeResponseBody) *CreateZnodeResponse
	GetBody() *CreateZnodeResponseBody
}

type CreateZnodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateZnodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateZnodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateZnodeResponse) GoString() string {
	return s.String()
}

func (s *CreateZnodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateZnodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateZnodeResponse) GetBody() *CreateZnodeResponseBody {
	return s.Body
}

func (s *CreateZnodeResponse) SetHeaders(v map[string]*string) *CreateZnodeResponse {
	s.Headers = v
	return s
}

func (s *CreateZnodeResponse) SetStatusCode(v int32) *CreateZnodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateZnodeResponse) SetBody(v *CreateZnodeResponseBody) *CreateZnodeResponse {
	s.Body = v
	return s
}

func (s *CreateZnodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
