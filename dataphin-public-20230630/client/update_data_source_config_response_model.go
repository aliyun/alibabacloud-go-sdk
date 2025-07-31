// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataSourceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataSourceConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataSourceConfigResponseBody) *UpdateDataSourceConfigResponse
	GetBody() *UpdateDataSourceConfigResponseBody
}

type UpdateDataSourceConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSourceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSourceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataSourceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataSourceConfigResponse) GetBody() *UpdateDataSourceConfigResponseBody {
	return s.Body
}

func (s *UpdateDataSourceConfigResponse) SetHeaders(v map[string]*string) *UpdateDataSourceConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSourceConfigResponse) SetStatusCode(v int32) *UpdateDataSourceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSourceConfigResponse) SetBody(v *UpdateDataSourceConfigResponseBody) *UpdateDataSourceConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateDataSourceConfigResponse) Validate() error {
	return dara.Validate(s)
}
