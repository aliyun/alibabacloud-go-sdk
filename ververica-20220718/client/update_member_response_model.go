// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMemberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMemberResponseBody) *UpdateMemberResponse
	GetBody() *UpdateMemberResponseBody
}

type UpdateMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMemberResponse) GetBody() *UpdateMemberResponseBody {
	return s.Body
}

func (s *UpdateMemberResponse) SetHeaders(v map[string]*string) *UpdateMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemberResponse) SetStatusCode(v int32) *UpdateMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemberResponse) SetBody(v *UpdateMemberResponseBody) *UpdateMemberResponse {
	s.Body = v
	return s
}

func (s *UpdateMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
