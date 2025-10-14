// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmAccessStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGtmAccessStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGtmAccessStrategyResponse
	GetStatusCode() *int32
	SetBody(v *AddGtmAccessStrategyResponseBody) *AddGtmAccessStrategyResponse
	GetBody() *AddGtmAccessStrategyResponseBody
}

type AddGtmAccessStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGtmAccessStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *AddGtmAccessStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGtmAccessStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGtmAccessStrategyResponse) GetBody() *AddGtmAccessStrategyResponseBody {
	return s.Body
}

func (s *AddGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *AddGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *AddGtmAccessStrategyResponse) SetStatusCode(v int32) *AddGtmAccessStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGtmAccessStrategyResponse) SetBody(v *AddGtmAccessStrategyResponseBody) *AddGtmAccessStrategyResponse {
	s.Body = v
	return s
}

func (s *AddGtmAccessStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
