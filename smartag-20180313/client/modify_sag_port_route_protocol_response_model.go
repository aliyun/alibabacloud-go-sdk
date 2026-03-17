// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagPortRouteProtocolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagPortRouteProtocolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagPortRouteProtocolResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagPortRouteProtocolResponseBody) *ModifySagPortRouteProtocolResponse
	GetBody() *ModifySagPortRouteProtocolResponseBody
}

type ModifySagPortRouteProtocolResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagPortRouteProtocolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagPortRouteProtocolResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagPortRouteProtocolResponse) GoString() string {
	return s.String()
}

func (s *ModifySagPortRouteProtocolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagPortRouteProtocolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagPortRouteProtocolResponse) GetBody() *ModifySagPortRouteProtocolResponseBody {
	return s.Body
}

func (s *ModifySagPortRouteProtocolResponse) SetHeaders(v map[string]*string) *ModifySagPortRouteProtocolResponse {
	s.Headers = v
	return s
}

func (s *ModifySagPortRouteProtocolResponse) SetStatusCode(v int32) *ModifySagPortRouteProtocolResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagPortRouteProtocolResponse) SetBody(v *ModifySagPortRouteProtocolResponseBody) *ModifySagPortRouteProtocolResponse {
	s.Body = v
	return s
}

func (s *ModifySagPortRouteProtocolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
