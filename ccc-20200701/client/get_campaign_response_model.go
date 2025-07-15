// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCampaignResponse
	GetStatusCode() *int32
	SetBody(v *GetCampaignResponseBody) *GetCampaignResponse
	GetBody() *GetCampaignResponseBody
}

type GetCampaignResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCampaignResponse) GoString() string {
	return s.String()
}

func (s *GetCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCampaignResponse) GetBody() *GetCampaignResponseBody {
	return s.Body
}

func (s *GetCampaignResponse) SetHeaders(v map[string]*string) *GetCampaignResponse {
	s.Headers = v
	return s
}

func (s *GetCampaignResponse) SetStatusCode(v int32) *GetCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCampaignResponse) SetBody(v *GetCampaignResponseBody) *GetCampaignResponse {
	s.Body = v
	return s
}

func (s *GetCampaignResponse) Validate() error {
	return dara.Validate(s)
}
