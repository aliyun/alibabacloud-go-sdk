// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHealthCheckConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHealthCheckConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHealthCheckConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHealthCheckConfigResponseBody) *ModifyHealthCheckConfigResponse
	GetBody() *ModifyHealthCheckConfigResponseBody
}

type ModifyHealthCheckConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHealthCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHealthCheckConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHealthCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHealthCheckConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHealthCheckConfigResponse) GetBody() *ModifyHealthCheckConfigResponseBody {
	return s.Body
}

func (s *ModifyHealthCheckConfigResponse) SetHeaders(v map[string]*string) *ModifyHealthCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyHealthCheckConfigResponse) SetStatusCode(v int32) *ModifyHealthCheckConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHealthCheckConfigResponse) SetBody(v *ModifyHealthCheckConfigResponseBody) *ModifyHealthCheckConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyHealthCheckConfigResponse) Validate() error {
	return dara.Validate(s)
}
