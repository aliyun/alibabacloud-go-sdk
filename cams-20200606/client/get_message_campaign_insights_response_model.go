// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCampaignInsightsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageCampaignInsightsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageCampaignInsightsResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageCampaignInsightsResponseBody) *GetMessageCampaignInsightsResponse
	GetBody() *GetMessageCampaignInsightsResponseBody
}

type GetMessageCampaignInsightsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageCampaignInsightsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageCampaignInsightsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCampaignInsightsResponse) GoString() string {
	return s.String()
}

func (s *GetMessageCampaignInsightsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageCampaignInsightsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageCampaignInsightsResponse) GetBody() *GetMessageCampaignInsightsResponseBody {
	return s.Body
}

func (s *GetMessageCampaignInsightsResponse) SetHeaders(v map[string]*string) *GetMessageCampaignInsightsResponse {
	s.Headers = v
	return s
}

func (s *GetMessageCampaignInsightsResponse) SetStatusCode(v int32) *GetMessageCampaignInsightsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageCampaignInsightsResponse) SetBody(v *GetMessageCampaignInsightsResponseBody) *GetMessageCampaignInsightsResponse {
	s.Body = v
	return s
}

func (s *GetMessageCampaignInsightsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
