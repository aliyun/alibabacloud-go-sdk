// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveInputSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMediaLiveInputSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMediaLiveInputSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateMediaLiveInputSecurityGroupResponseBody) *CreateMediaLiveInputSecurityGroupResponse
	GetBody() *CreateMediaLiveInputSecurityGroupResponseBody
}

type CreateMediaLiveInputSecurityGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMediaLiveInputSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMediaLiveInputSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveInputSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveInputSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMediaLiveInputSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMediaLiveInputSecurityGroupResponse) GetBody() *CreateMediaLiveInputSecurityGroupResponseBody {
	return s.Body
}

func (s *CreateMediaLiveInputSecurityGroupResponse) SetHeaders(v map[string]*string) *CreateMediaLiveInputSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMediaLiveInputSecurityGroupResponse) SetStatusCode(v int32) *CreateMediaLiveInputSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMediaLiveInputSecurityGroupResponse) SetBody(v *CreateMediaLiveInputSecurityGroupResponseBody) *CreateMediaLiveInputSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *CreateMediaLiveInputSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
