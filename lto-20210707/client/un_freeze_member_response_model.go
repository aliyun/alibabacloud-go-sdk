// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnFreezeMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnFreezeMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnFreezeMemberResponse
	GetStatusCode() *int32
	SetBody(v *UnFreezeMemberResponseBody) *UnFreezeMemberResponse
	GetBody() *UnFreezeMemberResponseBody
}

type UnFreezeMemberResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnFreezeMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnFreezeMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UnFreezeMemberResponse) GoString() string {
	return s.String()
}

func (s *UnFreezeMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnFreezeMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnFreezeMemberResponse) GetBody() *UnFreezeMemberResponseBody {
	return s.Body
}

func (s *UnFreezeMemberResponse) SetHeaders(v map[string]*string) *UnFreezeMemberResponse {
	s.Headers = v
	return s
}

func (s *UnFreezeMemberResponse) SetStatusCode(v int32) *UnFreezeMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UnFreezeMemberResponse) SetBody(v *UnFreezeMemberResponseBody) *UnFreezeMemberResponse {
	s.Body = v
	return s
}

func (s *UnFreezeMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
