// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLogstashChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLogstashChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLogstashChargeTypeResponseBody) *UpdateLogstashChargeTypeResponse
	GetBody() *UpdateLogstashChargeTypeResponseBody
}

type UpdateLogstashChargeTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLogstashChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLogstashChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogstashChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLogstashChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLogstashChargeTypeResponse) GetBody() *UpdateLogstashChargeTypeResponseBody {
	return s.Body
}

func (s *UpdateLogstashChargeTypeResponse) SetHeaders(v map[string]*string) *UpdateLogstashChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogstashChargeTypeResponse) SetStatusCode(v int32) *UpdateLogstashChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLogstashChargeTypeResponse) SetBody(v *UpdateLogstashChargeTypeResponseBody) *UpdateLogstashChargeTypeResponse {
	s.Body = v
	return s
}

func (s *UpdateLogstashChargeTypeResponse) Validate() error {
	return dara.Validate(s)
}
