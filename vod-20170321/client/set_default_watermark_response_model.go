// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultWatermarkResponseBody) *SetDefaultWatermarkResponse
	GetBody() *SetDefaultWatermarkResponseBody
}

type SetDefaultWatermarkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultWatermarkResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultWatermarkResponse) GetBody() *SetDefaultWatermarkResponseBody {
	return s.Body
}

func (s *SetDefaultWatermarkResponse) SetHeaders(v map[string]*string) *SetDefaultWatermarkResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultWatermarkResponse) SetStatusCode(v int32) *SetDefaultWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultWatermarkResponse) SetBody(v *SetDefaultWatermarkResponseBody) *SetDefaultWatermarkResponse {
	s.Body = v
	return s
}

func (s *SetDefaultWatermarkResponse) Validate() error {
	return dara.Validate(s)
}
