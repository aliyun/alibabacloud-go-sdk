// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerDefenseRuleSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyContainerDefenseRuleSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyContainerDefenseRuleSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyContainerDefenseRuleSwitchResponseBody) *ModifyContainerDefenseRuleSwitchResponse
	GetBody() *ModifyContainerDefenseRuleSwitchResponseBody
}

type ModifyContainerDefenseRuleSwitchResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyContainerDefenseRuleSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyContainerDefenseRuleSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyContainerDefenseRuleSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyContainerDefenseRuleSwitchResponse) GetBody() *ModifyContainerDefenseRuleSwitchResponseBody {
	return s.Body
}

func (s *ModifyContainerDefenseRuleSwitchResponse) SetHeaders(v map[string]*string) *ModifyContainerDefenseRuleSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchResponse) SetStatusCode(v int32) *ModifyContainerDefenseRuleSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchResponse) SetBody(v *ModifyContainerDefenseRuleSwitchResponseBody) *ModifyContainerDefenseRuleSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
