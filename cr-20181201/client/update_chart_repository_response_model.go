// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChartRepositoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChartRepositoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChartRepositoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChartRepositoryResponseBody) *UpdateChartRepositoryResponse
	GetBody() *UpdateChartRepositoryResponseBody
}

type UpdateChartRepositoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChartRepositoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateChartRepositoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChartRepositoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChartRepositoryResponse) GetBody() *UpdateChartRepositoryResponseBody {
	return s.Body
}

func (s *UpdateChartRepositoryResponse) SetHeaders(v map[string]*string) *UpdateChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateChartRepositoryResponse) SetStatusCode(v int32) *UpdateChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChartRepositoryResponse) SetBody(v *UpdateChartRepositoryResponseBody) *UpdateChartRepositoryResponse {
	s.Body = v
	return s
}

func (s *UpdateChartRepositoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
