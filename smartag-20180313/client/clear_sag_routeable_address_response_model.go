// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearSagRouteableAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearSagRouteableAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearSagRouteableAddressResponse
	GetStatusCode() *int32
	SetBody(v *ClearSagRouteableAddressResponseBody) *ClearSagRouteableAddressResponse
	GetBody() *ClearSagRouteableAddressResponseBody
}

type ClearSagRouteableAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearSagRouteableAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearSagRouteableAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearSagRouteableAddressResponse) GoString() string {
	return s.String()
}

func (s *ClearSagRouteableAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearSagRouteableAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearSagRouteableAddressResponse) GetBody() *ClearSagRouteableAddressResponseBody {
	return s.Body
}

func (s *ClearSagRouteableAddressResponse) SetHeaders(v map[string]*string) *ClearSagRouteableAddressResponse {
	s.Headers = v
	return s
}

func (s *ClearSagRouteableAddressResponse) SetStatusCode(v int32) *ClearSagRouteableAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearSagRouteableAddressResponse) SetBody(v *ClearSagRouteableAddressResponseBody) *ClearSagRouteableAddressResponse {
	s.Body = v
	return s
}

func (s *ClearSagRouteableAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
