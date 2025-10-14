// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRuleCapacitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNormalizationRuleCapacitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNormalizationRuleCapacitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListNormalizationRuleCapacitiesResponseBody) *ListNormalizationRuleCapacitiesResponse
	GetBody() *ListNormalizationRuleCapacitiesResponseBody
}

type ListNormalizationRuleCapacitiesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNormalizationRuleCapacitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNormalizationRuleCapacitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRuleCapacitiesResponse) GoString() string {
	return s.String()
}

func (s *ListNormalizationRuleCapacitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNormalizationRuleCapacitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNormalizationRuleCapacitiesResponse) GetBody() *ListNormalizationRuleCapacitiesResponseBody {
	return s.Body
}

func (s *ListNormalizationRuleCapacitiesResponse) SetHeaders(v map[string]*string) *ListNormalizationRuleCapacitiesResponse {
	s.Headers = v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponse) SetStatusCode(v int32) *ListNormalizationRuleCapacitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponse) SetBody(v *ListNormalizationRuleCapacitiesResponseBody) *ListNormalizationRuleCapacitiesResponse {
	s.Body = v
	return s
}

func (s *ListNormalizationRuleCapacitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
