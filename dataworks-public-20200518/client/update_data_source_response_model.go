// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataSourceResponseBody) *UpdateDataSourceResponse
	GetBody() *UpdateDataSourceResponseBody
}

type UpdateDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataSourceResponse) GetBody() *UpdateDataSourceResponseBody {
	return s.Body
}

func (s *UpdateDataSourceResponse) SetHeaders(v map[string]*string) *UpdateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSourceResponse) SetStatusCode(v int32) *UpdateDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSourceResponse) SetBody(v *UpdateDataSourceResponseBody) *UpdateDataSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
