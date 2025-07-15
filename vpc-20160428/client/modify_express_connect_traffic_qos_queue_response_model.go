// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectTrafficQosQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressConnectTrafficQosQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressConnectTrafficQosQueueResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressConnectTrafficQosQueueResponseBody) *ModifyExpressConnectTrafficQosQueueResponse
	GetBody() *ModifyExpressConnectTrafficQosQueueResponseBody
}

type ModifyExpressConnectTrafficQosQueueResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectTrafficQosQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectTrafficQosQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosQueueResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressConnectTrafficQosQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressConnectTrafficQosQueueResponse) GetBody() *ModifyExpressConnectTrafficQosQueueResponseBody {
	return s.Body
}

func (s *ModifyExpressConnectTrafficQosQueueResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectTrafficQosQueueResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueResponse) SetStatusCode(v int32) *ModifyExpressConnectTrafficQosQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueResponse) SetBody(v *ModifyExpressConnectTrafficQosQueueResponseBody) *ModifyExpressConnectTrafficQosQueueResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueResponse) Validate() error {
	return dara.Validate(s)
}
