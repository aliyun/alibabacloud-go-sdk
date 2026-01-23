// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityRuleSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQualityRuleSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQualityRuleSwitchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQualityRuleSwitchResponseBody) *UpdateQualityRuleSwitchResponse
	GetBody() *UpdateQualityRuleSwitchResponseBody
}

type UpdateQualityRuleSwitchResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQualityRuleSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQualityRuleSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityRuleSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateQualityRuleSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQualityRuleSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQualityRuleSwitchResponse) GetBody() *UpdateQualityRuleSwitchResponseBody {
	return s.Body
}

func (s *UpdateQualityRuleSwitchResponse) SetHeaders(v map[string]*string) *UpdateQualityRuleSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateQualityRuleSwitchResponse) SetStatusCode(v int32) *UpdateQualityRuleSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQualityRuleSwitchResponse) SetBody(v *UpdateQualityRuleSwitchResponseBody) *UpdateQualityRuleSwitchResponse {
	s.Body = v
	return s
}

func (s *UpdateQualityRuleSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
