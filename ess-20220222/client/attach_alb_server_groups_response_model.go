// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAlbServerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachAlbServerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachAlbServerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *AttachAlbServerGroupsResponseBody) *AttachAlbServerGroupsResponse
	GetBody() *AttachAlbServerGroupsResponseBody
}

type AttachAlbServerGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachAlbServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachAlbServerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachAlbServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *AttachAlbServerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachAlbServerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachAlbServerGroupsResponse) GetBody() *AttachAlbServerGroupsResponseBody {
	return s.Body
}

func (s *AttachAlbServerGroupsResponse) SetHeaders(v map[string]*string) *AttachAlbServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *AttachAlbServerGroupsResponse) SetStatusCode(v int32) *AttachAlbServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachAlbServerGroupsResponse) SetBody(v *AttachAlbServerGroupsResponseBody) *AttachAlbServerGroupsResponse {
	s.Body = v
	return s
}

func (s *AttachAlbServerGroupsResponse) Validate() error {
	return dara.Validate(s)
}
