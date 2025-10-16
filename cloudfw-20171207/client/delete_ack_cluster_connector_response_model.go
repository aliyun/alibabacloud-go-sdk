// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAckClusterConnectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAckClusterConnectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAckClusterConnectorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAckClusterConnectorResponseBody) *DeleteAckClusterConnectorResponse
	GetBody() *DeleteAckClusterConnectorResponseBody
}

type DeleteAckClusterConnectorResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAckClusterConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAckClusterConnectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAckClusterConnectorResponse) GoString() string {
	return s.String()
}

func (s *DeleteAckClusterConnectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAckClusterConnectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAckClusterConnectorResponse) GetBody() *DeleteAckClusterConnectorResponseBody {
	return s.Body
}

func (s *DeleteAckClusterConnectorResponse) SetHeaders(v map[string]*string) *DeleteAckClusterConnectorResponse {
	s.Headers = v
	return s
}

func (s *DeleteAckClusterConnectorResponse) SetStatusCode(v int32) *DeleteAckClusterConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAckClusterConnectorResponse) SetBody(v *DeleteAckClusterConnectorResponseBody) *DeleteAckClusterConnectorResponse {
	s.Body = v
	return s
}

func (s *DeleteAckClusterConnectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
