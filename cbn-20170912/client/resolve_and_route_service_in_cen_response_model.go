// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveAndRouteServiceInCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResolveAndRouteServiceInCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResolveAndRouteServiceInCenResponse
	GetStatusCode() *int32
	SetBody(v *ResolveAndRouteServiceInCenResponseBody) *ResolveAndRouteServiceInCenResponse
	GetBody() *ResolveAndRouteServiceInCenResponseBody
}

type ResolveAndRouteServiceInCenResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResolveAndRouteServiceInCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResolveAndRouteServiceInCenResponse) String() string {
	return dara.Prettify(s)
}

func (s ResolveAndRouteServiceInCenResponse) GoString() string {
	return s.String()
}

func (s *ResolveAndRouteServiceInCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResolveAndRouteServiceInCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResolveAndRouteServiceInCenResponse) GetBody() *ResolveAndRouteServiceInCenResponseBody {
	return s.Body
}

func (s *ResolveAndRouteServiceInCenResponse) SetHeaders(v map[string]*string) *ResolveAndRouteServiceInCenResponse {
	s.Headers = v
	return s
}

func (s *ResolveAndRouteServiceInCenResponse) SetStatusCode(v int32) *ResolveAndRouteServiceInCenResponse {
	s.StatusCode = &v
	return s
}

func (s *ResolveAndRouteServiceInCenResponse) SetBody(v *ResolveAndRouteServiceInCenResponseBody) *ResolveAndRouteServiceInCenResponse {
	s.Body = v
	return s
}

func (s *ResolveAndRouteServiceInCenResponse) Validate() error {
	return dara.Validate(s)
}
