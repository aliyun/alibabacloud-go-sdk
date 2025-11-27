// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPresetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPresetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPresetResponse
	GetStatusCode() *int32
	SetBody(v *SetPresetResponseBody) *SetPresetResponse
	GetBody() *SetPresetResponseBody
}

type SetPresetResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPresetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPresetResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPresetResponse) GoString() string {
	return s.String()
}

func (s *SetPresetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPresetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPresetResponse) GetBody() *SetPresetResponseBody {
	return s.Body
}

func (s *SetPresetResponse) SetHeaders(v map[string]*string) *SetPresetResponse {
	s.Headers = v
	return s
}

func (s *SetPresetResponse) SetStatusCode(v int32) *SetPresetResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPresetResponse) SetBody(v *SetPresetResponseBody) *SetPresetResponse {
	s.Body = v
	return s
}

func (s *SetPresetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
