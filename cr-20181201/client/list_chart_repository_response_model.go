// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChartRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChartRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChartRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *ListChartRepositoryResponseBody) *ListChartRepositoryResponse
	GetBody() *ListChartRepositoryResponseBody
}

type ListChartRepositoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChartRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *ListChartRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChartRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChartRepositoryResponse) GetBody() *ListChartRepositoryResponseBody {
	return s.Body
}

func (s *ListChartRepositoryResponse) SetHeaders(v map[string]*string) *ListChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *ListChartRepositoryResponse) SetStatusCode(v int32) *ListChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChartRepositoryResponse) SetBody(v *ListChartRepositoryResponseBody) *ListChartRepositoryResponse {
	s.Body = v
	return s
}

func (s *ListChartRepositoryResponse) Validate() error {
	return dara.Validate(s)
}
