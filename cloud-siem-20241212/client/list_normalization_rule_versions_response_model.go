// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRuleVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNormalizationRuleVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNormalizationRuleVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListNormalizationRuleVersionsResponseBody) *ListNormalizationRuleVersionsResponse
	GetBody() *ListNormalizationRuleVersionsResponseBody
}

type ListNormalizationRuleVersionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNormalizationRuleVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNormalizationRuleVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRuleVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListNormalizationRuleVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNormalizationRuleVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNormalizationRuleVersionsResponse) GetBody() *ListNormalizationRuleVersionsResponseBody {
	return s.Body
}

func (s *ListNormalizationRuleVersionsResponse) SetHeaders(v map[string]*string) *ListNormalizationRuleVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListNormalizationRuleVersionsResponse) SetStatusCode(v int32) *ListNormalizationRuleVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponse) SetBody(v *ListNormalizationRuleVersionsResponseBody) *ListNormalizationRuleVersionsResponse {
	s.Body = v
	return s
}

func (s *ListNormalizationRuleVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
