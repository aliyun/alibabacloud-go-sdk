// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseCampaignResponse
	GetStatusCode() *int32
	SetBody(v *PauseCampaignResponseBody) *PauseCampaignResponse
	GetBody() *PauseCampaignResponseBody
}

type PauseCampaignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseCampaignResponse) GoString() string {
	return s.String()
}

func (s *PauseCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseCampaignResponse) GetBody() *PauseCampaignResponseBody {
	return s.Body
}

func (s *PauseCampaignResponse) SetHeaders(v map[string]*string) *PauseCampaignResponse {
	s.Headers = v
	return s
}

func (s *PauseCampaignResponse) SetStatusCode(v int32) *PauseCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseCampaignResponse) SetBody(v *PauseCampaignResponseBody) *PauseCampaignResponse {
	s.Body = v
	return s
}

func (s *PauseCampaignResponse) Validate() error {
	return dara.Validate(s)
}
