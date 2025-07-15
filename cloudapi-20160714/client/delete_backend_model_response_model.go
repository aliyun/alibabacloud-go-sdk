// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackendModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackendModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackendModelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackendModelResponseBody) *DeleteBackendModelResponse
	GetBody() *DeleteBackendModelResponseBody
}

type DeleteBackendModelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackendModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackendModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackendModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackendModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackendModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackendModelResponse) GetBody() *DeleteBackendModelResponseBody {
	return s.Body
}

func (s *DeleteBackendModelResponse) SetHeaders(v map[string]*string) *DeleteBackendModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackendModelResponse) SetStatusCode(v int32) *DeleteBackendModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackendModelResponse) SetBody(v *DeleteBackendModelResponseBody) *DeleteBackendModelResponse {
	s.Body = v
	return s
}

func (s *DeleteBackendModelResponse) Validate() error {
	return dara.Validate(s)
}
