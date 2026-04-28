// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorCostCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshAdvisorCostCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshAdvisorCostCheckResponse
	GetStatusCode() *int32
	SetBody(v *RefreshAdvisorCostCheckResponseBody) *RefreshAdvisorCostCheckResponse
	GetBody() *RefreshAdvisorCostCheckResponseBody
}

type RefreshAdvisorCostCheckResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshAdvisorCostCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshAdvisorCostCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCostCheckResponse) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshAdvisorCostCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshAdvisorCostCheckResponse) GetBody() *RefreshAdvisorCostCheckResponseBody {
	return s.Body
}

func (s *RefreshAdvisorCostCheckResponse) SetHeaders(v map[string]*string) *RefreshAdvisorCostCheckResponse {
	s.Headers = v
	return s
}

func (s *RefreshAdvisorCostCheckResponse) SetStatusCode(v int32) *RefreshAdvisorCostCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponse) SetBody(v *RefreshAdvisorCostCheckResponseBody) *RefreshAdvisorCostCheckResponse {
	s.Body = v
	return s
}

func (s *RefreshAdvisorCostCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
