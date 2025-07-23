// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigImageRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigImageRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigImageRemarkResponse
	GetStatusCode() *int32
	SetBody(v *ConfigImageRemarkResponseBody) *ConfigImageRemarkResponse
	GetBody() *ConfigImageRemarkResponseBody
}

type ConfigImageRemarkResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigImageRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigImageRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigImageRemarkResponse) GoString() string {
	return s.String()
}

func (s *ConfigImageRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigImageRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigImageRemarkResponse) GetBody() *ConfigImageRemarkResponseBody {
	return s.Body
}

func (s *ConfigImageRemarkResponse) SetHeaders(v map[string]*string) *ConfigImageRemarkResponse {
	s.Headers = v
	return s
}

func (s *ConfigImageRemarkResponse) SetStatusCode(v int32) *ConfigImageRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigImageRemarkResponse) SetBody(v *ConfigImageRemarkResponseBody) *ConfigImageRemarkResponse {
	s.Body = v
	return s
}

func (s *ConfigImageRemarkResponse) Validate() error {
	return dara.Validate(s)
}
