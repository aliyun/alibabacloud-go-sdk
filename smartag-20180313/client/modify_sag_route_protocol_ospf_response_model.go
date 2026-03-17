// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagRouteProtocolOspfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagRouteProtocolOspfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagRouteProtocolOspfResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagRouteProtocolOspfResponseBody) *ModifySagRouteProtocolOspfResponse
	GetBody() *ModifySagRouteProtocolOspfResponseBody
}

type ModifySagRouteProtocolOspfResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagRouteProtocolOspfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagRouteProtocolOspfResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagRouteProtocolOspfResponse) GoString() string {
	return s.String()
}

func (s *ModifySagRouteProtocolOspfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagRouteProtocolOspfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagRouteProtocolOspfResponse) GetBody() *ModifySagRouteProtocolOspfResponseBody {
	return s.Body
}

func (s *ModifySagRouteProtocolOspfResponse) SetHeaders(v map[string]*string) *ModifySagRouteProtocolOspfResponse {
	s.Headers = v
	return s
}

func (s *ModifySagRouteProtocolOspfResponse) SetStatusCode(v int32) *ModifySagRouteProtocolOspfResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagRouteProtocolOspfResponse) SetBody(v *ModifySagRouteProtocolOspfResponseBody) *ModifySagRouteProtocolOspfResponse {
	s.Body = v
	return s
}

func (s *ModifySagRouteProtocolOspfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
