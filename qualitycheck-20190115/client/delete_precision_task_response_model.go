// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrecisionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrecisionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrecisionTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrecisionTaskResponseBody) *DeletePrecisionTaskResponse
	GetBody() *DeletePrecisionTaskResponseBody
}

type DeletePrecisionTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrecisionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrecisionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrecisionTaskResponse) GoString() string {
	return s.String()
}

func (s *DeletePrecisionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrecisionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrecisionTaskResponse) GetBody() *DeletePrecisionTaskResponseBody {
	return s.Body
}

func (s *DeletePrecisionTaskResponse) SetHeaders(v map[string]*string) *DeletePrecisionTaskResponse {
	s.Headers = v
	return s
}

func (s *DeletePrecisionTaskResponse) SetStatusCode(v int32) *DeletePrecisionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrecisionTaskResponse) SetBody(v *DeletePrecisionTaskResponseBody) *DeletePrecisionTaskResponse {
	s.Body = v
	return s
}

func (s *DeletePrecisionTaskResponse) Validate() error {
	return dara.Validate(s)
}
