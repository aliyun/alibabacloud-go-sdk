// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateInstanceResponse
	GetStatusCode() *int32
	SetBody(v *OperateInstanceResponseBody) *OperateInstanceResponse
	GetBody() *OperateInstanceResponseBody
}

type OperateInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateInstanceResponse) GoString() string {
	return s.String()
}

func (s *OperateInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateInstanceResponse) GetBody() *OperateInstanceResponseBody {
	return s.Body
}

func (s *OperateInstanceResponse) SetHeaders(v map[string]*string) *OperateInstanceResponse {
	s.Headers = v
	return s
}

func (s *OperateInstanceResponse) SetStatusCode(v int32) *OperateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateInstanceResponse) SetBody(v *OperateInstanceResponseBody) *OperateInstanceResponse {
	s.Body = v
	return s
}

func (s *OperateInstanceResponse) Validate() error {
	return dara.Validate(s)
}
