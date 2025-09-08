// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataSourceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataSourceLogResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataSourceLogResponseBody) *DeleteDataSourceLogResponse
	GetBody() *DeleteDataSourceLogResponseBody
}

type DeleteDataSourceLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceLogResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataSourceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataSourceLogResponse) GetBody() *DeleteDataSourceLogResponseBody {
	return s.Body
}

func (s *DeleteDataSourceLogResponse) SetHeaders(v map[string]*string) *DeleteDataSourceLogResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceLogResponse) SetStatusCode(v int32) *DeleteDataSourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceLogResponse) SetBody(v *DeleteDataSourceLogResponseBody) *DeleteDataSourceLogResponse {
	s.Body = v
	return s
}

func (s *DeleteDataSourceLogResponse) Validate() error {
	return dara.Validate(s)
}
