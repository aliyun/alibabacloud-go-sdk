// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNormalizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNormalizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNormalizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNormalizationRuleResponseBody) *DeleteNormalizationRuleResponse
	GetBody() *DeleteNormalizationRuleResponseBody
}

type DeleteNormalizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNormalizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNormalizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNormalizationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteNormalizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNormalizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNormalizationRuleResponse) GetBody() *DeleteNormalizationRuleResponseBody {
	return s.Body
}

func (s *DeleteNormalizationRuleResponse) SetHeaders(v map[string]*string) *DeleteNormalizationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteNormalizationRuleResponse) SetStatusCode(v int32) *DeleteNormalizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNormalizationRuleResponse) SetBody(v *DeleteNormalizationRuleResponseBody) *DeleteNormalizationRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteNormalizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
