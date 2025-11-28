// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFreezeMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FreezeMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FreezeMemberResponse
	GetStatusCode() *int32
	SetBody(v *FreezeMemberResponseBody) *FreezeMemberResponse
	GetBody() *FreezeMemberResponseBody
}

type FreezeMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FreezeMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FreezeMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s FreezeMemberResponse) GoString() string {
	return s.String()
}

func (s *FreezeMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FreezeMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FreezeMemberResponse) GetBody() *FreezeMemberResponseBody {
	return s.Body
}

func (s *FreezeMemberResponse) SetHeaders(v map[string]*string) *FreezeMemberResponse {
	s.Headers = v
	return s
}

func (s *FreezeMemberResponse) SetStatusCode(v int32) *FreezeMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *FreezeMemberResponse) SetBody(v *FreezeMemberResponseBody) *FreezeMemberResponse {
	s.Body = v
	return s
}

func (s *FreezeMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
