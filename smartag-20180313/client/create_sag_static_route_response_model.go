// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSagStaticRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSagStaticRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSagStaticRouteResponse
	GetStatusCode() *int32
	SetBody(v *CreateSagStaticRouteResponseBody) *CreateSagStaticRouteResponse
	GetBody() *CreateSagStaticRouteResponseBody
}

type CreateSagStaticRouteResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSagStaticRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSagStaticRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSagStaticRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateSagStaticRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSagStaticRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSagStaticRouteResponse) GetBody() *CreateSagStaticRouteResponseBody {
	return s.Body
}

func (s *CreateSagStaticRouteResponse) SetHeaders(v map[string]*string) *CreateSagStaticRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateSagStaticRouteResponse) SetStatusCode(v int32) *CreateSagStaticRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSagStaticRouteResponse) SetBody(v *CreateSagStaticRouteResponseBody) *CreateSagStaticRouteResponse {
	s.Body = v
	return s
}

func (s *CreateSagStaticRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
