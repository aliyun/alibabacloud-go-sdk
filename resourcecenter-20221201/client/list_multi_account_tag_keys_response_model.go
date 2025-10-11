// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountTagKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultiAccountTagKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultiAccountTagKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListMultiAccountTagKeysResponseBody) *ListMultiAccountTagKeysResponse
	GetBody() *ListMultiAccountTagKeysResponseBody
}

type ListMultiAccountTagKeysResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountTagKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultiAccountTagKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultiAccountTagKeysResponse) GetBody() *ListMultiAccountTagKeysResponseBody {
	return s.Body
}

func (s *ListMultiAccountTagKeysResponse) SetHeaders(v map[string]*string) *ListMultiAccountTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountTagKeysResponse) SetStatusCode(v int32) *ListMultiAccountTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountTagKeysResponse) SetBody(v *ListMultiAccountTagKeysResponseBody) *ListMultiAccountTagKeysResponse {
	s.Body = v
	return s
}

func (s *ListMultiAccountTagKeysResponse) Validate() error {
	return dara.Validate(s)
}
