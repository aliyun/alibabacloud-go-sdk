// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachHostAccountsFromUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachHostAccountsFromUserResponse
	GetStatusCode() *int32
	SetBody(v *DetachHostAccountsFromUserResponseBody) *DetachHostAccountsFromUserResponse
	GetBody() *DetachHostAccountsFromUserResponseBody
}

type DetachHostAccountsFromUserResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostAccountsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachHostAccountsFromUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserResponse) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachHostAccountsFromUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachHostAccountsFromUserResponse) GetBody() *DetachHostAccountsFromUserResponseBody {
	return s.Body
}

func (s *DetachHostAccountsFromUserResponse) SetHeaders(v map[string]*string) *DetachHostAccountsFromUserResponse {
	s.Headers = v
	return s
}

func (s *DetachHostAccountsFromUserResponse) SetStatusCode(v int32) *DetachHostAccountsFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostAccountsFromUserResponse) SetBody(v *DetachHostAccountsFromUserResponseBody) *DetachHostAccountsFromUserResponse {
	s.Body = v
	return s
}

func (s *DetachHostAccountsFromUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
