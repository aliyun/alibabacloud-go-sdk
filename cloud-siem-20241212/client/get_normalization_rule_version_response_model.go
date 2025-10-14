// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationRuleVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNormalizationRuleVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNormalizationRuleVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetNormalizationRuleVersionResponseBody) *GetNormalizationRuleVersionResponse
	GetBody() *GetNormalizationRuleVersionResponseBody
}

type GetNormalizationRuleVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNormalizationRuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNormalizationRuleVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationRuleVersionResponse) GoString() string {
	return s.String()
}

func (s *GetNormalizationRuleVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNormalizationRuleVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNormalizationRuleVersionResponse) GetBody() *GetNormalizationRuleVersionResponseBody {
	return s.Body
}

func (s *GetNormalizationRuleVersionResponse) SetHeaders(v map[string]*string) *GetNormalizationRuleVersionResponse {
	s.Headers = v
	return s
}

func (s *GetNormalizationRuleVersionResponse) SetStatusCode(v int32) *GetNormalizationRuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNormalizationRuleVersionResponse) SetBody(v *GetNormalizationRuleVersionResponseBody) *GetNormalizationRuleVersionResponse {
	s.Body = v
	return s
}

func (s *GetNormalizationRuleVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
