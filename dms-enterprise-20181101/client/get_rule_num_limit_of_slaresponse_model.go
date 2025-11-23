// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleNumLimitOfSLAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRuleNumLimitOfSLAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRuleNumLimitOfSLAResponse
	GetStatusCode() *int32
	SetBody(v *GetRuleNumLimitOfSLAResponseBody) *GetRuleNumLimitOfSLAResponse
	GetBody() *GetRuleNumLimitOfSLAResponseBody
}

type GetRuleNumLimitOfSLAResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleNumLimitOfSLAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleNumLimitOfSLAResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRuleNumLimitOfSLAResponse) GoString() string {
	return s.String()
}

func (s *GetRuleNumLimitOfSLAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRuleNumLimitOfSLAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRuleNumLimitOfSLAResponse) GetBody() *GetRuleNumLimitOfSLAResponseBody {
	return s.Body
}

func (s *GetRuleNumLimitOfSLAResponse) SetHeaders(v map[string]*string) *GetRuleNumLimitOfSLAResponse {
	s.Headers = v
	return s
}

func (s *GetRuleNumLimitOfSLAResponse) SetStatusCode(v int32) *GetRuleNumLimitOfSLAResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleNumLimitOfSLAResponse) SetBody(v *GetRuleNumLimitOfSLAResponseBody) *GetRuleNumLimitOfSLAResponse {
	s.Body = v
	return s
}

func (s *GetRuleNumLimitOfSLAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
