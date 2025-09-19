// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertStrategyResponseBody) *DeleteAlertStrategyResponse
	GetBody() *DeleteAlertStrategyResponseBody
}

type DeleteAlertStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertStrategyResponse) GetBody() *DeleteAlertStrategyResponseBody {
	return s.Body
}

func (s *DeleteAlertStrategyResponse) SetHeaders(v map[string]*string) *DeleteAlertStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertStrategyResponse) SetStatusCode(v int32) *DeleteAlertStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertStrategyResponse) SetBody(v *DeleteAlertStrategyResponseBody) *DeleteAlertStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertStrategyResponse) Validate() error {
	return dara.Validate(s)
}
