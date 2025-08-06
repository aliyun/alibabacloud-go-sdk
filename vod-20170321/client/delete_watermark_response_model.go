// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWatermarkResponseBody) *DeleteWatermarkResponse
	GetBody() *DeleteWatermarkResponseBody
}

type DeleteWatermarkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWatermarkResponse) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWatermarkResponse) GetBody() *DeleteWatermarkResponseBody {
	return s.Body
}

func (s *DeleteWatermarkResponse) SetHeaders(v map[string]*string) *DeleteWatermarkResponse {
	s.Headers = v
	return s
}

func (s *DeleteWatermarkResponse) SetStatusCode(v int32) *DeleteWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWatermarkResponse) SetBody(v *DeleteWatermarkResponseBody) *DeleteWatermarkResponse {
	s.Body = v
	return s
}

func (s *DeleteWatermarkResponse) Validate() error {
	return dara.Validate(s)
}
