// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCampaignResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCampaignResponseBody) *SubmitCampaignResponse
	GetBody() *SubmitCampaignResponseBody
}

type SubmitCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCampaignResponse) GoString() string {
	return s.String()
}

func (s *SubmitCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCampaignResponse) GetBody() *SubmitCampaignResponseBody {
	return s.Body
}

func (s *SubmitCampaignResponse) SetHeaders(v map[string]*string) *SubmitCampaignResponse {
	s.Headers = v
	return s
}

func (s *SubmitCampaignResponse) SetStatusCode(v int32) *SubmitCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCampaignResponse) SetBody(v *SubmitCampaignResponseBody) *SubmitCampaignResponse {
	s.Body = v
	return s
}

func (s *SubmitCampaignResponse) Validate() error {
	return dara.Validate(s)
}
