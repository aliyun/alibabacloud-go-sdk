// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasourceResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasourceResponseBody) *GetDatasourceResponse
	GetBody() *GetDatasourceResponseBody
}

type GetDatasourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasourceResponse) GoString() string {
	return s.String()
}

func (s *GetDatasourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasourceResponse) GetBody() *GetDatasourceResponseBody {
	return s.Body
}

func (s *GetDatasourceResponse) SetHeaders(v map[string]*string) *GetDatasourceResponse {
	s.Headers = v
	return s
}

func (s *GetDatasourceResponse) SetStatusCode(v int32) *GetDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasourceResponse) SetBody(v *GetDatasourceResponseBody) *GetDatasourceResponse {
	s.Body = v
	return s
}

func (s *GetDatasourceResponse) Validate() error {
	return dara.Validate(s)
}
