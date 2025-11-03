// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoSyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepoSyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepoSyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepoSyncTaskResponseBody) *CreateRepoSyncTaskResponse
	GetBody() *CreateRepoSyncTaskResponseBody
}

type CreateRepoSyncTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoSyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepoSyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepoSyncTaskResponse) GetBody() *CreateRepoSyncTaskResponseBody {
	return s.Body
}

func (s *CreateRepoSyncTaskResponse) SetHeaders(v map[string]*string) *CreateRepoSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoSyncTaskResponse) SetStatusCode(v int32) *CreateRepoSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoSyncTaskResponse) SetBody(v *CreateRepoSyncTaskResponseBody) *CreateRepoSyncTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRepoSyncTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
