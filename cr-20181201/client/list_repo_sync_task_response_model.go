// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoSyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepoSyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepoSyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListRepoSyncTaskResponseBody) *ListRepoSyncTaskResponse
	GetBody() *ListRepoSyncTaskResponseBody
}

type ListRepoSyncTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoSyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepoSyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepoSyncTaskResponse) GetBody() *ListRepoSyncTaskResponseBody {
	return s.Body
}

func (s *ListRepoSyncTaskResponse) SetHeaders(v map[string]*string) *ListRepoSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *ListRepoSyncTaskResponse) SetStatusCode(v int32) *ListRepoSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoSyncTaskResponse) SetBody(v *ListRepoSyncTaskResponseBody) *ListRepoSyncTaskResponse {
	s.Body = v
	return s
}

func (s *ListRepoSyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
