// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarningStrategyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWarningStrategyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWarningStrategyConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateWarningStrategyConfigResponseBody) *CreateWarningStrategyConfigResponse
	GetBody() *CreateWarningStrategyConfigResponseBody
}

type CreateWarningStrategyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWarningStrategyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateWarningStrategyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWarningStrategyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWarningStrategyConfigResponse) GetBody() *CreateWarningStrategyConfigResponseBody {
	return s.Body
}

func (s *CreateWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *CreateWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateWarningStrategyConfigResponse) SetStatusCode(v int32) *CreateWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWarningStrategyConfigResponse) SetBody(v *CreateWarningStrategyConfigResponseBody) *CreateWarningStrategyConfigResponse {
	s.Body = v
	return s
}

func (s *CreateWarningStrategyConfigResponse) Validate() error {
	return dara.Validate(s)
}
