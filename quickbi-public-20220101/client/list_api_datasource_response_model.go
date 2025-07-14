// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiDatasourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiDatasourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiDatasourceResponse
	GetStatusCode() *int32
	SetBody(v *ListApiDatasourceResponseBody) *ListApiDatasourceResponse
	GetBody() *ListApiDatasourceResponseBody
}

type ListApiDatasourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiDatasourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiDatasourceResponse) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiDatasourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiDatasourceResponse) GetBody() *ListApiDatasourceResponseBody {
	return s.Body
}

func (s *ListApiDatasourceResponse) SetHeaders(v map[string]*string) *ListApiDatasourceResponse {
	s.Headers = v
	return s
}

func (s *ListApiDatasourceResponse) SetStatusCode(v int32) *ListApiDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiDatasourceResponse) SetBody(v *ListApiDatasourceResponseBody) *ListApiDatasourceResponse {
	s.Body = v
	return s
}

func (s *ListApiDatasourceResponse) Validate() error {
	return dara.Validate(s)
}
