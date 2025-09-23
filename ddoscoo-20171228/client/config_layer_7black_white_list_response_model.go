// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7BlackWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer7BlackWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer7BlackWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer7BlackWhiteListResponseBody) *ConfigLayer7BlackWhiteListResponse
	GetBody() *ConfigLayer7BlackWhiteListResponseBody
}

type ConfigLayer7BlackWhiteListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7BlackWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer7BlackWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7BlackWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7BlackWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer7BlackWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer7BlackWhiteListResponse) GetBody() *ConfigLayer7BlackWhiteListResponseBody {
	return s.Body
}

func (s *ConfigLayer7BlackWhiteListResponse) SetHeaders(v map[string]*string) *ConfigLayer7BlackWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7BlackWhiteListResponse) SetStatusCode(v int32) *ConfigLayer7BlackWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListResponse) SetBody(v *ConfigLayer7BlackWhiteListResponseBody) *ConfigLayer7BlackWhiteListResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer7BlackWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
