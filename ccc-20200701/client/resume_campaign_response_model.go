// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeCampaignResponse
	GetStatusCode() *int32
	SetBody(v *ResumeCampaignResponseBody) *ResumeCampaignResponse
	GetBody() *ResumeCampaignResponseBody
}

type ResumeCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeCampaignResponse) GoString() string {
	return s.String()
}

func (s *ResumeCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeCampaignResponse) GetBody() *ResumeCampaignResponseBody {
	return s.Body
}

func (s *ResumeCampaignResponse) SetHeaders(v map[string]*string) *ResumeCampaignResponse {
	s.Headers = v
	return s
}

func (s *ResumeCampaignResponse) SetStatusCode(v int32) *ResumeCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeCampaignResponse) SetBody(v *ResumeCampaignResponseBody) *ResumeCampaignResponse {
	s.Body = v
	return s
}

func (s *ResumeCampaignResponse) Validate() error {
	return dara.Validate(s)
}
