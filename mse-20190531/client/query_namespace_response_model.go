// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *QueryNamespaceResponseBody) *QueryNamespaceResponse
	GetBody() *QueryNamespaceResponseBody
}

type QueryNamespaceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryNamespaceResponse) GoString() string {
	return s.String()
}

func (s *QueryNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryNamespaceResponse) GetBody() *QueryNamespaceResponseBody {
	return s.Body
}

func (s *QueryNamespaceResponse) SetHeaders(v map[string]*string) *QueryNamespaceResponse {
	s.Headers = v
	return s
}

func (s *QueryNamespaceResponse) SetStatusCode(v int32) *QueryNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNamespaceResponse) SetBody(v *QueryNamespaceResponseBody) *QueryNamespaceResponse {
	s.Body = v
	return s
}

func (s *QueryNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
