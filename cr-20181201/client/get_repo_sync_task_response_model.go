// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoSyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRepoSyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRepoSyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetRepoSyncTaskResponseBody) *GetRepoSyncTaskResponse
	GetBody() *GetRepoSyncTaskResponseBody
}

type GetRepoSyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoSyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRepoSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRepoSyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRepoSyncTaskResponse) GetBody() *GetRepoSyncTaskResponseBody {
	return s.Body
}

func (s *GetRepoSyncTaskResponse) SetHeaders(v map[string]*string) *GetRepoSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetRepoSyncTaskResponse) SetStatusCode(v int32) *GetRepoSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoSyncTaskResponse) SetBody(v *GetRepoSyncTaskResponseBody) *GetRepoSyncTaskResponse {
	s.Body = v
	return s
}

func (s *GetRepoSyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
