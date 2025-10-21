// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMemberResponse
	GetStatusCode() *int32
	SetBody(v *CreateMemberResponseBody) *CreateMemberResponse
	GetBody() *CreateMemberResponseBody
}

type CreateMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMemberResponse) GetBody() *CreateMemberResponseBody {
	return s.Body
}

func (s *CreateMemberResponse) SetHeaders(v map[string]*string) *CreateMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateMemberResponse) SetStatusCode(v int32) *CreateMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemberResponse) SetBody(v *CreateMemberResponseBody) *CreateMemberResponse {
	s.Body = v
	return s
}

func (s *CreateMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
