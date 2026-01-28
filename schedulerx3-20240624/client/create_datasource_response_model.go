// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatasourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatasourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatasourceResponseBody) *CreateDatasourceResponse
	GetBody() *CreateDatasourceResponseBody
}

type CreateDatasourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatasourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatasourceResponse) GetBody() *CreateDatasourceResponseBody {
	return s.Body
}

func (s *CreateDatasourceResponse) SetHeaders(v map[string]*string) *CreateDatasourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasourceResponse) SetStatusCode(v int32) *CreateDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasourceResponse) SetBody(v *CreateDatasourceResponseBody) *CreateDatasourceResponse {
	s.Body = v
	return s
}

func (s *CreateDatasourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
