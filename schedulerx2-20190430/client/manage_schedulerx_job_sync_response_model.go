// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxJobSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManageSchedulerxJobSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManageSchedulerxJobSyncResponse
	GetStatusCode() *int32
	SetBody(v *ManageSchedulerxJobSyncResponseBody) *ManageSchedulerxJobSyncResponse
	GetBody() *ManageSchedulerxJobSyncResponseBody
}

type ManageSchedulerxJobSyncResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManageSchedulerxJobSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageSchedulerxJobSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxJobSyncResponse) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxJobSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManageSchedulerxJobSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManageSchedulerxJobSyncResponse) GetBody() *ManageSchedulerxJobSyncResponseBody {
	return s.Body
}

func (s *ManageSchedulerxJobSyncResponse) SetHeaders(v map[string]*string) *ManageSchedulerxJobSyncResponse {
	s.Headers = v
	return s
}

func (s *ManageSchedulerxJobSyncResponse) SetStatusCode(v int32) *ManageSchedulerxJobSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageSchedulerxJobSyncResponse) SetBody(v *ManageSchedulerxJobSyncResponseBody) *ManageSchedulerxJobSyncResponse {
	s.Body = v
	return s
}

func (s *ManageSchedulerxJobSyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
