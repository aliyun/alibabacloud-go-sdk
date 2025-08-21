// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCacheSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebCacheSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebCacheSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebCacheSwitchResponseBody) *ModifyWebCacheSwitchResponse
	GetBody() *ModifyWebCacheSwitchResponseBody
}

type ModifyWebCacheSwitchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebCacheSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebCacheSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCacheSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebCacheSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebCacheSwitchResponse) GetBody() *ModifyWebCacheSwitchResponseBody {
	return s.Body
}

func (s *ModifyWebCacheSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebCacheSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebCacheSwitchResponse) SetStatusCode(v int32) *ModifyWebCacheSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebCacheSwitchResponse) SetBody(v *ModifyWebCacheSwitchResponseBody) *ModifyWebCacheSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyWebCacheSwitchResponse) Validate() error {
	return dara.Validate(s)
}
