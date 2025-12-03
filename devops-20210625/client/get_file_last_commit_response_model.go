// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileLastCommitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileLastCommitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileLastCommitResponse
	GetStatusCode() *int32
	SetBody(v *GetFileLastCommitResponseBody) *GetFileLastCommitResponse
	GetBody() *GetFileLastCommitResponseBody
}

type GetFileLastCommitResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileLastCommitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileLastCommitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileLastCommitResponse) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileLastCommitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileLastCommitResponse) GetBody() *GetFileLastCommitResponseBody {
	return s.Body
}

func (s *GetFileLastCommitResponse) SetHeaders(v map[string]*string) *GetFileLastCommitResponse {
	s.Headers = v
	return s
}

func (s *GetFileLastCommitResponse) SetStatusCode(v int32) *GetFileLastCommitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileLastCommitResponse) SetBody(v *GetFileLastCommitResponseBody) *GetFileLastCommitResponse {
	s.Body = v
	return s
}

func (s *GetFileLastCommitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
