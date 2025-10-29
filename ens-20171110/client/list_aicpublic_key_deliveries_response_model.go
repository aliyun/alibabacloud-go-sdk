// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICPublicKeyDeliveriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAICPublicKeyDeliveriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAICPublicKeyDeliveriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAICPublicKeyDeliveriesResponseBody) *ListAICPublicKeyDeliveriesResponse
	GetBody() *ListAICPublicKeyDeliveriesResponseBody
}

type ListAICPublicKeyDeliveriesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAICPublicKeyDeliveriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAICPublicKeyDeliveriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAICPublicKeyDeliveriesResponse) GoString() string {
	return s.String()
}

func (s *ListAICPublicKeyDeliveriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAICPublicKeyDeliveriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAICPublicKeyDeliveriesResponse) GetBody() *ListAICPublicKeyDeliveriesResponseBody {
	return s.Body
}

func (s *ListAICPublicKeyDeliveriesResponse) SetHeaders(v map[string]*string) *ListAICPublicKeyDeliveriesResponse {
	s.Headers = v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponse) SetStatusCode(v int32) *ListAICPublicKeyDeliveriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponse) SetBody(v *ListAICPublicKeyDeliveriesResponseBody) *ListAICPublicKeyDeliveriesResponse {
	s.Body = v
	return s
}

func (s *ListAICPublicKeyDeliveriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
