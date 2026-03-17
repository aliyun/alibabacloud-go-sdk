// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagExpressConnectInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagExpressConnectInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagExpressConnectInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagExpressConnectInterfaceResponseBody) *ModifySagExpressConnectInterfaceResponse
	GetBody() *ModifySagExpressConnectInterfaceResponseBody
}

type ModifySagExpressConnectInterfaceResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagExpressConnectInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagExpressConnectInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagExpressConnectInterfaceResponse) GoString() string {
	return s.String()
}

func (s *ModifySagExpressConnectInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagExpressConnectInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagExpressConnectInterfaceResponse) GetBody() *ModifySagExpressConnectInterfaceResponseBody {
	return s.Body
}

func (s *ModifySagExpressConnectInterfaceResponse) SetHeaders(v map[string]*string) *ModifySagExpressConnectInterfaceResponse {
	s.Headers = v
	return s
}

func (s *ModifySagExpressConnectInterfaceResponse) SetStatusCode(v int32) *ModifySagExpressConnectInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceResponse) SetBody(v *ModifySagExpressConnectInterfaceResponseBody) *ModifySagExpressConnectInterfaceResponse {
	s.Body = v
	return s
}

func (s *ModifySagExpressConnectInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
