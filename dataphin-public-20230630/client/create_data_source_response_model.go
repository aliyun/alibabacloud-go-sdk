// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataSourceResponseBody) *CreateDataSourceResponse
	GetBody() *CreateDataSourceResponseBody
}

type CreateDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataSourceResponse) GetBody() *CreateDataSourceResponseBody {
	return s.Body
}

func (s *CreateDataSourceResponse) SetHeaders(v map[string]*string) *CreateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceResponse) SetStatusCode(v int32) *CreateDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSourceResponse) SetBody(v *CreateDataSourceResponseBody) *CreateDataSourceResponse {
	s.Body = v
	return s
}

func (s *CreateDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
