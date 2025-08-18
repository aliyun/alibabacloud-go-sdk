// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKvsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKvsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKvsResponse
	GetStatusCode() *int32
	SetBody(v *ListKvsResponseBody) *ListKvsResponse
	GetBody() *ListKvsResponseBody
}

type ListKvsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKvsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKvsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKvsResponse) GoString() string {
	return s.String()
}

func (s *ListKvsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKvsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKvsResponse) GetBody() *ListKvsResponseBody {
	return s.Body
}

func (s *ListKvsResponse) SetHeaders(v map[string]*string) *ListKvsResponse {
	s.Headers = v
	return s
}

func (s *ListKvsResponse) SetStatusCode(v int32) *ListKvsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKvsResponse) SetBody(v *ListKvsResponseBody) *ListKvsResponse {
	s.Body = v
	return s
}

func (s *ListKvsResponse) Validate() error {
	return dara.Validate(s)
}
