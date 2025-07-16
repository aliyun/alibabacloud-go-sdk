// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse
	GetBody() *DeleteResourceResponseBody
}

type DeleteResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceResponse) GetBody() *DeleteResourceResponseBody {
	return s.Body
}

func (s *DeleteResourceResponse) SetHeaders(v map[string]*string) *DeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceResponse) SetStatusCode(v int32) *DeleteResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceResponse) SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceResponse) Validate() error {
	return dara.Validate(s)
}
