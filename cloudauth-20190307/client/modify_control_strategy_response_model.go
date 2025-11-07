// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyControlStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyControlStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyControlStrategyResponseBody) *ModifyControlStrategyResponse
	GetBody() *ModifyControlStrategyResponseBody
}

type ModifyControlStrategyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyControlStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyControlStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyControlStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyControlStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyControlStrategyResponse) GetBody() *ModifyControlStrategyResponseBody {
	return s.Body
}

func (s *ModifyControlStrategyResponse) SetHeaders(v map[string]*string) *ModifyControlStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyControlStrategyResponse) SetStatusCode(v int32) *ModifyControlStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyControlStrategyResponse) SetBody(v *ModifyControlStrategyResponseBody) *ModifyControlStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyControlStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
