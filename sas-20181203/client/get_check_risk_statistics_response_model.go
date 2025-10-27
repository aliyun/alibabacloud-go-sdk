// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckRiskStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckRiskStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckRiskStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckRiskStatisticsResponseBody) *GetCheckRiskStatisticsResponse
	GetBody() *GetCheckRiskStatisticsResponseBody
}

type GetCheckRiskStatisticsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckRiskStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckRiskStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRiskStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetCheckRiskStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckRiskStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckRiskStatisticsResponse) GetBody() *GetCheckRiskStatisticsResponseBody {
	return s.Body
}

func (s *GetCheckRiskStatisticsResponse) SetHeaders(v map[string]*string) *GetCheckRiskStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetCheckRiskStatisticsResponse) SetStatusCode(v int32) *GetCheckRiskStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckRiskStatisticsResponse) SetBody(v *GetCheckRiskStatisticsResponseBody) *GetCheckRiskStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetCheckRiskStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
