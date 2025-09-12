// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomResourceResponseBody) *DeleteCustomResourceResponse
	GetBody() *DeleteCustomResourceResponseBody
}

type DeleteCustomResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomResourceResponse) GetBody() *DeleteCustomResourceResponseBody {
	return s.Body
}

func (s *DeleteCustomResourceResponse) SetHeaders(v map[string]*string) *DeleteCustomResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomResourceResponse) SetStatusCode(v int32) *DeleteCustomResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomResourceResponse) SetBody(v *DeleteCustomResourceResponseBody) *DeleteCustomResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomResourceResponse) Validate() error {
	return dara.Validate(s)
}
