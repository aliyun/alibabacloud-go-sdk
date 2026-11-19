// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMiniGameOnlineStatusByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncMiniGameOnlineStatusByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncMiniGameOnlineStatusByAppResponse
	GetStatusCode() *int32
	SetBody(v *SyncMiniGameOnlineStatusByAppResponseBody) *SyncMiniGameOnlineStatusByAppResponse
	GetBody() *SyncMiniGameOnlineStatusByAppResponseBody
}

type SyncMiniGameOnlineStatusByAppResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncMiniGameOnlineStatusByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncMiniGameOnlineStatusByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncMiniGameOnlineStatusByAppResponse) GoString() string {
	return s.String()
}

func (s *SyncMiniGameOnlineStatusByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncMiniGameOnlineStatusByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncMiniGameOnlineStatusByAppResponse) GetBody() *SyncMiniGameOnlineStatusByAppResponseBody {
	return s.Body
}

func (s *SyncMiniGameOnlineStatusByAppResponse) SetHeaders(v map[string]*string) *SyncMiniGameOnlineStatusByAppResponse {
	s.Headers = v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponse) SetStatusCode(v int32) *SyncMiniGameOnlineStatusByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponse) SetBody(v *SyncMiniGameOnlineStatusByAppResponseBody) *SyncMiniGameOnlineStatusByAppResponse {
	s.Body = v
	return s
}

func (s *SyncMiniGameOnlineStatusByAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
