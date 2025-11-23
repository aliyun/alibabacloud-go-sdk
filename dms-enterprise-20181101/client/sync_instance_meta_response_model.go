// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncInstanceMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncInstanceMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncInstanceMetaResponse
	GetStatusCode() *int32
	SetBody(v *SyncInstanceMetaResponseBody) *SyncInstanceMetaResponse
	GetBody() *SyncInstanceMetaResponseBody
}

type SyncInstanceMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncInstanceMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncInstanceMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncInstanceMetaResponse) GoString() string {
	return s.String()
}

func (s *SyncInstanceMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncInstanceMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncInstanceMetaResponse) GetBody() *SyncInstanceMetaResponseBody {
	return s.Body
}

func (s *SyncInstanceMetaResponse) SetHeaders(v map[string]*string) *SyncInstanceMetaResponse {
	s.Headers = v
	return s
}

func (s *SyncInstanceMetaResponse) SetStatusCode(v int32) *SyncInstanceMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncInstanceMetaResponse) SetBody(v *SyncInstanceMetaResponseBody) *SyncInstanceMetaResponse {
	s.Body = v
	return s
}

func (s *SyncInstanceMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
