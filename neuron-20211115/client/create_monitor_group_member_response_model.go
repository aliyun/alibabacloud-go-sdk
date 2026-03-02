// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorGroupMemberResponse
	GetStatusCode() *int32
}

type CreateMonitorGroupMemberResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateMonitorGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorGroupMemberResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupMemberResponse) SetStatusCode(v int32) *CreateMonitorGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorGroupMemberResponse) Validate() error {
	return dara.Validate(s)
}
