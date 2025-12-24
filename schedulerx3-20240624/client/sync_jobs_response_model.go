// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncJobsResponse
	GetStatusCode() *int32
	SetBody(v *SyncJobsResponseBody) *SyncJobsResponse
	GetBody() *SyncJobsResponseBody
}

type SyncJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncJobsResponse) GoString() string {
	return s.String()
}

func (s *SyncJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncJobsResponse) GetBody() *SyncJobsResponseBody {
	return s.Body
}

func (s *SyncJobsResponse) SetHeaders(v map[string]*string) *SyncJobsResponse {
	s.Headers = v
	return s
}

func (s *SyncJobsResponse) SetStatusCode(v int32) *SyncJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncJobsResponse) SetBody(v *SyncJobsResponseBody) *SyncJobsResponse {
	s.Body = v
	return s
}

func (s *SyncJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
