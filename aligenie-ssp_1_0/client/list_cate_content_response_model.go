// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCateContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCateContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCateContentResponse
	GetStatusCode() *int32
	SetBody(v *ListCateContentResponseBody) *ListCateContentResponse
	GetBody() *ListCateContentResponseBody
}

type ListCateContentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCateContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCateContentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentResponse) GoString() string {
	return s.String()
}

func (s *ListCateContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCateContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCateContentResponse) GetBody() *ListCateContentResponseBody {
	return s.Body
}

func (s *ListCateContentResponse) SetHeaders(v map[string]*string) *ListCateContentResponse {
	s.Headers = v
	return s
}

func (s *ListCateContentResponse) SetStatusCode(v int32) *ListCateContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCateContentResponse) SetBody(v *ListCateContentResponseBody) *ListCateContentResponse {
	s.Body = v
	return s
}

func (s *ListCateContentResponse) Validate() error {
	return dara.Validate(s)
}
