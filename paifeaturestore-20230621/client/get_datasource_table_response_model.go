// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasourceTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasourceTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasourceTableResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasourceTableResponseBody) *GetDatasourceTableResponse
	GetBody() *GetDatasourceTableResponseBody
}

type GetDatasourceTableResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasourceTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasourceTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasourceTableResponse) GoString() string {
	return s.String()
}

func (s *GetDatasourceTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasourceTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasourceTableResponse) GetBody() *GetDatasourceTableResponseBody {
	return s.Body
}

func (s *GetDatasourceTableResponse) SetHeaders(v map[string]*string) *GetDatasourceTableResponse {
	s.Headers = v
	return s
}

func (s *GetDatasourceTableResponse) SetStatusCode(v int32) *GetDatasourceTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasourceTableResponse) SetBody(v *GetDatasourceTableResponseBody) *GetDatasourceTableResponse {
	s.Body = v
	return s
}

func (s *GetDatasourceTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
