// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDynamicTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDynamicTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDynamicTagGroupResponseBody) *DeleteDynamicTagGroupResponse
	GetBody() *DeleteDynamicTagGroupResponseBody
}

type DeleteDynamicTagGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDynamicTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDynamicTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicTagGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDynamicTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDynamicTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDynamicTagGroupResponse) GetBody() *DeleteDynamicTagGroupResponseBody {
	return s.Body
}

func (s *DeleteDynamicTagGroupResponse) SetHeaders(v map[string]*string) *DeleteDynamicTagGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDynamicTagGroupResponse) SetStatusCode(v int32) *DeleteDynamicTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDynamicTagGroupResponse) SetBody(v *DeleteDynamicTagGroupResponseBody) *DeleteDynamicTagGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteDynamicTagGroupResponse) Validate() error {
	return dara.Validate(s)
}
