// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgenticDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgenticDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgenticDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgenticDataSourceResponseBody) *UpdateAgenticDataSourceResponse
	GetBody() *UpdateAgenticDataSourceResponseBody
}

type UpdateAgenticDataSourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgenticDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgenticDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgenticDataSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgenticDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgenticDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgenticDataSourceResponse) GetBody() *UpdateAgenticDataSourceResponseBody {
	return s.Body
}

func (s *UpdateAgenticDataSourceResponse) SetHeaders(v map[string]*string) *UpdateAgenticDataSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgenticDataSourceResponse) SetStatusCode(v int32) *UpdateAgenticDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgenticDataSourceResponse) SetBody(v *UpdateAgenticDataSourceResponseBody) *UpdateAgenticDataSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateAgenticDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
