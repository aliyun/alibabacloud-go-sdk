// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyForwardStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyForwardStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyForwardStrategyResponseBody) *ModifyForwardStrategyResponse
	GetBody() *ModifyForwardStrategyResponseBody
}

type ModifyForwardStrategyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyForwardStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyForwardStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyForwardStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyForwardStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyForwardStrategyResponse) GetBody() *ModifyForwardStrategyResponseBody {
	return s.Body
}

func (s *ModifyForwardStrategyResponse) SetHeaders(v map[string]*string) *ModifyForwardStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyForwardStrategyResponse) SetStatusCode(v int32) *ModifyForwardStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyForwardStrategyResponse) SetBody(v *ModifyForwardStrategyResponseBody) *ModifyForwardStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyForwardStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
