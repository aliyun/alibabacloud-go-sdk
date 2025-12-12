// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncQualityCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncQualityCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncQualityCheckResponse
	GetStatusCode() *int32
	SetBody(v *SyncQualityCheckResponseBody) *SyncQualityCheckResponse
	GetBody() *SyncQualityCheckResponseBody
}

type SyncQualityCheckResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncQualityCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncQualityCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncQualityCheckResponse) GoString() string {
	return s.String()
}

func (s *SyncQualityCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncQualityCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncQualityCheckResponse) GetBody() *SyncQualityCheckResponseBody {
	return s.Body
}

func (s *SyncQualityCheckResponse) SetHeaders(v map[string]*string) *SyncQualityCheckResponse {
	s.Headers = v
	return s
}

func (s *SyncQualityCheckResponse) SetStatusCode(v int32) *SyncQualityCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncQualityCheckResponse) SetBody(v *SyncQualityCheckResponseBody) *SyncQualityCheckResponse {
	s.Body = v
	return s
}

func (s *SyncQualityCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
