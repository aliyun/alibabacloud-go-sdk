// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDatasourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDatasourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDatasourceResponseBody) *UpdateDatasourceResponse
	GetBody() *UpdateDatasourceResponseBody
}

type UpdateDatasourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDatasourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDatasourceResponse) GetBody() *UpdateDatasourceResponseBody {
	return s.Body
}

func (s *UpdateDatasourceResponse) SetHeaders(v map[string]*string) *UpdateDatasourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasourceResponse) SetStatusCode(v int32) *UpdateDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasourceResponse) SetBody(v *UpdateDatasourceResponseBody) *UpdateDatasourceResponse {
	s.Body = v
	return s
}

func (s *UpdateDatasourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
