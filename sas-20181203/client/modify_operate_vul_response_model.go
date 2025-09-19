// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOperateVulResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOperateVulResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOperateVulResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOperateVulResponseBody) *ModifyOperateVulResponse
	GetBody() *ModifyOperateVulResponseBody
}

type ModifyOperateVulResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOperateVulResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOperateVulResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOperateVulResponse) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOperateVulResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOperateVulResponse) GetBody() *ModifyOperateVulResponseBody {
	return s.Body
}

func (s *ModifyOperateVulResponse) SetHeaders(v map[string]*string) *ModifyOperateVulResponse {
	s.Headers = v
	return s
}

func (s *ModifyOperateVulResponse) SetStatusCode(v int32) *ModifyOperateVulResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOperateVulResponse) SetBody(v *ModifyOperateVulResponseBody) *ModifyOperateVulResponse {
	s.Body = v
	return s
}

func (s *ModifyOperateVulResponse) Validate() error {
	return dara.Validate(s)
}
