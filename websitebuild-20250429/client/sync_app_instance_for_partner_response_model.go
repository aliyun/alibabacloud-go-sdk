// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncAppInstanceForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncAppInstanceForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncAppInstanceForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *SyncAppInstanceForPartnerResponseBody) *SyncAppInstanceForPartnerResponse
	GetBody() *SyncAppInstanceForPartnerResponseBody
}

type SyncAppInstanceForPartnerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncAppInstanceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncAppInstanceForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncAppInstanceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *SyncAppInstanceForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncAppInstanceForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncAppInstanceForPartnerResponse) GetBody() *SyncAppInstanceForPartnerResponseBody {
	return s.Body
}

func (s *SyncAppInstanceForPartnerResponse) SetHeaders(v map[string]*string) *SyncAppInstanceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *SyncAppInstanceForPartnerResponse) SetStatusCode(v int32) *SyncAppInstanceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncAppInstanceForPartnerResponse) SetBody(v *SyncAppInstanceForPartnerResponseBody) *SyncAppInstanceForPartnerResponse {
	s.Body = v
	return s
}

func (s *SyncAppInstanceForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
