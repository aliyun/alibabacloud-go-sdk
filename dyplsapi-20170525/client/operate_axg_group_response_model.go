// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAxgGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateAxgGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateAxgGroupResponse
	GetStatusCode() *int32
	SetBody(v *OperateAxgGroupResponseBody) *OperateAxgGroupResponse
	GetBody() *OperateAxgGroupResponseBody
}

type OperateAxgGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAxgGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAxgGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateAxgGroupResponse) GoString() string {
	return s.String()
}

func (s *OperateAxgGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateAxgGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateAxgGroupResponse) GetBody() *OperateAxgGroupResponseBody {
	return s.Body
}

func (s *OperateAxgGroupResponse) SetHeaders(v map[string]*string) *OperateAxgGroupResponse {
	s.Headers = v
	return s
}

func (s *OperateAxgGroupResponse) SetStatusCode(v int32) *OperateAxgGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAxgGroupResponse) SetBody(v *OperateAxgGroupResponseBody) *OperateAxgGroupResponse {
	s.Body = v
	return s
}

func (s *OperateAxgGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
