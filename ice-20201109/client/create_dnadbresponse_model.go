// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDNADBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDNADBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDNADBResponse
	GetStatusCode() *int32
	SetBody(v *CreateDNADBResponseBody) *CreateDNADBResponse
	GetBody() *CreateDNADBResponseBody
}

type CreateDNADBResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDNADBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDNADBResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDNADBResponse) GoString() string {
	return s.String()
}

func (s *CreateDNADBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDNADBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDNADBResponse) GetBody() *CreateDNADBResponseBody {
	return s.Body
}

func (s *CreateDNADBResponse) SetHeaders(v map[string]*string) *CreateDNADBResponse {
	s.Headers = v
	return s
}

func (s *CreateDNADBResponse) SetStatusCode(v int32) *CreateDNADBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDNADBResponse) SetBody(v *CreateDNADBResponseBody) *CreateDNADBResponse {
	s.Body = v
	return s
}

func (s *CreateDNADBResponse) Validate() error {
	return dara.Validate(s)
}
