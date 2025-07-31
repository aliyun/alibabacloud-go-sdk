// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComputeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComputeSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComputeSourceResponseBody) *DeleteComputeSourceResponse
	GetBody() *DeleteComputeSourceResponseBody
}

type DeleteComputeSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComputeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComputeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteComputeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComputeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComputeSourceResponse) GetBody() *DeleteComputeSourceResponseBody {
	return s.Body
}

func (s *DeleteComputeSourceResponse) SetHeaders(v map[string]*string) *DeleteComputeSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteComputeSourceResponse) SetStatusCode(v int32) *DeleteComputeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComputeSourceResponse) SetBody(v *DeleteComputeSourceResponseBody) *DeleteComputeSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteComputeSourceResponse) Validate() error {
	return dara.Validate(s)
}
