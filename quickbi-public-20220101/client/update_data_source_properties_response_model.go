// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourcePropertiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataSourcePropertiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataSourcePropertiesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataSourcePropertiesResponseBody) *UpdateDataSourcePropertiesResponse
	GetBody() *UpdateDataSourcePropertiesResponseBody
}

type UpdateDataSourcePropertiesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSourcePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSourcePropertiesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourcePropertiesResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSourcePropertiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataSourcePropertiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataSourcePropertiesResponse) GetBody() *UpdateDataSourcePropertiesResponseBody {
	return s.Body
}

func (s *UpdateDataSourcePropertiesResponse) SetHeaders(v map[string]*string) *UpdateDataSourcePropertiesResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSourcePropertiesResponse) SetStatusCode(v int32) *UpdateDataSourcePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSourcePropertiesResponse) SetBody(v *UpdateDataSourcePropertiesResponseBody) *UpdateDataSourcePropertiesResponse {
	s.Body = v
	return s
}

func (s *UpdateDataSourcePropertiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
