// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultNormalizationRuleVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultNormalizationRuleVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultNormalizationRuleVersionResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultNormalizationRuleVersionResponseBody) *SetDefaultNormalizationRuleVersionResponse
	GetBody() *SetDefaultNormalizationRuleVersionResponseBody
}

type SetDefaultNormalizationRuleVersionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultNormalizationRuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultNormalizationRuleVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultNormalizationRuleVersionResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultNormalizationRuleVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultNormalizationRuleVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultNormalizationRuleVersionResponse) GetBody() *SetDefaultNormalizationRuleVersionResponseBody {
	return s.Body
}

func (s *SetDefaultNormalizationRuleVersionResponse) SetHeaders(v map[string]*string) *SetDefaultNormalizationRuleVersionResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponse) SetStatusCode(v int32) *SetDefaultNormalizationRuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponse) SetBody(v *SetDefaultNormalizationRuleVersionResponseBody) *SetDefaultNormalizationRuleVersionResponse {
	s.Body = v
	return s
}

func (s *SetDefaultNormalizationRuleVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
