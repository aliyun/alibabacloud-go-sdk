// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasMinEffectPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotasMinEffectPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotasMinEffectPlanResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotasMinEffectPlanResponseBody) *ListQuotasMinEffectPlanResponse
	GetBody() *ListQuotasMinEffectPlanResponseBody
}

type ListQuotasMinEffectPlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotasMinEffectPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotasMinEffectPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasMinEffectPlanResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasMinEffectPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotasMinEffectPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotasMinEffectPlanResponse) GetBody() *ListQuotasMinEffectPlanResponseBody {
	return s.Body
}

func (s *ListQuotasMinEffectPlanResponse) SetHeaders(v map[string]*string) *ListQuotasMinEffectPlanResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasMinEffectPlanResponse) SetStatusCode(v int32) *ListQuotasMinEffectPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotasMinEffectPlanResponse) SetBody(v *ListQuotasMinEffectPlanResponseBody) *ListQuotasMinEffectPlanResponse {
	s.Body = v
	return s
}

func (s *ListQuotasMinEffectPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
