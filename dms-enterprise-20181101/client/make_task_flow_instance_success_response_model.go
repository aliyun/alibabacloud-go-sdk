// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeTaskFlowInstanceSuccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MakeTaskFlowInstanceSuccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MakeTaskFlowInstanceSuccessResponse
	GetStatusCode() *int32
	SetBody(v *MakeTaskFlowInstanceSuccessResponseBody) *MakeTaskFlowInstanceSuccessResponse
	GetBody() *MakeTaskFlowInstanceSuccessResponseBody
}

type MakeTaskFlowInstanceSuccessResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MakeTaskFlowInstanceSuccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MakeTaskFlowInstanceSuccessResponse) String() string {
	return dara.Prettify(s)
}

func (s MakeTaskFlowInstanceSuccessResponse) GoString() string {
	return s.String()
}

func (s *MakeTaskFlowInstanceSuccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MakeTaskFlowInstanceSuccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MakeTaskFlowInstanceSuccessResponse) GetBody() *MakeTaskFlowInstanceSuccessResponseBody {
	return s.Body
}

func (s *MakeTaskFlowInstanceSuccessResponse) SetHeaders(v map[string]*string) *MakeTaskFlowInstanceSuccessResponse {
	s.Headers = v
	return s
}

func (s *MakeTaskFlowInstanceSuccessResponse) SetStatusCode(v int32) *MakeTaskFlowInstanceSuccessResponse {
	s.StatusCode = &v
	return s
}

func (s *MakeTaskFlowInstanceSuccessResponse) SetBody(v *MakeTaskFlowInstanceSuccessResponseBody) *MakeTaskFlowInstanceSuccessResponse {
	s.Body = v
	return s
}

func (s *MakeTaskFlowInstanceSuccessResponse) Validate() error {
	return dara.Validate(s)
}
