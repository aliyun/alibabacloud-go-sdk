// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteColumnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteColumnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteColumnsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteColumnsResponseBody) *DeleteColumnsResponse
	GetBody() *DeleteColumnsResponseBody
}

type DeleteColumnsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteColumnsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteColumnsResponse) GoString() string {
	return s.String()
}

func (s *DeleteColumnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteColumnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteColumnsResponse) GetBody() *DeleteColumnsResponseBody {
	return s.Body
}

func (s *DeleteColumnsResponse) SetHeaders(v map[string]*string) *DeleteColumnsResponse {
	s.Headers = v
	return s
}

func (s *DeleteColumnsResponse) SetStatusCode(v int32) *DeleteColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteColumnsResponse) SetBody(v *DeleteColumnsResponseBody) *DeleteColumnsResponse {
	s.Body = v
	return s
}

func (s *DeleteColumnsResponse) Validate() error {
	return dara.Validate(s)
}
