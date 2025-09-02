// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDagsResponse
	GetStatusCode() *int32
	SetBody(v *ListDagsResponseBody) *ListDagsResponse
	GetBody() *ListDagsResponseBody
}

type ListDagsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDagsResponse) GoString() string {
	return s.String()
}

func (s *ListDagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDagsResponse) GetBody() *ListDagsResponseBody {
	return s.Body
}

func (s *ListDagsResponse) SetHeaders(v map[string]*string) *ListDagsResponse {
	s.Headers = v
	return s
}

func (s *ListDagsResponse) SetStatusCode(v int32) *ListDagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDagsResponse) SetBody(v *ListDagsResponseBody) *ListDagsResponse {
	s.Body = v
	return s
}

func (s *ListDagsResponse) Validate() error {
	return dara.Validate(s)
}
