// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolSyncJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserPoolSyncJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserPoolSyncJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserPoolSyncJobsResponseBody) *ListUserPoolSyncJobsResponse
	GetBody() *ListUserPoolSyncJobsResponseBody
}

type ListUserPoolSyncJobsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserPoolSyncJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserPoolSyncJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolSyncJobsResponse) GoString() string {
	return s.String()
}

func (s *ListUserPoolSyncJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserPoolSyncJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserPoolSyncJobsResponse) GetBody() *ListUserPoolSyncJobsResponseBody {
	return s.Body
}

func (s *ListUserPoolSyncJobsResponse) SetHeaders(v map[string]*string) *ListUserPoolSyncJobsResponse {
	s.Headers = v
	return s
}

func (s *ListUserPoolSyncJobsResponse) SetStatusCode(v int32) *ListUserPoolSyncJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserPoolSyncJobsResponse) SetBody(v *ListUserPoolSyncJobsResponseBody) *ListUserPoolSyncJobsResponse {
	s.Body = v
	return s
}

func (s *ListUserPoolSyncJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
