// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExtensionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExtensionResponseBody) *DeleteExtensionResponse
	GetBody() *DeleteExtensionResponseBody
}

type DeleteExtensionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExtensionResponse) GoString() string {
	return s.String()
}

func (s *DeleteExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExtensionResponse) GetBody() *DeleteExtensionResponseBody {
	return s.Body
}

func (s *DeleteExtensionResponse) SetHeaders(v map[string]*string) *DeleteExtensionResponse {
	s.Headers = v
	return s
}

func (s *DeleteExtensionResponse) SetStatusCode(v int32) *DeleteExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExtensionResponse) SetBody(v *DeleteExtensionResponseBody) *DeleteExtensionResponse {
	s.Body = v
	return s
}

func (s *DeleteExtensionResponse) Validate() error {
	return dara.Validate(s)
}
