// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSagExpressConnectInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSagExpressConnectInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSagExpressConnectInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateSagExpressConnectInterfaceResponseBody) *CreateSagExpressConnectInterfaceResponse
	GetBody() *CreateSagExpressConnectInterfaceResponseBody
}

type CreateSagExpressConnectInterfaceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSagExpressConnectInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSagExpressConnectInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSagExpressConnectInterfaceResponse) GoString() string {
	return s.String()
}

func (s *CreateSagExpressConnectInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSagExpressConnectInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSagExpressConnectInterfaceResponse) GetBody() *CreateSagExpressConnectInterfaceResponseBody {
	return s.Body
}

func (s *CreateSagExpressConnectInterfaceResponse) SetHeaders(v map[string]*string) *CreateSagExpressConnectInterfaceResponse {
	s.Headers = v
	return s
}

func (s *CreateSagExpressConnectInterfaceResponse) SetStatusCode(v int32) *CreateSagExpressConnectInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceResponse) SetBody(v *CreateSagExpressConnectInterfaceResponseBody) *CreateSagExpressConnectInterfaceResponse {
	s.Body = v
	return s
}

func (s *CreateSagExpressConnectInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
