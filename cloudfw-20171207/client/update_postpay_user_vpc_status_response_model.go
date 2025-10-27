// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostpayUserVpcStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePostpayUserVpcStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePostpayUserVpcStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePostpayUserVpcStatusResponseBody) *UpdatePostpayUserVpcStatusResponse
	GetBody() *UpdatePostpayUserVpcStatusResponseBody
}

type UpdatePostpayUserVpcStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePostpayUserVpcStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePostpayUserVpcStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostpayUserVpcStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdatePostpayUserVpcStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePostpayUserVpcStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePostpayUserVpcStatusResponse) GetBody() *UpdatePostpayUserVpcStatusResponseBody {
	return s.Body
}

func (s *UpdatePostpayUserVpcStatusResponse) SetHeaders(v map[string]*string) *UpdatePostpayUserVpcStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdatePostpayUserVpcStatusResponse) SetStatusCode(v int32) *UpdatePostpayUserVpcStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePostpayUserVpcStatusResponse) SetBody(v *UpdatePostpayUserVpcStatusResponseBody) *UpdatePostpayUserVpcStatusResponse {
	s.Body = v
	return s
}

func (s *UpdatePostpayUserVpcStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
