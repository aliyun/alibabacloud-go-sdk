// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInnerIpWhitelistGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInnerIpWhitelistGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInnerIpWhitelistGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInnerIpWhitelistGroupResponseBody) *UpdateInnerIpWhitelistGroupResponse
	GetBody() *UpdateInnerIpWhitelistGroupResponseBody
}

type UpdateInnerIpWhitelistGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInnerIpWhitelistGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInnerIpWhitelistGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInnerIpWhitelistGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateInnerIpWhitelistGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInnerIpWhitelistGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInnerIpWhitelistGroupResponse) GetBody() *UpdateInnerIpWhitelistGroupResponseBody {
	return s.Body
}

func (s *UpdateInnerIpWhitelistGroupResponse) SetHeaders(v map[string]*string) *UpdateInnerIpWhitelistGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateInnerIpWhitelistGroupResponse) SetStatusCode(v int32) *UpdateInnerIpWhitelistGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInnerIpWhitelistGroupResponse) SetBody(v *UpdateInnerIpWhitelistGroupResponseBody) *UpdateInnerIpWhitelistGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateInnerIpWhitelistGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
