// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitRoutineStagingCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommitRoutineStagingCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommitRoutineStagingCodeResponse
	GetStatusCode() *int32
	SetBody(v *CommitRoutineStagingCodeResponseBody) *CommitRoutineStagingCodeResponse
	GetBody() *CommitRoutineStagingCodeResponseBody
}

type CommitRoutineStagingCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommitRoutineStagingCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommitRoutineStagingCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CommitRoutineStagingCodeResponse) GoString() string {
	return s.String()
}

func (s *CommitRoutineStagingCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommitRoutineStagingCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommitRoutineStagingCodeResponse) GetBody() *CommitRoutineStagingCodeResponseBody {
	return s.Body
}

func (s *CommitRoutineStagingCodeResponse) SetHeaders(v map[string]*string) *CommitRoutineStagingCodeResponse {
	s.Headers = v
	return s
}

func (s *CommitRoutineStagingCodeResponse) SetStatusCode(v int32) *CommitRoutineStagingCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitRoutineStagingCodeResponse) SetBody(v *CommitRoutineStagingCodeResponseBody) *CommitRoutineStagingCodeResponse {
	s.Body = v
	return s
}

func (s *CommitRoutineStagingCodeResponse) Validate() error {
	return dara.Validate(s)
}
