// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChartRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChartRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChartRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChartRepositoryResponseBody) *DeleteChartRepositoryResponse
	GetBody() *DeleteChartRepositoryResponseBody
}

type DeleteChartRepositoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChartRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteChartRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChartRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChartRepositoryResponse) GetBody() *DeleteChartRepositoryResponseBody {
	return s.Body
}

func (s *DeleteChartRepositoryResponse) SetHeaders(v map[string]*string) *DeleteChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteChartRepositoryResponse) SetStatusCode(v int32) *DeleteChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChartRepositoryResponse) SetBody(v *DeleteChartRepositoryResponseBody) *DeleteChartRepositoryResponse {
	s.Body = v
	return s
}

func (s *DeleteChartRepositoryResponse) Validate() error {
	return dara.Validate(s)
}
