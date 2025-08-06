// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *ListWatermarkResponseBody) *ListWatermarkResponse
	GetBody() *ListWatermarkResponseBody
}

type ListWatermarkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWatermarkResponse) GoString() string {
	return s.String()
}

func (s *ListWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWatermarkResponse) GetBody() *ListWatermarkResponseBody {
	return s.Body
}

func (s *ListWatermarkResponse) SetHeaders(v map[string]*string) *ListWatermarkResponse {
	s.Headers = v
	return s
}

func (s *ListWatermarkResponse) SetStatusCode(v int32) *ListWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWatermarkResponse) SetBody(v *ListWatermarkResponseBody) *ListWatermarkResponse {
	s.Body = v
	return s
}

func (s *ListWatermarkResponse) Validate() error {
	return dara.Validate(s)
}
