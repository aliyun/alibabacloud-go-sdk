// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemberResponse
	GetStatusCode() *int32
	SetBody(v *ListMemberResponseBody) *ListMemberResponse
	GetBody() *ListMemberResponseBody
}

type ListMemberResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemberResponse) GoString() string {
	return s.String()
}

func (s *ListMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemberResponse) GetBody() *ListMemberResponseBody {
	return s.Body
}

func (s *ListMemberResponse) SetHeaders(v map[string]*string) *ListMemberResponse {
	s.Headers = v
	return s
}

func (s *ListMemberResponse) SetStatusCode(v int32) *ListMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemberResponse) SetBody(v *ListMemberResponseBody) *ListMemberResponse {
	s.Body = v
	return s
}

func (s *ListMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
