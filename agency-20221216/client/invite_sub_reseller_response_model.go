// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteSubResellerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InviteSubResellerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InviteSubResellerResponse
	GetStatusCode() *int32
	SetBody(v *InviteSubResellerResponseBody) *InviteSubResellerResponse
	GetBody() *InviteSubResellerResponseBody
}

type InviteSubResellerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InviteSubResellerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InviteSubResellerResponse) String() string {
	return dara.Prettify(s)
}

func (s InviteSubResellerResponse) GoString() string {
	return s.String()
}

func (s *InviteSubResellerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InviteSubResellerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InviteSubResellerResponse) GetBody() *InviteSubResellerResponseBody {
	return s.Body
}

func (s *InviteSubResellerResponse) SetHeaders(v map[string]*string) *InviteSubResellerResponse {
	s.Headers = v
	return s
}

func (s *InviteSubResellerResponse) SetStatusCode(v int32) *InviteSubResellerResponse {
	s.StatusCode = &v
	return s
}

func (s *InviteSubResellerResponse) SetBody(v *InviteSubResellerResponseBody) *InviteSubResellerResponse {
	s.Body = v
	return s
}

func (s *InviteSubResellerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
