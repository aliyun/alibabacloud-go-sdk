// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWebAccessLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableWebAccessLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableWebAccessLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *DisableWebAccessLogConfigResponseBody) *DisableWebAccessLogConfigResponse
	GetBody() *DisableWebAccessLogConfigResponseBody
}

type DisableWebAccessLogConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableWebAccessLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableWebAccessLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableWebAccessLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DisableWebAccessLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableWebAccessLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableWebAccessLogConfigResponse) GetBody() *DisableWebAccessLogConfigResponseBody {
	return s.Body
}

func (s *DisableWebAccessLogConfigResponse) SetHeaders(v map[string]*string) *DisableWebAccessLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DisableWebAccessLogConfigResponse) SetStatusCode(v int32) *DisableWebAccessLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableWebAccessLogConfigResponse) SetBody(v *DisableWebAccessLogConfigResponseBody) *DisableWebAccessLogConfigResponse {
	s.Body = v
	return s
}

func (s *DisableWebAccessLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
