// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserGroupListByParentIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserGroupListByParentIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserGroupListByParentIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserGroupListByParentIdResponseBody) *QueryUserGroupListByParentIdResponse
	GetBody() *QueryUserGroupListByParentIdResponseBody
}

type QueryUserGroupListByParentIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserGroupListByParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserGroupListByParentIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserGroupListByParentIdResponse) GoString() string {
	return s.String()
}

func (s *QueryUserGroupListByParentIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserGroupListByParentIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserGroupListByParentIdResponse) GetBody() *QueryUserGroupListByParentIdResponseBody {
	return s.Body
}

func (s *QueryUserGroupListByParentIdResponse) SetHeaders(v map[string]*string) *QueryUserGroupListByParentIdResponse {
	s.Headers = v
	return s
}

func (s *QueryUserGroupListByParentIdResponse) SetStatusCode(v int32) *QueryUserGroupListByParentIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserGroupListByParentIdResponse) SetBody(v *QueryUserGroupListByParentIdResponseBody) *QueryUserGroupListByParentIdResponse {
	s.Body = v
	return s
}

func (s *QueryUserGroupListByParentIdResponse) Validate() error {
	return dara.Validate(s)
}
