// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRepositoryMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRepositoryMemberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRepositoryMemberResponseBody) *DeleteRepositoryMemberResponse
	GetBody() *DeleteRepositoryMemberResponseBody
}

type DeleteRepositoryMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepositoryMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRepositoryMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRepositoryMemberResponse) GetBody() *DeleteRepositoryMemberResponseBody {
	return s.Body
}

func (s *DeleteRepositoryMemberResponse) SetHeaders(v map[string]*string) *DeleteRepositoryMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryMemberResponse) SetStatusCode(v int32) *DeleteRepositoryMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryMemberResponse) SetBody(v *DeleteRepositoryMemberResponseBody) *DeleteRepositoryMemberResponse {
	s.Body = v
	return s
}

func (s *DeleteRepositoryMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
