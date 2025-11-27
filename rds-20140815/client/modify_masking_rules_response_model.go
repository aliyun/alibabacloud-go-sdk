// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaskingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMaskingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMaskingRulesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMaskingRulesResponseBody) *ModifyMaskingRulesResponse
	GetBody() *ModifyMaskingRulesResponseBody
}

type ModifyMaskingRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMaskingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMaskingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaskingRulesResponse) GoString() string {
	return s.String()
}

func (s *ModifyMaskingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMaskingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMaskingRulesResponse) GetBody() *ModifyMaskingRulesResponseBody {
	return s.Body
}

func (s *ModifyMaskingRulesResponse) SetHeaders(v map[string]*string) *ModifyMaskingRulesResponse {
	s.Headers = v
	return s
}

func (s *ModifyMaskingRulesResponse) SetStatusCode(v int32) *ModifyMaskingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMaskingRulesResponse) SetBody(v *ModifyMaskingRulesResponseBody) *ModifyMaskingRulesResponse {
	s.Body = v
	return s
}

func (s *ModifyMaskingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
