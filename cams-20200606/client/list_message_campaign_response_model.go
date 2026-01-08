// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessageCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessageCampaignResponse
	GetStatusCode() *int32
	SetBody(v *ListMessageCampaignResponseBody) *ListMessageCampaignResponse
	GetBody() *ListMessageCampaignResponseBody
}

type ListMessageCampaignResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessageCampaignResponse) GoString() string {
	return s.String()
}

func (s *ListMessageCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessageCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessageCampaignResponse) GetBody() *ListMessageCampaignResponseBody {
	return s.Body
}

func (s *ListMessageCampaignResponse) SetHeaders(v map[string]*string) *ListMessageCampaignResponse {
	s.Headers = v
	return s
}

func (s *ListMessageCampaignResponse) SetStatusCode(v int32) *ListMessageCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageCampaignResponse) SetBody(v *ListMessageCampaignResponseBody) *ListMessageCampaignResponse {
	s.Body = v
	return s
}

func (s *ListMessageCampaignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
