// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarningStrategyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWarningStrategyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWarningStrategyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWarningStrategyConfigResponseBody) *DeleteWarningStrategyConfigResponse
	GetBody() *DeleteWarningStrategyConfigResponseBody
}

type DeleteWarningStrategyConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWarningStrategyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteWarningStrategyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWarningStrategyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWarningStrategyConfigResponse) GetBody() *DeleteWarningStrategyConfigResponseBody {
	return s.Body
}

func (s *DeleteWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *DeleteWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteWarningStrategyConfigResponse) SetStatusCode(v int32) *DeleteWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWarningStrategyConfigResponse) SetBody(v *DeleteWarningStrategyConfigResponseBody) *DeleteWarningStrategyConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteWarningStrategyConfigResponse) Validate() error {
	return dara.Validate(s)
}
