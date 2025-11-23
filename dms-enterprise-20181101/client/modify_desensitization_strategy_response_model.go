// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesensitizationStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesensitizationStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesensitizationStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesensitizationStrategyResponseBody) *ModifyDesensitizationStrategyResponse
	GetBody() *ModifyDesensitizationStrategyResponseBody
}

type ModifyDesensitizationStrategyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesensitizationStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesensitizationStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesensitizationStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesensitizationStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesensitizationStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesensitizationStrategyResponse) GetBody() *ModifyDesensitizationStrategyResponseBody {
	return s.Body
}

func (s *ModifyDesensitizationStrategyResponse) SetHeaders(v map[string]*string) *ModifyDesensitizationStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesensitizationStrategyResponse) SetStatusCode(v int32) *ModifyDesensitizationStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesensitizationStrategyResponse) SetBody(v *ModifyDesensitizationStrategyResponseBody) *ModifyDesensitizationStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyDesensitizationStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
