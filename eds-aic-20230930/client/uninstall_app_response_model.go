// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallAppResponse
	GetStatusCode() *int32
	SetBody(v *UninstallAppResponseBody) *UninstallAppResponse
	GetBody() *UninstallAppResponseBody
}

type UninstallAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallAppResponse) GoString() string {
	return s.String()
}

func (s *UninstallAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallAppResponse) GetBody() *UninstallAppResponseBody {
	return s.Body
}

func (s *UninstallAppResponse) SetHeaders(v map[string]*string) *UninstallAppResponse {
	s.Headers = v
	return s
}

func (s *UninstallAppResponse) SetStatusCode(v int32) *UninstallAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallAppResponse) SetBody(v *UninstallAppResponseBody) *UninstallAppResponse {
	s.Body = v
	return s
}

func (s *UninstallAppResponse) Validate() error {
	return dara.Validate(s)
}
