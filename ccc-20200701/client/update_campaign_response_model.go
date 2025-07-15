// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCampaignResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCampaignResponseBody) *UpdateCampaignResponse
	GetBody() *UpdateCampaignResponseBody
}

type UpdateCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCampaignResponse) GoString() string {
	return s.String()
}

func (s *UpdateCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCampaignResponse) GetBody() *UpdateCampaignResponseBody {
	return s.Body
}

func (s *UpdateCampaignResponse) SetHeaders(v map[string]*string) *UpdateCampaignResponse {
	s.Headers = v
	return s
}

func (s *UpdateCampaignResponse) SetStatusCode(v int32) *UpdateCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCampaignResponse) SetBody(v *UpdateCampaignResponseBody) *UpdateCampaignResponse {
	s.Body = v
	return s
}

func (s *UpdateCampaignResponse) Validate() error {
	return dara.Validate(s)
}
