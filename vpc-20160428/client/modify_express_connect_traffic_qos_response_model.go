// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectTrafficQosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressConnectTrafficQosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressConnectTrafficQosResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressConnectTrafficQosResponseBody) *ModifyExpressConnectTrafficQosResponse
	GetBody() *ModifyExpressConnectTrafficQosResponseBody
}

type ModifyExpressConnectTrafficQosResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectTrafficQosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectTrafficQosResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressConnectTrafficQosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressConnectTrafficQosResponse) GetBody() *ModifyExpressConnectTrafficQosResponseBody {
	return s.Body
}

func (s *ModifyExpressConnectTrafficQosResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectTrafficQosResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectTrafficQosResponse) SetStatusCode(v int32) *ModifyExpressConnectTrafficQosResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosResponse) SetBody(v *ModifyExpressConnectTrafficQosResponseBody) *ModifyExpressConnectTrafficQosResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressConnectTrafficQosResponse) Validate() error {
	return dara.Validate(s)
}
