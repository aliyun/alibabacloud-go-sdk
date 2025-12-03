// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommitStatusesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCommitStatusesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCommitStatusesResponse
	GetStatusCode() *int32
	SetBody(v *ListCommitStatusesResponseBody) *ListCommitStatusesResponse
	GetBody() *ListCommitStatusesResponseBody
}

type ListCommitStatusesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommitStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommitStatusesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCommitStatusesResponse) GoString() string {
	return s.String()
}

func (s *ListCommitStatusesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCommitStatusesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCommitStatusesResponse) GetBody() *ListCommitStatusesResponseBody {
	return s.Body
}

func (s *ListCommitStatusesResponse) SetHeaders(v map[string]*string) *ListCommitStatusesResponse {
	s.Headers = v
	return s
}

func (s *ListCommitStatusesResponse) SetStatusCode(v int32) *ListCommitStatusesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommitStatusesResponse) SetBody(v *ListCommitStatusesResponseBody) *ListCommitStatusesResponse {
	s.Body = v
	return s
}

func (s *ListCommitStatusesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
