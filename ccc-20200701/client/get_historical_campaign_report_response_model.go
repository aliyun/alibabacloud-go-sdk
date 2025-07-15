// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalCampaignReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHistoricalCampaignReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHistoricalCampaignReportResponse
	GetStatusCode() *int32
	SetBody(v *GetHistoricalCampaignReportResponseBody) *GetHistoricalCampaignReportResponse
	GetBody() *GetHistoricalCampaignReportResponseBody
}

type GetHistoricalCampaignReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoricalCampaignReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoricalCampaignReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalCampaignReportResponse) GoString() string {
	return s.String()
}

func (s *GetHistoricalCampaignReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHistoricalCampaignReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHistoricalCampaignReportResponse) GetBody() *GetHistoricalCampaignReportResponseBody {
	return s.Body
}

func (s *GetHistoricalCampaignReportResponse) SetHeaders(v map[string]*string) *GetHistoricalCampaignReportResponse {
	s.Headers = v
	return s
}

func (s *GetHistoricalCampaignReportResponse) SetStatusCode(v int32) *GetHistoricalCampaignReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoricalCampaignReportResponse) SetBody(v *GetHistoricalCampaignReportResponseBody) *GetHistoricalCampaignReportResponse {
	s.Body = v
	return s
}

func (s *GetHistoricalCampaignReportResponse) Validate() error {
	return dara.Validate(s)
}
