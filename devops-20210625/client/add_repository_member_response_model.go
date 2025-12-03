// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRepositoryMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRepositoryMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRepositoryMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddRepositoryMemberResponseBody) *AddRepositoryMemberResponse
	GetBody() *AddRepositoryMemberResponseBody
}

type AddRepositoryMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRepositoryMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRepositoryMemberResponse) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRepositoryMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRepositoryMemberResponse) GetBody() *AddRepositoryMemberResponseBody {
	return s.Body
}

func (s *AddRepositoryMemberResponse) SetHeaders(v map[string]*string) *AddRepositoryMemberResponse {
	s.Headers = v
	return s
}

func (s *AddRepositoryMemberResponse) SetStatusCode(v int32) *AddRepositoryMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRepositoryMemberResponse) SetBody(v *AddRepositoryMemberResponseBody) *AddRepositoryMemberResponse {
	s.Body = v
	return s
}

func (s *AddRepositoryMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
