// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradePolarClawSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradePolarClawSkillsResponse
	GetStatusCode() *int32
	SetBody(v *UpgradePolarClawSkillsResponseBody) *UpgradePolarClawSkillsResponse
	GetBody() *UpgradePolarClawSkillsResponseBody
}

type UpgradePolarClawSkillsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradePolarClawSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradePolarClawSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawSkillsResponse) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradePolarClawSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradePolarClawSkillsResponse) GetBody() *UpgradePolarClawSkillsResponseBody {
	return s.Body
}

func (s *UpgradePolarClawSkillsResponse) SetHeaders(v map[string]*string) *UpgradePolarClawSkillsResponse {
	s.Headers = v
	return s
}

func (s *UpgradePolarClawSkillsResponse) SetStatusCode(v int32) *UpgradePolarClawSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradePolarClawSkillsResponse) SetBody(v *UpgradePolarClawSkillsResponseBody) *UpgradePolarClawSkillsResponse {
	s.Body = v
	return s
}

func (s *UpgradePolarClawSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
