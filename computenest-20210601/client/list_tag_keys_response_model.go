// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse
	GetBody() *ListTagKeysResponseBody
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagKeysResponse) GetBody() *ListTagKeysResponseBody {
	return s.Body
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

func (s *ListTagKeysResponse) Validate() error {
	return dara.Validate(s)
}
