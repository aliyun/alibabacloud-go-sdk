// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagRouteProtocolBgpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagRouteProtocolBgpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagRouteProtocolBgpResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagRouteProtocolBgpResponseBody) *ModifySagRouteProtocolBgpResponse
	GetBody() *ModifySagRouteProtocolBgpResponseBody
}

type ModifySagRouteProtocolBgpResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagRouteProtocolBgpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagRouteProtocolBgpResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagRouteProtocolBgpResponse) GoString() string {
	return s.String()
}

func (s *ModifySagRouteProtocolBgpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagRouteProtocolBgpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagRouteProtocolBgpResponse) GetBody() *ModifySagRouteProtocolBgpResponseBody {
	return s.Body
}

func (s *ModifySagRouteProtocolBgpResponse) SetHeaders(v map[string]*string) *ModifySagRouteProtocolBgpResponse {
	s.Headers = v
	return s
}

func (s *ModifySagRouteProtocolBgpResponse) SetStatusCode(v int32) *ModifySagRouteProtocolBgpResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagRouteProtocolBgpResponse) SetBody(v *ModifySagRouteProtocolBgpResponseBody) *ModifySagRouteProtocolBgpResponse {
	s.Body = v
	return s
}

func (s *ModifySagRouteProtocolBgpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
