// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSQLReviewOriginSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSQLReviewOriginSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSQLReviewOriginSQLResponse
	GetStatusCode() *int32
	SetBody(v *ListSQLReviewOriginSQLResponseBody) *ListSQLReviewOriginSQLResponse
	GetBody() *ListSQLReviewOriginSQLResponseBody
}

type ListSQLReviewOriginSQLResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSQLReviewOriginSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSQLReviewOriginSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSQLReviewOriginSQLResponse) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSQLReviewOriginSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSQLReviewOriginSQLResponse) GetBody() *ListSQLReviewOriginSQLResponseBody {
	return s.Body
}

func (s *ListSQLReviewOriginSQLResponse) SetHeaders(v map[string]*string) *ListSQLReviewOriginSQLResponse {
	s.Headers = v
	return s
}

func (s *ListSQLReviewOriginSQLResponse) SetStatusCode(v int32) *ListSQLReviewOriginSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponse) SetBody(v *ListSQLReviewOriginSQLResponseBody) *ListSQLReviewOriginSQLResponse {
	s.Body = v
	return s
}

func (s *ListSQLReviewOriginSQLResponse) Validate() error {
	return dara.Validate(s)
}
