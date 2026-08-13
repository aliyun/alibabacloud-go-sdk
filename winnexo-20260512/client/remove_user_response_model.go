// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserResponseBody) *RemoveUserResponse
	GetBody() *RemoveUserResponseBody
}

type RemoveUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserResponse) GetBody() *RemoveUserResponseBody {
	return s.Body
}

func (s *RemoveUserResponse) SetHeaders(v map[string]*string) *RemoveUserResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserResponse) SetStatusCode(v int32) *RemoveUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserResponse) SetBody(v *RemoveUserResponseBody) *RemoveUserResponse {
	s.Body = v
	return s
}

func (s *RemoveUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
