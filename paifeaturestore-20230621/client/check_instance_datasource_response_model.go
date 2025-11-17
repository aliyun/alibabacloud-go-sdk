// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceDatasourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInstanceDatasourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInstanceDatasourceResponse
	GetStatusCode() *int32
	SetBody(v *CheckInstanceDatasourceResponseBody) *CheckInstanceDatasourceResponse
	GetBody() *CheckInstanceDatasourceResponseBody
}

type CheckInstanceDatasourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceDatasourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceDatasourceResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceDatasourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInstanceDatasourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInstanceDatasourceResponse) GetBody() *CheckInstanceDatasourceResponseBody {
	return s.Body
}

func (s *CheckInstanceDatasourceResponse) SetHeaders(v map[string]*string) *CheckInstanceDatasourceResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceDatasourceResponse) SetStatusCode(v int32) *CheckInstanceDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceDatasourceResponse) SetBody(v *CheckInstanceDatasourceResponseBody) *CheckInstanceDatasourceResponse {
	s.Body = v
	return s
}

func (s *CheckInstanceDatasourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
