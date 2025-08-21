// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserDeviceListByTmeUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserDeviceListByTmeUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserDeviceListByTmeUserIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserDeviceListByTmeUserIdResponseBody) *QueryUserDeviceListByTmeUserIdResponse
	GetBody() *QueryUserDeviceListByTmeUserIdResponseBody
}

type QueryUserDeviceListByTmeUserIdResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserDeviceListByTmeUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserDeviceListByTmeUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserDeviceListByTmeUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryUserDeviceListByTmeUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserDeviceListByTmeUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserDeviceListByTmeUserIdResponse) GetBody() *QueryUserDeviceListByTmeUserIdResponseBody {
	return s.Body
}

func (s *QueryUserDeviceListByTmeUserIdResponse) SetHeaders(v map[string]*string) *QueryUserDeviceListByTmeUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponse) SetStatusCode(v int32) *QueryUserDeviceListByTmeUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponse) SetBody(v *QueryUserDeviceListByTmeUserIdResponseBody) *QueryUserDeviceListByTmeUserIdResponse {
	s.Body = v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdResponse) Validate() error {
	return dara.Validate(s)
}
