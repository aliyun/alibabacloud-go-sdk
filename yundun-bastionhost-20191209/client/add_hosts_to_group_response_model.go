// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHostsToGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddHostsToGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddHostsToGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddHostsToGroupResponseBody) *AddHostsToGroupResponse
	GetBody() *AddHostsToGroupResponseBody
}

type AddHostsToGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddHostsToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddHostsToGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddHostsToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddHostsToGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddHostsToGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddHostsToGroupResponse) GetBody() *AddHostsToGroupResponseBody {
	return s.Body
}

func (s *AddHostsToGroupResponse) SetHeaders(v map[string]*string) *AddHostsToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddHostsToGroupResponse) SetStatusCode(v int32) *AddHostsToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHostsToGroupResponse) SetBody(v *AddHostsToGroupResponseBody) *AddHostsToGroupResponse {
	s.Body = v
	return s
}

func (s *AddHostsToGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
