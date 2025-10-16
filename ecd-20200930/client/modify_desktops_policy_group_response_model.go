// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopsPolicyGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopsPolicyGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopsPolicyGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopsPolicyGroupResponseBody) *ModifyDesktopsPolicyGroupResponse
	GetBody() *ModifyDesktopsPolicyGroupResponseBody
}

type ModifyDesktopsPolicyGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopsPolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopsPolicyGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopsPolicyGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopsPolicyGroupResponse) GetBody() *ModifyDesktopsPolicyGroupResponseBody {
	return s.Body
}

func (s *ModifyDesktopsPolicyGroupResponse) SetHeaders(v map[string]*string) *ModifyDesktopsPolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponse) SetStatusCode(v int32) *ModifyDesktopsPolicyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponse) SetBody(v *ModifyDesktopsPolicyGroupResponseBody) *ModifyDesktopsPolicyGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
