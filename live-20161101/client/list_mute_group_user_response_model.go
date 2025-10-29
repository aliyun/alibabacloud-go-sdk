// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMuteGroupUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMuteGroupUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMuteGroupUserResponse
	GetStatusCode() *int32
	SetBody(v *ListMuteGroupUserResponseBody) *ListMuteGroupUserResponse
	GetBody() *ListMuteGroupUserResponseBody
}

type ListMuteGroupUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMuteGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMuteGroupUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMuteGroupUserResponse) GoString() string {
	return s.String()
}

func (s *ListMuteGroupUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMuteGroupUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMuteGroupUserResponse) GetBody() *ListMuteGroupUserResponseBody {
	return s.Body
}

func (s *ListMuteGroupUserResponse) SetHeaders(v map[string]*string) *ListMuteGroupUserResponse {
	s.Headers = v
	return s
}

func (s *ListMuteGroupUserResponse) SetStatusCode(v int32) *ListMuteGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMuteGroupUserResponse) SetBody(v *ListMuteGroupUserResponseBody) *ListMuteGroupUserResponse {
	s.Body = v
	return s
}

func (s *ListMuteGroupUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
