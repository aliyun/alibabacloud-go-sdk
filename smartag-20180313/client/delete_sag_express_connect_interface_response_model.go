// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSagExpressConnectInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSagExpressConnectInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSagExpressConnectInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSagExpressConnectInterfaceResponseBody) *DeleteSagExpressConnectInterfaceResponse
	GetBody() *DeleteSagExpressConnectInterfaceResponseBody
}

type DeleteSagExpressConnectInterfaceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSagExpressConnectInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSagExpressConnectInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSagExpressConnectInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSagExpressConnectInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSagExpressConnectInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSagExpressConnectInterfaceResponse) GetBody() *DeleteSagExpressConnectInterfaceResponseBody {
	return s.Body
}

func (s *DeleteSagExpressConnectInterfaceResponse) SetHeaders(v map[string]*string) *DeleteSagExpressConnectInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSagExpressConnectInterfaceResponse) SetStatusCode(v int32) *DeleteSagExpressConnectInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceResponse) SetBody(v *DeleteSagExpressConnectInterfaceResponseBody) *DeleteSagExpressConnectInterfaceResponse {
	s.Body = v
	return s
}

func (s *DeleteSagExpressConnectInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
