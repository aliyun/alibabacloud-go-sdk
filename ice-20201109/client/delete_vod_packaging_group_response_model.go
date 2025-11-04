// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodPackagingGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVodPackagingGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVodPackagingGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVodPackagingGroupResponseBody) *DeleteVodPackagingGroupResponse
	GetBody() *DeleteVodPackagingGroupResponseBody
}

type DeleteVodPackagingGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVodPackagingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVodPackagingGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodPackagingGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodPackagingGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVodPackagingGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVodPackagingGroupResponse) GetBody() *DeleteVodPackagingGroupResponseBody {
	return s.Body
}

func (s *DeleteVodPackagingGroupResponse) SetHeaders(v map[string]*string) *DeleteVodPackagingGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodPackagingGroupResponse) SetStatusCode(v int32) *DeleteVodPackagingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVodPackagingGroupResponse) SetBody(v *DeleteVodPackagingGroupResponseBody) *DeleteVodPackagingGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteVodPackagingGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
