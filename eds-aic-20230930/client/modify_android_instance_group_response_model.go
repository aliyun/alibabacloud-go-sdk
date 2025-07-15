// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAndroidInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAndroidInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAndroidInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAndroidInstanceGroupResponseBody) *ModifyAndroidInstanceGroupResponse
	GetBody() *ModifyAndroidInstanceGroupResponseBody
}

type ModifyAndroidInstanceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAndroidInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAndroidInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAndroidInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAndroidInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAndroidInstanceGroupResponse) GetBody() *ModifyAndroidInstanceGroupResponseBody {
	return s.Body
}

func (s *ModifyAndroidInstanceGroupResponse) SetHeaders(v map[string]*string) *ModifyAndroidInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAndroidInstanceGroupResponse) SetStatusCode(v int32) *ModifyAndroidInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAndroidInstanceGroupResponse) SetBody(v *ModifyAndroidInstanceGroupResponseBody) *ModifyAndroidInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyAndroidInstanceGroupResponse) Validate() error {
	return dara.Validate(s)
}
