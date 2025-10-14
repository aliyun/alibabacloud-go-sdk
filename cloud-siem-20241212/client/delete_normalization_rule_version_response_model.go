// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNormalizationRuleVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNormalizationRuleVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNormalizationRuleVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNormalizationRuleVersionResponseBody) *DeleteNormalizationRuleVersionResponse
	GetBody() *DeleteNormalizationRuleVersionResponseBody
}

type DeleteNormalizationRuleVersionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNormalizationRuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNormalizationRuleVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNormalizationRuleVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteNormalizationRuleVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNormalizationRuleVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNormalizationRuleVersionResponse) GetBody() *DeleteNormalizationRuleVersionResponseBody {
	return s.Body
}

func (s *DeleteNormalizationRuleVersionResponse) SetHeaders(v map[string]*string) *DeleteNormalizationRuleVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteNormalizationRuleVersionResponse) SetStatusCode(v int32) *DeleteNormalizationRuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNormalizationRuleVersionResponse) SetBody(v *DeleteNormalizationRuleVersionResponseBody) *DeleteNormalizationRuleVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteNormalizationRuleVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
