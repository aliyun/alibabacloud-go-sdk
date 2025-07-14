// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentViewReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecentViewReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecentViewReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecentViewReportsResponseBody) *ListRecentViewReportsResponse
	GetBody() *ListRecentViewReportsResponseBody
}

type ListRecentViewReportsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecentViewReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecentViewReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecentViewReportsResponse) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecentViewReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecentViewReportsResponse) GetBody() *ListRecentViewReportsResponseBody {
	return s.Body
}

func (s *ListRecentViewReportsResponse) SetHeaders(v map[string]*string) *ListRecentViewReportsResponse {
	s.Headers = v
	return s
}

func (s *ListRecentViewReportsResponse) SetStatusCode(v int32) *ListRecentViewReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentViewReportsResponse) SetBody(v *ListRecentViewReportsResponseBody) *ListRecentViewReportsResponse {
	s.Body = v
	return s
}

func (s *ListRecentViewReportsResponse) Validate() error {
	return dara.Validate(s)
}
