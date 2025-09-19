// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateApplicationResponse
	GetStatusCode() *int32
	SetBody(v *OperateApplicationResponseBody) *OperateApplicationResponse
	GetBody() *OperateApplicationResponseBody
}

type OperateApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateApplicationResponse) GoString() string {
	return s.String()
}

func (s *OperateApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateApplicationResponse) GetBody() *OperateApplicationResponseBody {
	return s.Body
}

func (s *OperateApplicationResponse) SetHeaders(v map[string]*string) *OperateApplicationResponse {
	s.Headers = v
	return s
}

func (s *OperateApplicationResponse) SetStatusCode(v int32) *OperateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateApplicationResponse) SetBody(v *OperateApplicationResponseBody) *OperateApplicationResponse {
	s.Body = v
	return s
}

func (s *OperateApplicationResponse) Validate() error {
	return dara.Validate(s)
}
