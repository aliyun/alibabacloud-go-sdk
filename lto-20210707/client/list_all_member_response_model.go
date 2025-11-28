// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllMemberResponse
	GetStatusCode() *int32
	SetBody(v *ListAllMemberResponseBody) *ListAllMemberResponse
	GetBody() *ListAllMemberResponseBody
}

type ListAllMemberResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllMemberResponse) GoString() string {
	return s.String()
}

func (s *ListAllMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllMemberResponse) GetBody() *ListAllMemberResponseBody {
	return s.Body
}

func (s *ListAllMemberResponse) SetHeaders(v map[string]*string) *ListAllMemberResponse {
	s.Headers = v
	return s
}

func (s *ListAllMemberResponse) SetStatusCode(v int32) *ListAllMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllMemberResponse) SetBody(v *ListAllMemberResponseBody) *ListAllMemberResponse {
	s.Body = v
	return s
}

func (s *ListAllMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
