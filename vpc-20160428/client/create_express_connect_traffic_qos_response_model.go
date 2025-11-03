// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectTrafficQosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExpressConnectTrafficQosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExpressConnectTrafficQosResponse
	GetStatusCode() *int32
	SetBody(v *CreateExpressConnectTrafficQosResponseBody) *CreateExpressConnectTrafficQosResponse
	GetBody() *CreateExpressConnectTrafficQosResponseBody
}

type CreateExpressConnectTrafficQosResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressConnectTrafficQosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressConnectTrafficQosResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExpressConnectTrafficQosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExpressConnectTrafficQosResponse) GetBody() *CreateExpressConnectTrafficQosResponseBody {
	return s.Body
}

func (s *CreateExpressConnectTrafficQosResponse) SetHeaders(v map[string]*string) *CreateExpressConnectTrafficQosResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressConnectTrafficQosResponse) SetStatusCode(v int32) *CreateExpressConnectTrafficQosResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressConnectTrafficQosResponse) SetBody(v *CreateExpressConnectTrafficQosResponseBody) *CreateExpressConnectTrafficQosResponse {
	s.Body = v
	return s
}

func (s *CreateExpressConnectTrafficQosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
