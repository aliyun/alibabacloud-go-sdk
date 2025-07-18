// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHadoopDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHadoopDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHadoopDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHadoopDataSourceResponseBody) *DeleteHadoopDataSourceResponse
	GetBody() *DeleteHadoopDataSourceResponseBody
}

type DeleteHadoopDataSourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHadoopDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHadoopDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHadoopDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteHadoopDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHadoopDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHadoopDataSourceResponse) GetBody() *DeleteHadoopDataSourceResponseBody {
	return s.Body
}

func (s *DeleteHadoopDataSourceResponse) SetHeaders(v map[string]*string) *DeleteHadoopDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteHadoopDataSourceResponse) SetStatusCode(v int32) *DeleteHadoopDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHadoopDataSourceResponse) SetBody(v *DeleteHadoopDataSourceResponseBody) *DeleteHadoopDataSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteHadoopDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
