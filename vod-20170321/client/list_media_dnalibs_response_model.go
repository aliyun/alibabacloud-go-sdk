// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaDNALibsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaDNALibsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaDNALibsResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaDNALibsResponseBody) *ListMediaDNALibsResponse
	GetBody() *ListMediaDNALibsResponseBody
}

type ListMediaDNALibsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaDNALibsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaDNALibsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNALibsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaDNALibsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaDNALibsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaDNALibsResponse) GetBody() *ListMediaDNALibsResponseBody {
	return s.Body
}

func (s *ListMediaDNALibsResponse) SetHeaders(v map[string]*string) *ListMediaDNALibsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaDNALibsResponse) SetStatusCode(v int32) *ListMediaDNALibsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaDNALibsResponse) SetBody(v *ListMediaDNALibsResponseBody) *ListMediaDNALibsResponse {
	s.Body = v
	return s
}

func (s *ListMediaDNALibsResponse) Validate() error {
	return dara.Validate(s)
}
