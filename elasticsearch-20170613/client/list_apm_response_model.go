// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApmResponse
	GetStatusCode() *int32
	SetBody(v *ListApmResponseBody) *ListApmResponse
	GetBody() *ListApmResponseBody
}

type ListApmResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApmResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApmResponse) GoString() string {
	return s.String()
}

func (s *ListApmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApmResponse) GetBody() *ListApmResponseBody {
	return s.Body
}

func (s *ListApmResponse) SetHeaders(v map[string]*string) *ListApmResponse {
	s.Headers = v
	return s
}

func (s *ListApmResponse) SetStatusCode(v int32) *ListApmResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApmResponse) SetBody(v *ListApmResponseBody) *ListApmResponse {
	s.Body = v
	return s
}

func (s *ListApmResponse) Validate() error {
	return dara.Validate(s)
}
