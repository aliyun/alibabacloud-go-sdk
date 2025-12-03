// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommitStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCommitStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCommitStatusResponse
	GetStatusCode() *int32
	SetBody(v *CreateCommitStatusResponseBody) *CreateCommitStatusResponse
	GetBody() *CreateCommitStatusResponseBody
}

type CreateCommitStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCommitStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCommitStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitStatusResponse) GoString() string {
	return s.String()
}

func (s *CreateCommitStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCommitStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCommitStatusResponse) GetBody() *CreateCommitStatusResponseBody {
	return s.Body
}

func (s *CreateCommitStatusResponse) SetHeaders(v map[string]*string) *CreateCommitStatusResponse {
	s.Headers = v
	return s
}

func (s *CreateCommitStatusResponse) SetStatusCode(v int32) *CreateCommitStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCommitStatusResponse) SetBody(v *CreateCommitStatusResponseBody) *CreateCommitStatusResponse {
	s.Body = v
	return s
}

func (s *CreateCommitStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
