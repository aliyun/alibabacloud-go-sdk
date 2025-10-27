// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStrategyResponseBody) *ModifyStrategyResponse
	GetBody() *ModifyStrategyResponseBody
}

type ModifyStrategyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStrategyResponse) GetBody() *ModifyStrategyResponseBody {
	return s.Body
}

func (s *ModifyStrategyResponse) SetHeaders(v map[string]*string) *ModifyStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyStrategyResponse) SetStatusCode(v int32) *ModifyStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStrategyResponse) SetBody(v *ModifyStrategyResponseBody) *ModifyStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
