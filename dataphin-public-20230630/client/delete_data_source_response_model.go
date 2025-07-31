// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse
	GetBody() *DeleteDataSourceResponseBody
}

type DeleteDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataSourceResponse) GetBody() *DeleteDataSourceResponseBody {
	return s.Body
}

func (s *DeleteDataSourceResponse) SetHeaders(v map[string]*string) *DeleteDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceResponse) SetStatusCode(v int32) *DeleteDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
