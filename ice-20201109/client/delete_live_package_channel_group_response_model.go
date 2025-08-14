// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageChannelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLivePackageChannelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLivePackageChannelGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLivePackageChannelGroupResponseBody) *DeleteLivePackageChannelGroupResponse
	GetBody() *DeleteLivePackageChannelGroupResponseBody
}

type DeleteLivePackageChannelGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivePackageChannelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivePackageChannelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageChannelGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageChannelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLivePackageChannelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLivePackageChannelGroupResponse) GetBody() *DeleteLivePackageChannelGroupResponseBody {
	return s.Body
}

func (s *DeleteLivePackageChannelGroupResponse) SetHeaders(v map[string]*string) *DeleteLivePackageChannelGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivePackageChannelGroupResponse) SetStatusCode(v int32) *DeleteLivePackageChannelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivePackageChannelGroupResponse) SetBody(v *DeleteLivePackageChannelGroupResponseBody) *DeleteLivePackageChannelGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteLivePackageChannelGroupResponse) Validate() error {
	return dara.Validate(s)
}
