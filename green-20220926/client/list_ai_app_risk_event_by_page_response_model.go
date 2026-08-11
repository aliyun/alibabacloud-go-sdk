// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppRiskEventByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAiAppRiskEventByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAiAppRiskEventByPageResponse
	GetStatusCode() *int32
	SetBody(v *ListAiAppRiskEventByPageResponseBody) *ListAiAppRiskEventByPageResponse
	GetBody() *ListAiAppRiskEventByPageResponseBody
}

type ListAiAppRiskEventByPageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiAppRiskEventByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAiAppRiskEventByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppRiskEventByPageResponse) GoString() string {
	return s.String()
}

func (s *ListAiAppRiskEventByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAiAppRiskEventByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAiAppRiskEventByPageResponse) GetBody() *ListAiAppRiskEventByPageResponseBody {
	return s.Body
}

func (s *ListAiAppRiskEventByPageResponse) SetHeaders(v map[string]*string) *ListAiAppRiskEventByPageResponse {
	s.Headers = v
	return s
}

func (s *ListAiAppRiskEventByPageResponse) SetStatusCode(v int32) *ListAiAppRiskEventByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponse) SetBody(v *ListAiAppRiskEventByPageResponseBody) *ListAiAppRiskEventByPageResponse {
	s.Body = v
	return s
}

func (s *ListAiAppRiskEventByPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
