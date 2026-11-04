// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVisionDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVisionDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVisionDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryVisionDeviceInfoResponseBody) *QueryVisionDeviceInfoResponse
	GetBody() *QueryVisionDeviceInfoResponseBody
}

type QueryVisionDeviceInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVisionDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVisionDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVisionDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryVisionDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVisionDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVisionDeviceInfoResponse) GetBody() *QueryVisionDeviceInfoResponseBody {
	return s.Body
}

func (s *QueryVisionDeviceInfoResponse) SetHeaders(v map[string]*string) *QueryVisionDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryVisionDeviceInfoResponse) SetStatusCode(v int32) *QueryVisionDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVisionDeviceInfoResponse) SetBody(v *QueryVisionDeviceInfoResponseBody) *QueryVisionDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *QueryVisionDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
