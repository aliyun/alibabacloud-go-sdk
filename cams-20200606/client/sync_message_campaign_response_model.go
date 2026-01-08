// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMessageCampaignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncMessageCampaignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncMessageCampaignResponse
	GetStatusCode() *int32
	SetBody(v *SyncMessageCampaignResponseBody) *SyncMessageCampaignResponse
	GetBody() *SyncMessageCampaignResponseBody
}

type SyncMessageCampaignResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncMessageCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncMessageCampaignResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncMessageCampaignResponse) GoString() string {
	return s.String()
}

func (s *SyncMessageCampaignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncMessageCampaignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncMessageCampaignResponse) GetBody() *SyncMessageCampaignResponseBody {
	return s.Body
}

func (s *SyncMessageCampaignResponse) SetHeaders(v map[string]*string) *SyncMessageCampaignResponse {
	s.Headers = v
	return s
}

func (s *SyncMessageCampaignResponse) SetStatusCode(v int32) *SyncMessageCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncMessageCampaignResponse) SetBody(v *SyncMessageCampaignResponseBody) *SyncMessageCampaignResponse {
	s.Body = v
	return s
}

func (s *SyncMessageCampaignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
