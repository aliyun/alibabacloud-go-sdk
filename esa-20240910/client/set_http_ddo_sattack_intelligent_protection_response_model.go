// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackIntelligentProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetHttpDDoSAttackIntelligentProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetHttpDDoSAttackIntelligentProtectionResponse
	GetStatusCode() *int32
	SetBody(v *SetHttpDDoSAttackIntelligentProtectionResponseBody) *SetHttpDDoSAttackIntelligentProtectionResponse
	GetBody() *SetHttpDDoSAttackIntelligentProtectionResponseBody
}

type SetHttpDDoSAttackIntelligentProtectionResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetHttpDDoSAttackIntelligentProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetHttpDDoSAttackIntelligentProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackIntelligentProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) GetBody() *SetHttpDDoSAttackIntelligentProtectionResponseBody {
	return s.Body
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) SetHeaders(v map[string]*string) *SetHttpDDoSAttackIntelligentProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) SetStatusCode(v int32) *SetHttpDDoSAttackIntelligentProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) SetBody(v *SetHttpDDoSAttackIntelligentProtectionResponseBody) *SetHttpDDoSAttackIntelligentProtectionResponse {
	s.Body = v
	return s
}

func (s *SetHttpDDoSAttackIntelligentProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
