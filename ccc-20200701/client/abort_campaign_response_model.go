// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbortCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbortCampaignResponse
	GetStatusCode() *int32
	SetBody(v *AbortCampaignResponseBody) *AbortCampaignResponse
	GetBody() *AbortCampaignResponseBody
}

type AbortCampaignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbortCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbortCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s AbortCampaignResponse) GoString() string {
	return s.String()
}

func (s *AbortCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbortCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbortCampaignResponse) GetBody() *AbortCampaignResponseBody {
	return s.Body
}

func (s *AbortCampaignResponse) SetHeaders(v map[string]*string) *AbortCampaignResponse {
	s.Headers = v
	return s
}

func (s *AbortCampaignResponse) SetStatusCode(v int32) *AbortCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortCampaignResponse) SetBody(v *AbortCampaignResponseBody) *AbortCampaignResponse {
	s.Body = v
	return s
}

func (s *AbortCampaignResponse) Validate() error {
	return dara.Validate(s)
}
