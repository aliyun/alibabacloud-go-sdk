// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMessageCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMessageCampaignResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMessageCampaignResponseBody) *DeleteMessageCampaignResponse
	GetBody() *DeleteMessageCampaignResponseBody
}

type DeleteMessageCampaignResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessageCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessageCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageCampaignResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMessageCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMessageCampaignResponse) GetBody() *DeleteMessageCampaignResponseBody {
	return s.Body
}

func (s *DeleteMessageCampaignResponse) SetHeaders(v map[string]*string) *DeleteMessageCampaignResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageCampaignResponse) SetStatusCode(v int32) *DeleteMessageCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageCampaignResponse) SetBody(v *DeleteMessageCampaignResponseBody) *DeleteMessageCampaignResponse {
	s.Body = v
	return s
}

func (s *DeleteMessageCampaignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
