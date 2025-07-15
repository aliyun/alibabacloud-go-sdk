// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeCampaignStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRealtimeCampaignStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRealtimeCampaignStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetRealtimeCampaignStatsResponseBody) *GetRealtimeCampaignStatsResponse
	GetBody() *GetRealtimeCampaignStatsResponseBody
}

type GetRealtimeCampaignStatsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealtimeCampaignStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealtimeCampaignStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeCampaignStatsResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeCampaignStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRealtimeCampaignStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRealtimeCampaignStatsResponse) GetBody() *GetRealtimeCampaignStatsResponseBody {
	return s.Body
}

func (s *GetRealtimeCampaignStatsResponse) SetHeaders(v map[string]*string) *GetRealtimeCampaignStatsResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeCampaignStatsResponse) SetStatusCode(v int32) *GetRealtimeCampaignStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponse) SetBody(v *GetRealtimeCampaignStatsResponseBody) *GetRealtimeCampaignStatsResponse {
	s.Body = v
	return s
}

func (s *GetRealtimeCampaignStatsResponse) Validate() error {
	return dara.Validate(s)
}
