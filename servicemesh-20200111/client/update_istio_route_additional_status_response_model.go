// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioRouteAdditionalStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIstioRouteAdditionalStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIstioRouteAdditionalStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIstioRouteAdditionalStatusResponseBody) *UpdateIstioRouteAdditionalStatusResponse
	GetBody() *UpdateIstioRouteAdditionalStatusResponseBody
}

type UpdateIstioRouteAdditionalStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIstioRouteAdditionalStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIstioRouteAdditionalStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioRouteAdditionalStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateIstioRouteAdditionalStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIstioRouteAdditionalStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIstioRouteAdditionalStatusResponse) GetBody() *UpdateIstioRouteAdditionalStatusResponseBody {
	return s.Body
}

func (s *UpdateIstioRouteAdditionalStatusResponse) SetHeaders(v map[string]*string) *UpdateIstioRouteAdditionalStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusResponse) SetStatusCode(v int32) *UpdateIstioRouteAdditionalStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusResponse) SetBody(v *UpdateIstioRouteAdditionalStatusResponseBody) *UpdateIstioRouteAdditionalStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateIstioRouteAdditionalStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
