// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMessageCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMessageCampaignResponse
	GetStatusCode() *int32
	SetBody(v *CreateMessageCampaignResponseBody) *CreateMessageCampaignResponse
	GetBody() *CreateMessageCampaignResponseBody
}

type CreateMessageCampaignResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMessageCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMessageCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageCampaignResponse) GoString() string {
	return s.String()
}

func (s *CreateMessageCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMessageCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMessageCampaignResponse) GetBody() *CreateMessageCampaignResponseBody {
	return s.Body
}

func (s *CreateMessageCampaignResponse) SetHeaders(v map[string]*string) *CreateMessageCampaignResponse {
	s.Headers = v
	return s
}

func (s *CreateMessageCampaignResponse) SetStatusCode(v int32) *CreateMessageCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMessageCampaignResponse) SetBody(v *CreateMessageCampaignResponseBody) *CreateMessageCampaignResponse {
	s.Body = v
	return s
}

func (s *CreateMessageCampaignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
