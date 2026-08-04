// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindCustomerInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindCustomerInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindCustomerInfoResponse
	GetStatusCode() *int32
	SetBody(v *FindCustomerInfoResponseBody) *FindCustomerInfoResponse
	GetBody() *FindCustomerInfoResponseBody
}

type FindCustomerInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindCustomerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindCustomerInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerInfoResponse) GoString() string {
	return s.String()
}

func (s *FindCustomerInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindCustomerInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindCustomerInfoResponse) GetBody() *FindCustomerInfoResponseBody {
	return s.Body
}

func (s *FindCustomerInfoResponse) SetHeaders(v map[string]*string) *FindCustomerInfoResponse {
	s.Headers = v
	return s
}

func (s *FindCustomerInfoResponse) SetStatusCode(v int32) *FindCustomerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *FindCustomerInfoResponse) SetBody(v *FindCustomerInfoResponseBody) *FindCustomerInfoResponse {
	s.Body = v
	return s
}

func (s *FindCustomerInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
