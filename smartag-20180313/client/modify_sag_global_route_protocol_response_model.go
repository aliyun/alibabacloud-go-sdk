// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagGlobalRouteProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagGlobalRouteProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagGlobalRouteProtocolResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagGlobalRouteProtocolResponseBody) *ModifySagGlobalRouteProtocolResponse
	GetBody() *ModifySagGlobalRouteProtocolResponseBody
}

type ModifySagGlobalRouteProtocolResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagGlobalRouteProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagGlobalRouteProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagGlobalRouteProtocolResponse) GoString() string {
	return s.String()
}

func (s *ModifySagGlobalRouteProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagGlobalRouteProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagGlobalRouteProtocolResponse) GetBody() *ModifySagGlobalRouteProtocolResponseBody {
	return s.Body
}

func (s *ModifySagGlobalRouteProtocolResponse) SetHeaders(v map[string]*string) *ModifySagGlobalRouteProtocolResponse {
	s.Headers = v
	return s
}

func (s *ModifySagGlobalRouteProtocolResponse) SetStatusCode(v int32) *ModifySagGlobalRouteProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolResponse) SetBody(v *ModifySagGlobalRouteProtocolResponseBody) *ModifySagGlobalRouteProtocolResponse {
	s.Body = v
	return s
}

func (s *ModifySagGlobalRouteProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
