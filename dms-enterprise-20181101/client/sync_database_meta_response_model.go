// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDatabaseMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncDatabaseMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncDatabaseMetaResponse
	GetStatusCode() *int32
	SetBody(v *SyncDatabaseMetaResponseBody) *SyncDatabaseMetaResponse
	GetBody() *SyncDatabaseMetaResponseBody
}

type SyncDatabaseMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDatabaseMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDatabaseMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncDatabaseMetaResponse) GoString() string {
	return s.String()
}

func (s *SyncDatabaseMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncDatabaseMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncDatabaseMetaResponse) GetBody() *SyncDatabaseMetaResponseBody {
	return s.Body
}

func (s *SyncDatabaseMetaResponse) SetHeaders(v map[string]*string) *SyncDatabaseMetaResponse {
	s.Headers = v
	return s
}

func (s *SyncDatabaseMetaResponse) SetStatusCode(v int32) *SyncDatabaseMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDatabaseMetaResponse) SetBody(v *SyncDatabaseMetaResponseBody) *SyncDatabaseMetaResponse {
	s.Body = v
	return s
}

func (s *SyncDatabaseMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
