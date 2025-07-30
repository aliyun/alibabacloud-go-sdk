// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarningStrategyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWarningStrategyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWarningStrategyConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListWarningStrategyConfigResponseBody) *ListWarningStrategyConfigResponse
	GetBody() *ListWarningStrategyConfigResponseBody
}

type ListWarningStrategyConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWarningStrategyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *ListWarningStrategyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWarningStrategyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWarningStrategyConfigResponse) GetBody() *ListWarningStrategyConfigResponseBody {
	return s.Body
}

func (s *ListWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *ListWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *ListWarningStrategyConfigResponse) SetStatusCode(v int32) *ListWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarningStrategyConfigResponse) SetBody(v *ListWarningStrategyConfigResponseBody) *ListWarningStrategyConfigResponse {
	s.Body = v
	return s
}

func (s *ListWarningStrategyConfigResponse) Validate() error {
	return dara.Validate(s)
}
