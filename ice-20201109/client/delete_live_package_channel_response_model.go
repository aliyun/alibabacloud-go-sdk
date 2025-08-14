// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLivePackageChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLivePackageChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLivePackageChannelResponseBody) *DeleteLivePackageChannelResponse
	GetBody() *DeleteLivePackageChannelResponseBody
}

type DeleteLivePackageChannelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivePackageChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivePackageChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLivePackageChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLivePackageChannelResponse) GetBody() *DeleteLivePackageChannelResponseBody {
	return s.Body
}

func (s *DeleteLivePackageChannelResponse) SetHeaders(v map[string]*string) *DeleteLivePackageChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivePackageChannelResponse) SetStatusCode(v int32) *DeleteLivePackageChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivePackageChannelResponse) SetBody(v *DeleteLivePackageChannelResponseBody) *DeleteLivePackageChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteLivePackageChannelResponse) Validate() error {
	return dara.Validate(s)
}
