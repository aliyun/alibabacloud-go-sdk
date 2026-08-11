// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRiskStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRiskStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetRiskStatsResponseBody) *GetRiskStatsResponse
	GetBody() *GetRiskStatsResponseBody
}

type GetRiskStatsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRiskStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRiskStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRiskStatsResponse) GoString() string {
	return s.String()
}

func (s *GetRiskStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRiskStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRiskStatsResponse) GetBody() *GetRiskStatsResponseBody {
	return s.Body
}

func (s *GetRiskStatsResponse) SetHeaders(v map[string]*string) *GetRiskStatsResponse {
	s.Headers = v
	return s
}

func (s *GetRiskStatsResponse) SetStatusCode(v int32) *GetRiskStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRiskStatsResponse) SetBody(v *GetRiskStatsResponseBody) *GetRiskStatsResponse {
	s.Body = v
	return s
}

func (s *GetRiskStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
