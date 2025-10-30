// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRulePriorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRulePriorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRulePriorityResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRulePriorityResponseBody) *ModifyRulePriorityResponse
	GetBody() *ModifyRulePriorityResponseBody
}

type ModifyRulePriorityResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRulePriorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRulePriorityResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRulePriorityResponse) GoString() string {
	return s.String()
}

func (s *ModifyRulePriorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRulePriorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRulePriorityResponse) GetBody() *ModifyRulePriorityResponseBody {
	return s.Body
}

func (s *ModifyRulePriorityResponse) SetHeaders(v map[string]*string) *ModifyRulePriorityResponse {
	s.Headers = v
	return s
}

func (s *ModifyRulePriorityResponse) SetStatusCode(v int32) *ModifyRulePriorityResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRulePriorityResponse) SetBody(v *ModifyRulePriorityResponseBody) *ModifyRulePriorityResponse {
	s.Body = v
	return s
}

func (s *ModifyRulePriorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
