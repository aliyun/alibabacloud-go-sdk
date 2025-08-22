// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitStagingRoutineCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommitStagingRoutineCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommitStagingRoutineCodeResponse
	GetStatusCode() *int32
	SetBody(v *CommitStagingRoutineCodeResponseBody) *CommitStagingRoutineCodeResponse
	GetBody() *CommitStagingRoutineCodeResponseBody
}

type CommitStagingRoutineCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommitStagingRoutineCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommitStagingRoutineCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CommitStagingRoutineCodeResponse) GoString() string {
	return s.String()
}

func (s *CommitStagingRoutineCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommitStagingRoutineCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommitStagingRoutineCodeResponse) GetBody() *CommitStagingRoutineCodeResponseBody {
	return s.Body
}

func (s *CommitStagingRoutineCodeResponse) SetHeaders(v map[string]*string) *CommitStagingRoutineCodeResponse {
	s.Headers = v
	return s
}

func (s *CommitStagingRoutineCodeResponse) SetStatusCode(v int32) *CommitStagingRoutineCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitStagingRoutineCodeResponse) SetBody(v *CommitStagingRoutineCodeResponseBody) *CommitStagingRoutineCodeResponse {
	s.Body = v
	return s
}

func (s *CommitStagingRoutineCodeResponse) Validate() error {
	return dara.Validate(s)
}
