// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMmsDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMmsDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMmsDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMmsDataSourceResponseBody) *DeleteMmsDataSourceResponse
	GetBody() *DeleteMmsDataSourceResponseBody
}

type DeleteMmsDataSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMmsDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMmsDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMmsDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteMmsDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMmsDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMmsDataSourceResponse) GetBody() *DeleteMmsDataSourceResponseBody {
	return s.Body
}

func (s *DeleteMmsDataSourceResponse) SetHeaders(v map[string]*string) *DeleteMmsDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteMmsDataSourceResponse) SetStatusCode(v int32) *DeleteMmsDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMmsDataSourceResponse) SetBody(v *DeleteMmsDataSourceResponseBody) *DeleteMmsDataSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteMmsDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
