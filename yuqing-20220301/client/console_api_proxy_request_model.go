// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleApiProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ConsoleBody) *ConsoleApiProxyRequest
	GetBody() *ConsoleBody
}

type ConsoleApiProxyRequest struct {
	Body *ConsoleBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsoleApiProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s ConsoleApiProxyRequest) GoString() string {
	return s.String()
}

func (s *ConsoleApiProxyRequest) GetBody() *ConsoleBody {
	return s.Body
}

func (s *ConsoleApiProxyRequest) SetBody(v *ConsoleBody) *ConsoleApiProxyRequest {
	s.Body = v
	return s
}

func (s *ConsoleApiProxyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
