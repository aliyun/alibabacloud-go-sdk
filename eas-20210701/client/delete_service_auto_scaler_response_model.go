// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceAutoScalerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceAutoScalerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceAutoScalerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceAutoScalerResponseBody) *DeleteServiceAutoScalerResponse
	GetBody() *DeleteServiceAutoScalerResponseBody
}

type DeleteServiceAutoScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceAutoScalerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceAutoScalerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceAutoScalerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceAutoScalerResponse) GetBody() *DeleteServiceAutoScalerResponseBody {
	return s.Body
}

func (s *DeleteServiceAutoScalerResponse) SetHeaders(v map[string]*string) *DeleteServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceAutoScalerResponse) SetStatusCode(v int32) *DeleteServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceAutoScalerResponse) SetBody(v *DeleteServiceAutoScalerResponseBody) *DeleteServiceAutoScalerResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceAutoScalerResponse) Validate() error {
	return dara.Validate(s)
}
