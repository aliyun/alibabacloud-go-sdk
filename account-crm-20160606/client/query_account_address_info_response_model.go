// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountAddressInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountAddressInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountAddressInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountAddressInfoResponseBody) *QueryAccountAddressInfoResponse
	GetBody() *QueryAccountAddressInfoResponseBody
}

type QueryAccountAddressInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountAddressInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountAddressInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountAddressInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountAddressInfoResponse) GetBody() *QueryAccountAddressInfoResponseBody {
	return s.Body
}

func (s *QueryAccountAddressInfoResponse) SetHeaders(v map[string]*string) *QueryAccountAddressInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountAddressInfoResponse) SetStatusCode(v int32) *QueryAccountAddressInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountAddressInfoResponse) SetBody(v *QueryAccountAddressInfoResponseBody) *QueryAccountAddressInfoResponse {
	s.Body = v
	return s
}

func (s *QueryAccountAddressInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
