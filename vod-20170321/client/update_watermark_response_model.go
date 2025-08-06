// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWatermarkResponseBody) *UpdateWatermarkResponse
	GetBody() *UpdateWatermarkResponseBody
}

type UpdateWatermarkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWatermarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWatermarkResponse) GetBody() *UpdateWatermarkResponseBody {
	return s.Body
}

func (s *UpdateWatermarkResponse) SetHeaders(v map[string]*string) *UpdateWatermarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateWatermarkResponse) SetStatusCode(v int32) *UpdateWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWatermarkResponse) SetBody(v *UpdateWatermarkResponseBody) *UpdateWatermarkResponse {
	s.Body = v
	return s
}

func (s *UpdateWatermarkResponse) Validate() error {
	return dara.Validate(s)
}
