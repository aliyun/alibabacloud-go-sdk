// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarningStrategyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWarningStrategyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWarningStrategyConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWarningStrategyConfigResponseBody) *UpdateWarningStrategyConfigResponse
	GetBody() *UpdateWarningStrategyConfigResponseBody
}

type UpdateWarningStrategyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWarningStrategyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateWarningStrategyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWarningStrategyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWarningStrategyConfigResponse) GetBody() *UpdateWarningStrategyConfigResponseBody {
	return s.Body
}

func (s *UpdateWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *UpdateWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateWarningStrategyConfigResponse) SetStatusCode(v int32) *UpdateWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWarningStrategyConfigResponse) SetBody(v *UpdateWarningStrategyConfigResponseBody) *UpdateWarningStrategyConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateWarningStrategyConfigResponse) Validate() error {
	return dara.Validate(s)
}
