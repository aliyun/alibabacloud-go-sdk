// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLhMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLhMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLhMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddLhMembersResponseBody) *AddLhMembersResponse
	GetBody() *AddLhMembersResponseBody
}

type AddLhMembersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLhMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLhMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLhMembersResponse) GoString() string {
	return s.String()
}

func (s *AddLhMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLhMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLhMembersResponse) GetBody() *AddLhMembersResponseBody {
	return s.Body
}

func (s *AddLhMembersResponse) SetHeaders(v map[string]*string) *AddLhMembersResponse {
	s.Headers = v
	return s
}

func (s *AddLhMembersResponse) SetStatusCode(v int32) *AddLhMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLhMembersResponse) SetBody(v *AddLhMembersResponseBody) *AddLhMembersResponse {
	s.Body = v
	return s
}

func (s *AddLhMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
