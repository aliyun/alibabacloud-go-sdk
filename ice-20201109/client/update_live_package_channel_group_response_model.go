// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageChannelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLivePackageChannelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLivePackageChannelGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLivePackageChannelGroupResponseBody) *UpdateLivePackageChannelGroupResponse
	GetBody() *UpdateLivePackageChannelGroupResponseBody
}

type UpdateLivePackageChannelGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLivePackageChannelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLivePackageChannelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLivePackageChannelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLivePackageChannelGroupResponse) GetBody() *UpdateLivePackageChannelGroupResponseBody {
	return s.Body
}

func (s *UpdateLivePackageChannelGroupResponse) SetHeaders(v map[string]*string) *UpdateLivePackageChannelGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivePackageChannelGroupResponse) SetStatusCode(v int32) *UpdateLivePackageChannelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivePackageChannelGroupResponse) SetBody(v *UpdateLivePackageChannelGroupResponseBody) *UpdateLivePackageChannelGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateLivePackageChannelGroupResponse) Validate() error {
	return dara.Validate(s)
}
