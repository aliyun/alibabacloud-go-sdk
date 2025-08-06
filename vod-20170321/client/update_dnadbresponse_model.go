// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDNADBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDNADBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDNADBResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDNADBResponseBody) *UpdateDNADBResponse
	GetBody() *UpdateDNADBResponseBody
}

type UpdateDNADBResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDNADBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDNADBResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDNADBResponse) GoString() string {
	return s.String()
}

func (s *UpdateDNADBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDNADBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDNADBResponse) GetBody() *UpdateDNADBResponseBody {
	return s.Body
}

func (s *UpdateDNADBResponse) SetHeaders(v map[string]*string) *UpdateDNADBResponse {
	s.Headers = v
	return s
}

func (s *UpdateDNADBResponse) SetStatusCode(v int32) *UpdateDNADBResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDNADBResponse) SetBody(v *UpdateDNADBResponseBody) *UpdateDNADBResponse {
	s.Body = v
	return s
}

func (s *UpdateDNADBResponse) Validate() error {
	return dara.Validate(s)
}
