// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKillInstanceSessionTaskWithMaintainUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKillInstanceSessionTaskWithMaintainUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKillInstanceSessionTaskWithMaintainUserResponse
	GetStatusCode() *int32
	SetBody(v *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) *CreateKillInstanceSessionTaskWithMaintainUserResponse
	GetBody() *CreateKillInstanceSessionTaskWithMaintainUserResponseBody
}

type CreateKillInstanceSessionTaskWithMaintainUserResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKillInstanceSessionTaskWithMaintainUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKillInstanceSessionTaskWithMaintainUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKillInstanceSessionTaskWithMaintainUserResponse) GoString() string {
	return s.String()
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponse) GetBody() *CreateKillInstanceSessionTaskWithMaintainUserResponseBody {
	return s.Body
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponse) SetHeaders(v map[string]*string) *CreateKillInstanceSessionTaskWithMaintainUserResponse {
	s.Headers = v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponse) SetStatusCode(v int32) *CreateKillInstanceSessionTaskWithMaintainUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponse) SetBody(v *CreateKillInstanceSessionTaskWithMaintainUserResponseBody) *CreateKillInstanceSessionTaskWithMaintainUserResponse {
	s.Body = v
	return s
}

func (s *CreateKillInstanceSessionTaskWithMaintainUserResponse) Validate() error {
	return dara.Validate(s)
}
