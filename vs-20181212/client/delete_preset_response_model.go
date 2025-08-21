// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePresetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePresetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePresetResponse
	GetStatusCode() *int32
	SetBody(v *DeletePresetResponseBody) *DeletePresetResponse
	GetBody() *DeletePresetResponseBody
}

type DeletePresetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePresetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePresetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePresetResponse) GoString() string {
	return s.String()
}

func (s *DeletePresetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePresetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePresetResponse) GetBody() *DeletePresetResponseBody {
	return s.Body
}

func (s *DeletePresetResponse) SetHeaders(v map[string]*string) *DeletePresetResponse {
	s.Headers = v
	return s
}

func (s *DeletePresetResponse) SetStatusCode(v int32) *DeletePresetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePresetResponse) SetBody(v *DeletePresetResponseBody) *DeletePresetResponse {
	s.Body = v
	return s
}

func (s *DeletePresetResponse) Validate() error {
	return dara.Validate(s)
}
