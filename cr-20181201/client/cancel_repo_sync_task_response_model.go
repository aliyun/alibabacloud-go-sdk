// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRepoSyncTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelRepoSyncTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelRepoSyncTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelRepoSyncTaskResponseBody) *CancelRepoSyncTaskResponse
	GetBody() *CancelRepoSyncTaskResponseBody
}

type CancelRepoSyncTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRepoSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRepoSyncTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelRepoSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelRepoSyncTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelRepoSyncTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelRepoSyncTaskResponse) GetBody() *CancelRepoSyncTaskResponseBody {
	return s.Body
}

func (s *CancelRepoSyncTaskResponse) SetHeaders(v map[string]*string) *CancelRepoSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelRepoSyncTaskResponse) SetStatusCode(v int32) *CancelRepoSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRepoSyncTaskResponse) SetBody(v *CancelRepoSyncTaskResponseBody) *CancelRepoSyncTaskResponse {
	s.Body = v
	return s
}

func (s *CancelRepoSyncTaskResponse) Validate() error {
	return dara.Validate(s)
}
