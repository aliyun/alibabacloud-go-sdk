// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCampaignTrendingReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCampaignTrendingReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCampaignTrendingReportResponse
	GetStatusCode() *int32
	SetBody(v *ListCampaignTrendingReportResponseBody) *ListCampaignTrendingReportResponse
	GetBody() *ListCampaignTrendingReportResponseBody
}

type ListCampaignTrendingReportResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCampaignTrendingReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCampaignTrendingReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignTrendingReportResponse) GoString() string {
	return s.String()
}

func (s *ListCampaignTrendingReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCampaignTrendingReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCampaignTrendingReportResponse) GetBody() *ListCampaignTrendingReportResponseBody {
	return s.Body
}

func (s *ListCampaignTrendingReportResponse) SetHeaders(v map[string]*string) *ListCampaignTrendingReportResponse {
	s.Headers = v
	return s
}

func (s *ListCampaignTrendingReportResponse) SetStatusCode(v int32) *ListCampaignTrendingReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCampaignTrendingReportResponse) SetBody(v *ListCampaignTrendingReportResponseBody) *ListCampaignTrendingReportResponse {
	s.Body = v
	return s
}

func (s *ListCampaignTrendingReportResponse) Validate() error {
	return dara.Validate(s)
}
