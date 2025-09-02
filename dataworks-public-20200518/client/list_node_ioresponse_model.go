// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeIOResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeIOResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeIOResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeIOResponseBody) *ListNodeIOResponse
	GetBody() *ListNodeIOResponseBody
}

type ListNodeIOResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeIOResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeIOResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeIOResponse) GoString() string {
	return s.String()
}

func (s *ListNodeIOResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeIOResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeIOResponse) GetBody() *ListNodeIOResponseBody {
	return s.Body
}

func (s *ListNodeIOResponse) SetHeaders(v map[string]*string) *ListNodeIOResponse {
	s.Headers = v
	return s
}

func (s *ListNodeIOResponse) SetStatusCode(v int32) *ListNodeIOResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeIOResponse) SetBody(v *ListNodeIOResponseBody) *ListNodeIOResponse {
	s.Body = v
	return s
}

func (s *ListNodeIOResponse) Validate() error {
	return dara.Validate(s)
}
