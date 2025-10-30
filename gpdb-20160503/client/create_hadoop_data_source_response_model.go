// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHadoopDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHadoopDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHadoopDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateHadoopDataSourceResponseBody) *CreateHadoopDataSourceResponse
	GetBody() *CreateHadoopDataSourceResponseBody
}

type CreateHadoopDataSourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHadoopDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHadoopDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHadoopDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateHadoopDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHadoopDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHadoopDataSourceResponse) GetBody() *CreateHadoopDataSourceResponseBody {
	return s.Body
}

func (s *CreateHadoopDataSourceResponse) SetHeaders(v map[string]*string) *CreateHadoopDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateHadoopDataSourceResponse) SetStatusCode(v int32) *CreateHadoopDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHadoopDataSourceResponse) SetBody(v *CreateHadoopDataSourceResponseBody) *CreateHadoopDataSourceResponse {
	s.Body = v
	return s
}

func (s *CreateHadoopDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
