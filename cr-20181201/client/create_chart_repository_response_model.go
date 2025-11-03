// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChartRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChartRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChartRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateChartRepositoryResponseBody) *CreateChartRepositoryResponse
	GetBody() *CreateChartRepositoryResponseBody
}

type CreateChartRepositoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChartRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *CreateChartRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChartRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChartRepositoryResponse) GetBody() *CreateChartRepositoryResponseBody {
	return s.Body
}

func (s *CreateChartRepositoryResponse) SetHeaders(v map[string]*string) *CreateChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *CreateChartRepositoryResponse) SetStatusCode(v int32) *CreateChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChartRepositoryResponse) SetBody(v *CreateChartRepositoryResponseBody) *CreateChartRepositoryResponse {
	s.Body = v
	return s
}

func (s *CreateChartRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
