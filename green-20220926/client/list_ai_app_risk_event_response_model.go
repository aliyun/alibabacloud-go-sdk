// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppRiskEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAiAppRiskEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAiAppRiskEventResponse
	GetStatusCode() *int32
	SetBody(v *ListAiAppRiskEventResponseBody) *ListAiAppRiskEventResponse
	GetBody() *ListAiAppRiskEventResponseBody
}

type ListAiAppRiskEventResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiAppRiskEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAiAppRiskEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppRiskEventResponse) GoString() string {
	return s.String()
}

func (s *ListAiAppRiskEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAiAppRiskEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAiAppRiskEventResponse) GetBody() *ListAiAppRiskEventResponseBody {
	return s.Body
}

func (s *ListAiAppRiskEventResponse) SetHeaders(v map[string]*string) *ListAiAppRiskEventResponse {
	s.Headers = v
	return s
}

func (s *ListAiAppRiskEventResponse) SetStatusCode(v int32) *ListAiAppRiskEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiAppRiskEventResponse) SetBody(v *ListAiAppRiskEventResponseBody) *ListAiAppRiskEventResponse {
	s.Body = v
	return s
}

func (s *ListAiAppRiskEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
