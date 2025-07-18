// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIndexResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIndexResponseBody) *DeleteIndexResponse
	GetBody() *DeleteIndexResponseBody
}

type DeleteIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIndexResponse) GetBody() *DeleteIndexResponseBody {
	return s.Body
}

func (s *DeleteIndexResponse) SetHeaders(v map[string]*string) *DeleteIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexResponse) SetStatusCode(v int32) *DeleteIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexResponse) SetBody(v *DeleteIndexResponseBody) *DeleteIndexResponse {
	s.Body = v
	return s
}

func (s *DeleteIndexResponse) Validate() error {
	return dara.Validate(s)
}
