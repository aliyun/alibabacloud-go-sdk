// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteListStrategyStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWhiteListStrategyStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWhiteListStrategyStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWhiteListStrategyStatusResponseBody) *UpdateWhiteListStrategyStatusResponse
	GetBody() *UpdateWhiteListStrategyStatusResponseBody
}

type UpdateWhiteListStrategyStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWhiteListStrategyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWhiteListStrategyStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteListStrategyStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhiteListStrategyStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWhiteListStrategyStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWhiteListStrategyStatusResponse) GetBody() *UpdateWhiteListStrategyStatusResponseBody {
	return s.Body
}

func (s *UpdateWhiteListStrategyStatusResponse) SetHeaders(v map[string]*string) *UpdateWhiteListStrategyStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateWhiteListStrategyStatusResponse) SetStatusCode(v int32) *UpdateWhiteListStrategyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWhiteListStrategyStatusResponse) SetBody(v *UpdateWhiteListStrategyStatusResponseBody) *UpdateWhiteListStrategyStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateWhiteListStrategyStatusResponse) Validate() error {
	return dara.Validate(s)
}
