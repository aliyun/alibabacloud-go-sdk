// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncRCKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncRCKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncRCKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *SyncRCKeyPairResponseBody) *SyncRCKeyPairResponse
	GetBody() *SyncRCKeyPairResponseBody
}

type SyncRCKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncRCKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncRCKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncRCKeyPairResponse) GoString() string {
	return s.String()
}

func (s *SyncRCKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncRCKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncRCKeyPairResponse) GetBody() *SyncRCKeyPairResponseBody {
	return s.Body
}

func (s *SyncRCKeyPairResponse) SetHeaders(v map[string]*string) *SyncRCKeyPairResponse {
	s.Headers = v
	return s
}

func (s *SyncRCKeyPairResponse) SetStatusCode(v int32) *SyncRCKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncRCKeyPairResponse) SetBody(v *SyncRCKeyPairResponseBody) *SyncRCKeyPairResponse {
	s.Body = v
	return s
}

func (s *SyncRCKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
