// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommitServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommitServiceResponse
	GetStatusCode() *int32
	SetBody(v *CommitServiceResponseBody) *CommitServiceResponse
	GetBody() *CommitServiceResponseBody
}

type CommitServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommitServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommitServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CommitServiceResponse) GoString() string {
	return s.String()
}

func (s *CommitServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommitServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommitServiceResponse) GetBody() *CommitServiceResponseBody {
	return s.Body
}

func (s *CommitServiceResponse) SetHeaders(v map[string]*string) *CommitServiceResponse {
	s.Headers = v
	return s
}

func (s *CommitServiceResponse) SetStatusCode(v int32) *CommitServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CommitServiceResponse) SetBody(v *CommitServiceResponseBody) *CommitServiceResponse {
	s.Body = v
	return s
}

func (s *CommitServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
