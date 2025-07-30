// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserBatchJoinGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UserBatchJoinGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UserBatchJoinGroupResponse
	GetStatusCode() *int32
	SetBody(v *UserBatchJoinGroupResponseBody) *UserBatchJoinGroupResponse
	GetBody() *UserBatchJoinGroupResponseBody
}

type UserBatchJoinGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UserBatchJoinGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UserBatchJoinGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UserBatchJoinGroupResponse) GoString() string {
	return s.String()
}

func (s *UserBatchJoinGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UserBatchJoinGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UserBatchJoinGroupResponse) GetBody() *UserBatchJoinGroupResponseBody {
	return s.Body
}

func (s *UserBatchJoinGroupResponse) SetHeaders(v map[string]*string) *UserBatchJoinGroupResponse {
	s.Headers = v
	return s
}

func (s *UserBatchJoinGroupResponse) SetStatusCode(v int32) *UserBatchJoinGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UserBatchJoinGroupResponse) SetBody(v *UserBatchJoinGroupResponseBody) *UserBatchJoinGroupResponse {
	s.Body = v
	return s
}

func (s *UserBatchJoinGroupResponse) Validate() error {
	return dara.Validate(s)
}
