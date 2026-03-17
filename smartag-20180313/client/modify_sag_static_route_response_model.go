// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagStaticRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagStaticRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagStaticRouteResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagStaticRouteResponseBody) *ModifySagStaticRouteResponse
	GetBody() *ModifySagStaticRouteResponseBody
}

type ModifySagStaticRouteResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagStaticRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagStaticRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagStaticRouteResponse) GoString() string {
	return s.String()
}

func (s *ModifySagStaticRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagStaticRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagStaticRouteResponse) GetBody() *ModifySagStaticRouteResponseBody {
	return s.Body
}

func (s *ModifySagStaticRouteResponse) SetHeaders(v map[string]*string) *ModifySagStaticRouteResponse {
	s.Headers = v
	return s
}

func (s *ModifySagStaticRouteResponse) SetStatusCode(v int32) *ModifySagStaticRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagStaticRouteResponse) SetBody(v *ModifySagStaticRouteResponseBody) *ModifySagStaticRouteResponse {
	s.Body = v
	return s
}

func (s *ModifySagStaticRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
