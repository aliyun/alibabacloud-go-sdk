// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAlbServerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachAlbServerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachAlbServerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DetachAlbServerGroupsResponseBody) *DetachAlbServerGroupsResponse
	GetBody() *DetachAlbServerGroupsResponseBody
}

type DetachAlbServerGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachAlbServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachAlbServerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachAlbServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DetachAlbServerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachAlbServerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachAlbServerGroupsResponse) GetBody() *DetachAlbServerGroupsResponseBody {
	return s.Body
}

func (s *DetachAlbServerGroupsResponse) SetHeaders(v map[string]*string) *DetachAlbServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DetachAlbServerGroupsResponse) SetStatusCode(v int32) *DetachAlbServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachAlbServerGroupsResponse) SetBody(v *DetachAlbServerGroupsResponseBody) *DetachAlbServerGroupsResponse {
	s.Body = v
	return s
}

func (s *DetachAlbServerGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
