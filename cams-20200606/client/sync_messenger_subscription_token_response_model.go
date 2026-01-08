// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMessengerSubscriptionTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncMessengerSubscriptionTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncMessengerSubscriptionTokenResponse
	GetStatusCode() *int32
	SetBody(v *SyncMessengerSubscriptionTokenResponseBody) *SyncMessengerSubscriptionTokenResponse
	GetBody() *SyncMessengerSubscriptionTokenResponseBody
}

type SyncMessengerSubscriptionTokenResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncMessengerSubscriptionTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncMessengerSubscriptionTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncMessengerSubscriptionTokenResponse) GoString() string {
	return s.String()
}

func (s *SyncMessengerSubscriptionTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncMessengerSubscriptionTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncMessengerSubscriptionTokenResponse) GetBody() *SyncMessengerSubscriptionTokenResponseBody {
	return s.Body
}

func (s *SyncMessengerSubscriptionTokenResponse) SetHeaders(v map[string]*string) *SyncMessengerSubscriptionTokenResponse {
	s.Headers = v
	return s
}

func (s *SyncMessengerSubscriptionTokenResponse) SetStatusCode(v int32) *SyncMessengerSubscriptionTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenResponse) SetBody(v *SyncMessengerSubscriptionTokenResponseBody) *SyncMessengerSubscriptionTokenResponse {
	s.Body = v
	return s
}

func (s *SyncMessengerSubscriptionTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
