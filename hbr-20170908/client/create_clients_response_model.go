// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClientsResponse
	GetStatusCode() *int32
	SetBody(v *CreateClientsResponseBody) *CreateClientsResponse
	GetBody() *CreateClientsResponseBody
}

type CreateClientsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClientsResponse) GoString() string {
	return s.String()
}

func (s *CreateClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClientsResponse) GetBody() *CreateClientsResponseBody {
	return s.Body
}

func (s *CreateClientsResponse) SetHeaders(v map[string]*string) *CreateClientsResponse {
	s.Headers = v
	return s
}

func (s *CreateClientsResponse) SetStatusCode(v int32) *CreateClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientsResponse) SetBody(v *CreateClientsResponseBody) *CreateClientsResponse {
	s.Body = v
	return s
}

func (s *CreateClientsResponse) Validate() error {
	return dara.Validate(s)
}
