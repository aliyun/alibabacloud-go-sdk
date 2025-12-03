// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppMemberResponse
	GetStatusCode() *int32
	SetBody(v string) *UpdateAppMemberResponse
	GetBody() *string
}

type UpdateAppMemberResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppMemberResponse) GetBody() *string {
	return s.Body
}

func (s *UpdateAppMemberResponse) SetHeaders(v map[string]*string) *UpdateAppMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppMemberResponse) SetStatusCode(v int32) *UpdateAppMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppMemberResponse) SetBody(v string) *UpdateAppMemberResponse {
	s.Body = &v
	return s
}

func (s *UpdateAppMemberResponse) Validate() error {
	return dara.Validate(s)
}
