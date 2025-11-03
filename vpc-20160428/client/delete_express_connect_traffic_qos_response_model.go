// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectTrafficQosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExpressConnectTrafficQosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExpressConnectTrafficQosResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExpressConnectTrafficQosResponseBody) *DeleteExpressConnectTrafficQosResponse
	GetBody() *DeleteExpressConnectTrafficQosResponseBody
}

type DeleteExpressConnectTrafficQosResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectTrafficQosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectTrafficQosResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectTrafficQosResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectTrafficQosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExpressConnectTrafficQosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExpressConnectTrafficQosResponse) GetBody() *DeleteExpressConnectTrafficQosResponseBody {
	return s.Body
}

func (s *DeleteExpressConnectTrafficQosResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectTrafficQosResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectTrafficQosResponse) SetStatusCode(v int32) *DeleteExpressConnectTrafficQosResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosResponse) SetBody(v *DeleteExpressConnectTrafficQosResponseBody) *DeleteExpressConnectTrafficQosResponse {
	s.Body = v
	return s
}

func (s *DeleteExpressConnectTrafficQosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
