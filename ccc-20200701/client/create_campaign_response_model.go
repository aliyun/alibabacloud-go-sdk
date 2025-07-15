// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCampaignResponse
	GetStatusCode() *int32
	SetBody(v *CreateCampaignResponseBody) *CreateCampaignResponse
	GetBody() *CreateCampaignResponseBody
}

type CreateCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignResponse) GoString() string {
	return s.String()
}

func (s *CreateCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCampaignResponse) GetBody() *CreateCampaignResponseBody {
	return s.Body
}

func (s *CreateCampaignResponse) SetHeaders(v map[string]*string) *CreateCampaignResponse {
	s.Headers = v
	return s
}

func (s *CreateCampaignResponse) SetStatusCode(v int32) *CreateCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCampaignResponse) SetBody(v *CreateCampaignResponseBody) *CreateCampaignResponse {
	s.Body = v
	return s
}

func (s *CreateCampaignResponse) Validate() error {
	return dara.Validate(s)
}
