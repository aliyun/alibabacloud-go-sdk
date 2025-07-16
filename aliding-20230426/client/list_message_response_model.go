// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessageResponse
	GetStatusCode() *int32
	SetBody(v *ListMessageResponseBody) *ListMessageResponse
	GetBody() *ListMessageResponseBody
}

type ListMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessageResponse) GoString() string {
	return s.String()
}

func (s *ListMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessageResponse) GetBody() *ListMessageResponseBody {
	return s.Body
}

func (s *ListMessageResponse) SetHeaders(v map[string]*string) *ListMessageResponse {
	s.Headers = v
	return s
}

func (s *ListMessageResponse) SetStatusCode(v int32) *ListMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageResponse) SetBody(v *ListMessageResponseBody) *ListMessageResponse {
	s.Body = v
	return s
}

func (s *ListMessageResponse) Validate() error {
	return dara.Validate(s)
}
