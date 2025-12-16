// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSortExpressionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSortExpressionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSortExpressionsResponse
	GetStatusCode() *int32
	SetBody(v *ListSortExpressionsResponseBody) *ListSortExpressionsResponse
	GetBody() *ListSortExpressionsResponseBody
}

type ListSortExpressionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSortExpressionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSortExpressionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSortExpressionsResponse) GoString() string {
	return s.String()
}

func (s *ListSortExpressionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSortExpressionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSortExpressionsResponse) GetBody() *ListSortExpressionsResponseBody {
	return s.Body
}

func (s *ListSortExpressionsResponse) SetHeaders(v map[string]*string) *ListSortExpressionsResponse {
	s.Headers = v
	return s
}

func (s *ListSortExpressionsResponse) SetStatusCode(v int32) *ListSortExpressionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSortExpressionsResponse) SetBody(v *ListSortExpressionsResponseBody) *ListSortExpressionsResponse {
	s.Body = v
	return s
}

func (s *ListSortExpressionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
