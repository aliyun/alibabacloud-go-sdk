// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChartRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChartRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChartRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *GetChartRepositoryResponseBody) *GetChartRepositoryResponse
	GetBody() *GetChartRepositoryResponseBody
}

type GetChartRepositoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChartRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *GetChartRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChartRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChartRepositoryResponse) GetBody() *GetChartRepositoryResponseBody {
	return s.Body
}

func (s *GetChartRepositoryResponse) SetHeaders(v map[string]*string) *GetChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *GetChartRepositoryResponse) SetStatusCode(v int32) *GetChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChartRepositoryResponse) SetBody(v *GetChartRepositoryResponseBody) *GetChartRepositoryResponse {
	s.Body = v
	return s
}

func (s *GetChartRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
