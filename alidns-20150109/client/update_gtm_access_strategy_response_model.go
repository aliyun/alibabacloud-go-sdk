// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmAccessStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGtmAccessStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGtmAccessStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGtmAccessStrategyResponseBody) *UpdateGtmAccessStrategyResponse
	GetBody() *UpdateGtmAccessStrategyResponseBody
}

type UpdateGtmAccessStrategyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGtmAccessStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmAccessStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGtmAccessStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGtmAccessStrategyResponse) GetBody() *UpdateGtmAccessStrategyResponseBody {
	return s.Body
}

func (s *UpdateGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *UpdateGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmAccessStrategyResponse) SetStatusCode(v int32) *UpdateGtmAccessStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGtmAccessStrategyResponse) SetBody(v *UpdateGtmAccessStrategyResponseBody) *UpdateGtmAccessStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateGtmAccessStrategyResponse) Validate() error {
	return dara.Validate(s)
}
