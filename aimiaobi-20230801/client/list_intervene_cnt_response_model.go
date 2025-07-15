// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterveneCntResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterveneCntResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterveneCntResponse
	GetStatusCode() *int32
	SetBody(v *ListInterveneCntResponseBody) *ListInterveneCntResponse
	GetBody() *ListInterveneCntResponseBody
}

type ListInterveneCntResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterveneCntResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterveneCntResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneCntResponse) GoString() string {
	return s.String()
}

func (s *ListInterveneCntResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterveneCntResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterveneCntResponse) GetBody() *ListInterveneCntResponseBody {
	return s.Body
}

func (s *ListInterveneCntResponse) SetHeaders(v map[string]*string) *ListInterveneCntResponse {
	s.Headers = v
	return s
}

func (s *ListInterveneCntResponse) SetStatusCode(v int32) *ListInterveneCntResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterveneCntResponse) SetBody(v *ListInterveneCntResponseBody) *ListInterveneCntResponse {
	s.Body = v
	return s
}

func (s *ListInterveneCntResponse) Validate() error {
	return dara.Validate(s)
}
