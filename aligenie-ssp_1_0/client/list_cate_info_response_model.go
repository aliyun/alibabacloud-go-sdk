// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCateInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListCateInfoResponseBody) *ListCateInfoResponse
	GetBody() *ListCateInfoResponseBody
}

type ListCateInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCateInfoResponse) GoString() string {
	return s.String()
}

func (s *ListCateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCateInfoResponse) GetBody() *ListCateInfoResponseBody {
	return s.Body
}

func (s *ListCateInfoResponse) SetHeaders(v map[string]*string) *ListCateInfoResponse {
	s.Headers = v
	return s
}

func (s *ListCateInfoResponse) SetStatusCode(v int32) *ListCateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCateInfoResponse) SetBody(v *ListCateInfoResponseBody) *ListCateInfoResponse {
	s.Body = v
	return s
}

func (s *ListCateInfoResponse) Validate() error {
	return dara.Validate(s)
}
