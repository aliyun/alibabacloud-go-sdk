// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddMemberResponseBody) *AddMemberResponse
	GetBody() *AddMemberResponseBody
}

type AddMemberResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMemberResponse) GoString() string {
	return s.String()
}

func (s *AddMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMemberResponse) GetBody() *AddMemberResponseBody {
	return s.Body
}

func (s *AddMemberResponse) SetHeaders(v map[string]*string) *AddMemberResponse {
	s.Headers = v
	return s
}

func (s *AddMemberResponse) SetStatusCode(v int32) *AddMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMemberResponse) SetBody(v *AddMemberResponseBody) *AddMemberResponse {
	s.Body = v
	return s
}

func (s *AddMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
