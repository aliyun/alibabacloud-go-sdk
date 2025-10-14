// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNormalizationRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNormalizationRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListNormalizationRulesResponseBody) *ListNormalizationRulesResponse
	GetBody() *ListNormalizationRulesResponseBody
}

type ListNormalizationRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNormalizationRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNormalizationRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRulesResponse) GoString() string {
	return s.String()
}

func (s *ListNormalizationRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNormalizationRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNormalizationRulesResponse) GetBody() *ListNormalizationRulesResponseBody {
	return s.Body
}

func (s *ListNormalizationRulesResponse) SetHeaders(v map[string]*string) *ListNormalizationRulesResponse {
	s.Headers = v
	return s
}

func (s *ListNormalizationRulesResponse) SetStatusCode(v int32) *ListNormalizationRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNormalizationRulesResponse) SetBody(v *ListNormalizationRulesResponseBody) *ListNormalizationRulesResponse {
	s.Body = v
	return s
}

func (s *ListNormalizationRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
