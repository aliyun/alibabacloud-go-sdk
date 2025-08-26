// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataLakeTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataLakeTableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataLakeTableResponseBody) *DeleteDataLakeTableResponse
	GetBody() *DeleteDataLakeTableResponseBody
}

type DeleteDataLakeTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLakeTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataLakeTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataLakeTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataLakeTableResponse) GetBody() *DeleteDataLakeTableResponseBody {
	return s.Body
}

func (s *DeleteDataLakeTableResponse) SetHeaders(v map[string]*string) *DeleteDataLakeTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLakeTableResponse) SetStatusCode(v int32) *DeleteDataLakeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLakeTableResponse) SetBody(v *DeleteDataLakeTableResponseBody) *DeleteDataLakeTableResponse {
	s.Body = v
	return s
}

func (s *DeleteDataLakeTableResponse) Validate() error {
	return dara.Validate(s)
}
