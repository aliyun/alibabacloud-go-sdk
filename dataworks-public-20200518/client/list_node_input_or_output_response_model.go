// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeInputOrOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeInputOrOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeInputOrOutputResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeInputOrOutputResponseBody) *ListNodeInputOrOutputResponse
	GetBody() *ListNodeInputOrOutputResponseBody
}

type ListNodeInputOrOutputResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeInputOrOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeInputOrOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInputOrOutputResponse) GoString() string {
	return s.String()
}

func (s *ListNodeInputOrOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeInputOrOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeInputOrOutputResponse) GetBody() *ListNodeInputOrOutputResponseBody {
	return s.Body
}

func (s *ListNodeInputOrOutputResponse) SetHeaders(v map[string]*string) *ListNodeInputOrOutputResponse {
	s.Headers = v
	return s
}

func (s *ListNodeInputOrOutputResponse) SetStatusCode(v int32) *ListNodeInputOrOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeInputOrOutputResponse) SetBody(v *ListNodeInputOrOutputResponseBody) *ListNodeInputOrOutputResponse {
	s.Body = v
	return s
}

func (s *ListNodeInputOrOutputResponse) Validate() error {
	return dara.Validate(s)
}
