// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAndroidInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAndroidInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAndroidInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAndroidInstanceGroupResponseBody) *DeleteAndroidInstanceGroupResponse
	GetBody() *DeleteAndroidInstanceGroupResponseBody
}

type DeleteAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAndroidInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAndroidInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAndroidInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAndroidInstanceGroupResponse) GetBody() *DeleteAndroidInstanceGroupResponseBody {
	return s.Body
}

func (s *DeleteAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *DeleteAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAndroidInstanceGroupResponse) SetStatusCode(v int32) *DeleteAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAndroidInstanceGroupResponse) SetBody(v *DeleteAndroidInstanceGroupResponseBody) *DeleteAndroidInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteAndroidInstanceGroupResponse) Validate() error {
	return dara.Validate(s)
}
