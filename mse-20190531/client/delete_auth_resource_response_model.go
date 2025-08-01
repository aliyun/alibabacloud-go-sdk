// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAuthResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAuthResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAuthResourceResponseBody) *DeleteAuthResourceResponse
	GetBody() *DeleteAuthResourceResponseBody
}

type DeleteAuthResourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAuthResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAuthResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuthResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAuthResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAuthResourceResponse) GetBody() *DeleteAuthResourceResponseBody {
	return s.Body
}

func (s *DeleteAuthResourceResponse) SetHeaders(v map[string]*string) *DeleteAuthResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuthResourceResponse) SetStatusCode(v int32) *DeleteAuthResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuthResourceResponse) SetBody(v *DeleteAuthResourceResponseBody) *DeleteAuthResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteAuthResourceResponse) Validate() error {
	return dara.Validate(s)
}
