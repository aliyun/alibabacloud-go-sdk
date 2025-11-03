// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectTrafficQosQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExpressConnectTrafficQosQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExpressConnectTrafficQosQueueResponse
	GetStatusCode() *int32
	SetBody(v *CreateExpressConnectTrafficQosQueueResponseBody) *CreateExpressConnectTrafficQosQueueResponse
	GetBody() *CreateExpressConnectTrafficQosQueueResponseBody
}

type CreateExpressConnectTrafficQosQueueResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressConnectTrafficQosQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressConnectTrafficQosQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExpressConnectTrafficQosQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExpressConnectTrafficQosQueueResponse) GetBody() *CreateExpressConnectTrafficQosQueueResponseBody {
	return s.Body
}

func (s *CreateExpressConnectTrafficQosQueueResponse) SetHeaders(v map[string]*string) *CreateExpressConnectTrafficQosQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueResponse) SetStatusCode(v int32) *CreateExpressConnectTrafficQosQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueResponse) SetBody(v *CreateExpressConnectTrafficQosQueueResponseBody) *CreateExpressConnectTrafficQosQueueResponse {
	s.Body = v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
