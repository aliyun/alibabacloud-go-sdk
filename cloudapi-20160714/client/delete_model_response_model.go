// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelResponseBody) *DeleteModelResponse
	GetBody() *DeleteModelResponseBody
}

type DeleteModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelResponse) GetBody() *DeleteModelResponseBody {
	return s.Body
}

func (s *DeleteModelResponse) SetHeaders(v map[string]*string) *DeleteModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelResponse) SetStatusCode(v int32) *DeleteModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelResponse) SetBody(v *DeleteModelResponseBody) *DeleteModelResponse {
	s.Body = v
	return s
}

func (s *DeleteModelResponse) Validate() error {
	return dara.Validate(s)
}
