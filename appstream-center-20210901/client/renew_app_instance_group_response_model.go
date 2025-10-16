// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewAppInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewAppInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *RenewAppInstanceGroupResponseBody) *RenewAppInstanceGroupResponse
	GetBody() *RenewAppInstanceGroupResponseBody
}

type RenewAppInstanceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAppInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewAppInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewAppInstanceGroupResponse) GetBody() *RenewAppInstanceGroupResponseBody {
	return s.Body
}

func (s *RenewAppInstanceGroupResponse) SetHeaders(v map[string]*string) *RenewAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *RenewAppInstanceGroupResponse) SetStatusCode(v int32) *RenewAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAppInstanceGroupResponse) SetBody(v *RenewAppInstanceGroupResponseBody) *RenewAppInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *RenewAppInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
