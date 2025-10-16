// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostpayUserInternetStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePostpayUserInternetStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePostpayUserInternetStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePostpayUserInternetStatusResponseBody) *UpdatePostpayUserInternetStatusResponse
	GetBody() *UpdatePostpayUserInternetStatusResponseBody
}

type UpdatePostpayUserInternetStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePostpayUserInternetStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePostpayUserInternetStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostpayUserInternetStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdatePostpayUserInternetStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePostpayUserInternetStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePostpayUserInternetStatusResponse) GetBody() *UpdatePostpayUserInternetStatusResponseBody {
	return s.Body
}

func (s *UpdatePostpayUserInternetStatusResponse) SetHeaders(v map[string]*string) *UpdatePostpayUserInternetStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdatePostpayUserInternetStatusResponse) SetStatusCode(v int32) *UpdatePostpayUserInternetStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePostpayUserInternetStatusResponse) SetBody(v *UpdatePostpayUserInternetStatusResponseBody) *UpdatePostpayUserInternetStatusResponse {
	s.Body = v
	return s
}

func (s *UpdatePostpayUserInternetStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
