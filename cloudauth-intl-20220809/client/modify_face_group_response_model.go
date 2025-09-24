// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFaceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFaceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFaceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFaceGroupResponseBody) *ModifyFaceGroupResponse
	GetBody() *ModifyFaceGroupResponseBody
}

type ModifyFaceGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFaceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFaceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFaceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyFaceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFaceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFaceGroupResponse) GetBody() *ModifyFaceGroupResponseBody {
	return s.Body
}

func (s *ModifyFaceGroupResponse) SetHeaders(v map[string]*string) *ModifyFaceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyFaceGroupResponse) SetStatusCode(v int32) *ModifyFaceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFaceGroupResponse) SetBody(v *ModifyFaceGroupResponseBody) *ModifyFaceGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyFaceGroupResponse) Validate() error {
	return dara.Validate(s)
}
