// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAckClusterConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAckClusterConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAckClusterConnectorResponse
	GetStatusCode() *int32
	SetBody(v *CreateAckClusterConnectorResponseBody) *CreateAckClusterConnectorResponse
	GetBody() *CreateAckClusterConnectorResponseBody
}

type CreateAckClusterConnectorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAckClusterConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAckClusterConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAckClusterConnectorResponse) GoString() string {
	return s.String()
}

func (s *CreateAckClusterConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAckClusterConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAckClusterConnectorResponse) GetBody() *CreateAckClusterConnectorResponseBody {
	return s.Body
}

func (s *CreateAckClusterConnectorResponse) SetHeaders(v map[string]*string) *CreateAckClusterConnectorResponse {
	s.Headers = v
	return s
}

func (s *CreateAckClusterConnectorResponse) SetStatusCode(v int32) *CreateAckClusterConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAckClusterConnectorResponse) SetBody(v *CreateAckClusterConnectorResponseBody) *CreateAckClusterConnectorResponse {
	s.Body = v
	return s
}

func (s *CreateAckClusterConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
