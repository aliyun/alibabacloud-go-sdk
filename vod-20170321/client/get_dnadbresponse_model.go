// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDNADBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDNADBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDNADBResponse
	GetStatusCode() *int32
	SetBody(v *GetDNADBResponseBody) *GetDNADBResponse
	GetBody() *GetDNADBResponseBody
}

type GetDNADBResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDNADBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDNADBResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDNADBResponse) GoString() string {
	return s.String()
}

func (s *GetDNADBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDNADBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDNADBResponse) GetBody() *GetDNADBResponseBody {
	return s.Body
}

func (s *GetDNADBResponse) SetHeaders(v map[string]*string) *GetDNADBResponse {
	s.Headers = v
	return s
}

func (s *GetDNADBResponse) SetStatusCode(v int32) *GetDNADBResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDNADBResponse) SetBody(v *GetDNADBResponseBody) *GetDNADBResponse {
	s.Body = v
	return s
}

func (s *GetDNADBResponse) Validate() error {
	return dara.Validate(s)
}
