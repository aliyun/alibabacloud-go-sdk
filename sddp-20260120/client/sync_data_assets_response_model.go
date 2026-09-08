// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDataAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncDataAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncDataAssetsResponse
	GetStatusCode() *int32
	SetBody(v *SyncDataAssetsResponseBody) *SyncDataAssetsResponse
	GetBody() *SyncDataAssetsResponseBody
}

type SyncDataAssetsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDataAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDataAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncDataAssetsResponse) GoString() string {
	return s.String()
}

func (s *SyncDataAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncDataAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncDataAssetsResponse) GetBody() *SyncDataAssetsResponseBody {
	return s.Body
}

func (s *SyncDataAssetsResponse) SetHeaders(v map[string]*string) *SyncDataAssetsResponse {
	s.Headers = v
	return s
}

func (s *SyncDataAssetsResponse) SetStatusCode(v int32) *SyncDataAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDataAssetsResponse) SetBody(v *SyncDataAssetsResponseBody) *SyncDataAssetsResponse {
	s.Body = v
	return s
}

func (s *SyncDataAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
