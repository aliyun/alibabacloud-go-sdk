// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceMemberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceMemberResponseBody) *UpdateResourceMemberResponse
	GetBody() *UpdateResourceMemberResponseBody
}

type UpdateResourceMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceMemberResponse) GetBody() *UpdateResourceMemberResponseBody {
	return s.Body
}

func (s *UpdateResourceMemberResponse) SetHeaders(v map[string]*string) *UpdateResourceMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceMemberResponse) SetStatusCode(v int32) *UpdateResourceMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceMemberResponse) SetBody(v *UpdateResourceMemberResponseBody) *UpdateResourceMemberResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
