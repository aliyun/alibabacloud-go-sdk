// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommitWithMultipleFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCommitWithMultipleFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCommitWithMultipleFilesResponse
	GetStatusCode() *int32
	SetBody(v *CreateCommitWithMultipleFilesResponseBody) *CreateCommitWithMultipleFilesResponse
	GetBody() *CreateCommitWithMultipleFilesResponseBody
}

type CreateCommitWithMultipleFilesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCommitWithMultipleFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCommitWithMultipleFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitWithMultipleFilesResponse) GoString() string {
	return s.String()
}

func (s *CreateCommitWithMultipleFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCommitWithMultipleFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCommitWithMultipleFilesResponse) GetBody() *CreateCommitWithMultipleFilesResponseBody {
	return s.Body
}

func (s *CreateCommitWithMultipleFilesResponse) SetHeaders(v map[string]*string) *CreateCommitWithMultipleFilesResponse {
	s.Headers = v
	return s
}

func (s *CreateCommitWithMultipleFilesResponse) SetStatusCode(v int32) *CreateCommitWithMultipleFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCommitWithMultipleFilesResponse) SetBody(v *CreateCommitWithMultipleFilesResponseBody) *CreateCommitWithMultipleFilesResponse {
	s.Body = v
	return s
}

func (s *CreateCommitWithMultipleFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
