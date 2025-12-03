// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppMembersResponse
	GetStatusCode() *int32
	SetBody(v string) *CreateAppMembersResponse
	GetBody() *string
}

type CreateAppMembersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppMembersResponse) GoString() string {
	return s.String()
}

func (s *CreateAppMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppMembersResponse) GetBody() *string {
	return s.Body
}

func (s *CreateAppMembersResponse) SetHeaders(v map[string]*string) *CreateAppMembersResponse {
	s.Headers = v
	return s
}

func (s *CreateAppMembersResponse) SetStatusCode(v int32) *CreateAppMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppMembersResponse) SetBody(v string) *CreateAppMembersResponse {
	s.Body = &v
	return s
}

func (s *CreateAppMembersResponse) Validate() error {
	return dara.Validate(s)
}
