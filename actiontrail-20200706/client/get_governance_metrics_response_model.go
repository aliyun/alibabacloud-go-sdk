// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGovernanceMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGovernanceMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetGovernanceMetricsResponseBody) *GetGovernanceMetricsResponse
	GetBody() *GetGovernanceMetricsResponseBody
}

type GetGovernanceMetricsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGovernanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGovernanceMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetGovernanceMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGovernanceMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGovernanceMetricsResponse) GetBody() *GetGovernanceMetricsResponseBody {
	return s.Body
}

func (s *GetGovernanceMetricsResponse) SetHeaders(v map[string]*string) *GetGovernanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetGovernanceMetricsResponse) SetStatusCode(v int32) *GetGovernanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGovernanceMetricsResponse) SetBody(v *GetGovernanceMetricsResponseBody) *GetGovernanceMetricsResponse {
	s.Body = v
	return s
}

func (s *GetGovernanceMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
