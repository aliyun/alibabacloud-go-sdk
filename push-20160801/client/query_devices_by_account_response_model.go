// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicesByAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDevicesByAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDevicesByAccountResponse
	GetStatusCode() *int32
	SetBody(v *QueryDevicesByAccountResponseBody) *QueryDevicesByAccountResponse
	GetBody() *QueryDevicesByAccountResponseBody
}

type QueryDevicesByAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDevicesByAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDevicesByAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicesByAccountResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDevicesByAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDevicesByAccountResponse) GetBody() *QueryDevicesByAccountResponseBody {
	return s.Body
}

func (s *QueryDevicesByAccountResponse) SetHeaders(v map[string]*string) *QueryDevicesByAccountResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicesByAccountResponse) SetStatusCode(v int32) *QueryDevicesByAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicesByAccountResponse) SetBody(v *QueryDevicesByAccountResponseBody) *QueryDevicesByAccountResponse {
	s.Body = v
	return s
}

func (s *QueryDevicesByAccountResponse) Validate() error {
	return dara.Validate(s)
}
