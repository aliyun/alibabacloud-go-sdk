// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyThreatIntelligenceSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyThreatIntelligenceSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyThreatIntelligenceSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyThreatIntelligenceSwitchResponseBody) *ModifyThreatIntelligenceSwitchResponse
	GetBody() *ModifyThreatIntelligenceSwitchResponseBody
}

type ModifyThreatIntelligenceSwitchResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyThreatIntelligenceSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyThreatIntelligenceSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyThreatIntelligenceSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyThreatIntelligenceSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyThreatIntelligenceSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyThreatIntelligenceSwitchResponse) GetBody() *ModifyThreatIntelligenceSwitchResponseBody {
	return s.Body
}

func (s *ModifyThreatIntelligenceSwitchResponse) SetHeaders(v map[string]*string) *ModifyThreatIntelligenceSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyThreatIntelligenceSwitchResponse) SetStatusCode(v int32) *ModifyThreatIntelligenceSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyThreatIntelligenceSwitchResponse) SetBody(v *ModifyThreatIntelligenceSwitchResponseBody) *ModifyThreatIntelligenceSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyThreatIntelligenceSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
