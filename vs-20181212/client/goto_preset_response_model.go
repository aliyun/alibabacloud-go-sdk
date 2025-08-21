// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGotoPresetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GotoPresetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GotoPresetResponse
	GetStatusCode() *int32
	SetBody(v *GotoPresetResponseBody) *GotoPresetResponse
	GetBody() *GotoPresetResponseBody
}

type GotoPresetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GotoPresetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GotoPresetResponse) String() string {
	return dara.Prettify(s)
}

func (s GotoPresetResponse) GoString() string {
	return s.String()
}

func (s *GotoPresetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GotoPresetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GotoPresetResponse) GetBody() *GotoPresetResponseBody {
	return s.Body
}

func (s *GotoPresetResponse) SetHeaders(v map[string]*string) *GotoPresetResponse {
	s.Headers = v
	return s
}

func (s *GotoPresetResponse) SetStatusCode(v int32) *GotoPresetResponse {
	s.StatusCode = &v
	return s
}

func (s *GotoPresetResponse) SetBody(v *GotoPresetResponseBody) *GotoPresetResponse {
	s.Body = v
	return s
}

func (s *GotoPresetResponse) Validate() error {
	return dara.Validate(s)
}
