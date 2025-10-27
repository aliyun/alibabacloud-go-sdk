// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostpayUserNatStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePostpayUserNatStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePostpayUserNatStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePostpayUserNatStatusResponseBody) *UpdatePostpayUserNatStatusResponse
	GetBody() *UpdatePostpayUserNatStatusResponseBody
}

type UpdatePostpayUserNatStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePostpayUserNatStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePostpayUserNatStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostpayUserNatStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdatePostpayUserNatStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePostpayUserNatStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePostpayUserNatStatusResponse) GetBody() *UpdatePostpayUserNatStatusResponseBody {
	return s.Body
}

func (s *UpdatePostpayUserNatStatusResponse) SetHeaders(v map[string]*string) *UpdatePostpayUserNatStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdatePostpayUserNatStatusResponse) SetStatusCode(v int32) *UpdatePostpayUserNatStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePostpayUserNatStatusResponse) SetBody(v *UpdatePostpayUserNatStatusResponseBody) *UpdatePostpayUserNatStatusResponse {
	s.Body = v
	return s
}

func (s *UpdatePostpayUserNatStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
