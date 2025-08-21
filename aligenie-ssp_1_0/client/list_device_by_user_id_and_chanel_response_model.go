// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceByUserIdAndChanelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceByUserIdAndChanelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceByUserIdAndChanelResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceByUserIdAndChanelResponseBody) *ListDeviceByUserIdAndChanelResponse
	GetBody() *ListDeviceByUserIdAndChanelResponseBody
}

type ListDeviceByUserIdAndChanelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceByUserIdAndChanelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceByUserIdAndChanelResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceByUserIdAndChanelResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceByUserIdAndChanelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceByUserIdAndChanelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceByUserIdAndChanelResponse) GetBody() *ListDeviceByUserIdAndChanelResponseBody {
	return s.Body
}

func (s *ListDeviceByUserIdAndChanelResponse) SetHeaders(v map[string]*string) *ListDeviceByUserIdAndChanelResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponse) SetStatusCode(v int32) *ListDeviceByUserIdAndChanelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponse) SetBody(v *ListDeviceByUserIdAndChanelResponseBody) *ListDeviceByUserIdAndChanelResponse {
	s.Body = v
	return s
}

func (s *ListDeviceByUserIdAndChanelResponse) Validate() error {
	return dara.Validate(s)
}
