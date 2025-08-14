// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLiveInputSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaLiveInputSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaLiveInputSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaLiveInputSecurityGroupResponseBody) *DeleteMediaLiveInputSecurityGroupResponse
	GetBody() *DeleteMediaLiveInputSecurityGroupResponseBody
}

type DeleteMediaLiveInputSecurityGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaLiveInputSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaLiveInputSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLiveInputSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaLiveInputSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaLiveInputSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaLiveInputSecurityGroupResponse) GetBody() *DeleteMediaLiveInputSecurityGroupResponseBody {
	return s.Body
}

func (s *DeleteMediaLiveInputSecurityGroupResponse) SetHeaders(v map[string]*string) *DeleteMediaLiveInputSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaLiveInputSecurityGroupResponse) SetStatusCode(v int32) *DeleteMediaLiveInputSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaLiveInputSecurityGroupResponse) SetBody(v *DeleteMediaLiveInputSecurityGroupResponseBody) *DeleteMediaLiveInputSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaLiveInputSecurityGroupResponse) Validate() error {
	return dara.Validate(s)
}
