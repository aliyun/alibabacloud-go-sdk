// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCatalogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCatalogsResponse
	GetStatusCode() *int32
	SetBody(v *ListCatalogsResponseBody) *ListCatalogsResponse
	GetBody() *ListCatalogsResponseBody
}

type ListCatalogsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCatalogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCatalogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsResponse) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCatalogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCatalogsResponse) GetBody() *ListCatalogsResponseBody {
	return s.Body
}

func (s *ListCatalogsResponse) SetHeaders(v map[string]*string) *ListCatalogsResponse {
	s.Headers = v
	return s
}

func (s *ListCatalogsResponse) SetStatusCode(v int32) *ListCatalogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCatalogsResponse) SetBody(v *ListCatalogsResponseBody) *ListCatalogsResponse {
	s.Body = v
	return s
}

func (s *ListCatalogsResponse) Validate() error {
	return dara.Validate(s)
}
