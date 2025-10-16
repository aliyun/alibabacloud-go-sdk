// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAckClusterConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAckClusterConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAckClusterConnectorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAckClusterConnectorResponseBody) *UpdateAckClusterConnectorResponse
	GetBody() *UpdateAckClusterConnectorResponseBody
}

type UpdateAckClusterConnectorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAckClusterConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAckClusterConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAckClusterConnectorResponse) GoString() string {
	return s.String()
}

func (s *UpdateAckClusterConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAckClusterConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAckClusterConnectorResponse) GetBody() *UpdateAckClusterConnectorResponseBody {
	return s.Body
}

func (s *UpdateAckClusterConnectorResponse) SetHeaders(v map[string]*string) *UpdateAckClusterConnectorResponse {
	s.Headers = v
	return s
}

func (s *UpdateAckClusterConnectorResponse) SetStatusCode(v int32) *UpdateAckClusterConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAckClusterConnectorResponse) SetBody(v *UpdateAckClusterConnectorResponseBody) *UpdateAckClusterConnectorResponse {
	s.Body = v
	return s
}

func (s *UpdateAckClusterConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
