// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDNADBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDNADBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDNADBResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDNADBResponseBody) *DeleteDNADBResponse
	GetBody() *DeleteDNADBResponseBody
}

type DeleteDNADBResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDNADBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDNADBResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDNADBResponse) GoString() string {
	return s.String()
}

func (s *DeleteDNADBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDNADBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDNADBResponse) GetBody() *DeleteDNADBResponseBody {
	return s.Body
}

func (s *DeleteDNADBResponse) SetHeaders(v map[string]*string) *DeleteDNADBResponse {
	s.Headers = v
	return s
}

func (s *DeleteDNADBResponse) SetStatusCode(v int32) *DeleteDNADBResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDNADBResponse) SetBody(v *DeleteDNADBResponseBody) *DeleteDNADBResponse {
	s.Body = v
	return s
}

func (s *DeleteDNADBResponse) Validate() error {
	return dara.Validate(s)
}
