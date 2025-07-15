// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectTrafficQosQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExpressConnectTrafficQosQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExpressConnectTrafficQosQueueResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExpressConnectTrafficQosQueueResponseBody) *DeleteExpressConnectTrafficQosQueueResponse
	GetBody() *DeleteExpressConnectTrafficQosQueueResponseBody
}

type DeleteExpressConnectTrafficQosQueueResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectTrafficQosQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectTrafficQosQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectTrafficQosQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectTrafficQosQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExpressConnectTrafficQosQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExpressConnectTrafficQosQueueResponse) GetBody() *DeleteExpressConnectTrafficQosQueueResponseBody {
	return s.Body
}

func (s *DeleteExpressConnectTrafficQosQueueResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectTrafficQosQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueResponse) SetStatusCode(v int32) *DeleteExpressConnectTrafficQosQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueResponse) SetBody(v *DeleteExpressConnectTrafficQosQueueResponseBody) *DeleteExpressConnectTrafficQosQueueResponse {
	s.Body = v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueResponse) Validate() error {
	return dara.Validate(s)
}
