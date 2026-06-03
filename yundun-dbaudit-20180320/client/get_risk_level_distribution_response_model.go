// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskLevelDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRiskLevelDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRiskLevelDistributionResponse
	GetStatusCode() *int32
	SetBody(v *GetRiskLevelDistributionResponseBody) *GetRiskLevelDistributionResponse
	GetBody() *GetRiskLevelDistributionResponseBody
}

type GetRiskLevelDistributionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRiskLevelDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRiskLevelDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRiskLevelDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetRiskLevelDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRiskLevelDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRiskLevelDistributionResponse) GetBody() *GetRiskLevelDistributionResponseBody {
	return s.Body
}

func (s *GetRiskLevelDistributionResponse) SetHeaders(v map[string]*string) *GetRiskLevelDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetRiskLevelDistributionResponse) SetStatusCode(v int32) *GetRiskLevelDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRiskLevelDistributionResponse) SetBody(v *GetRiskLevelDistributionResponseBody) *GetRiskLevelDistributionResponse {
	s.Body = v
	return s
}

func (s *GetRiskLevelDistributionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
