// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommitFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommitFileResponse
	GetStatusCode() *int32
	SetBody(v *CommitFileResponseBody) *CommitFileResponse
	GetBody() *CommitFileResponseBody
}

type CommitFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommitFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommitFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CommitFileResponse) GoString() string {
	return s.String()
}

func (s *CommitFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommitFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommitFileResponse) GetBody() *CommitFileResponseBody {
	return s.Body
}

func (s *CommitFileResponse) SetHeaders(v map[string]*string) *CommitFileResponse {
	s.Headers = v
	return s
}

func (s *CommitFileResponse) SetStatusCode(v int32) *CommitFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitFileResponse) SetBody(v *CommitFileResponseBody) *CommitFileResponse {
	s.Body = v
	return s
}

func (s *CommitFileResponse) Validate() error {
	return dara.Validate(s)
}
