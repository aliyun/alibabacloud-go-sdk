// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryVisionDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchQueryVisionDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchQueryVisionDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *BatchQueryVisionDeviceInfoResponseBody) *BatchQueryVisionDeviceInfoResponse
	GetBody() *BatchQueryVisionDeviceInfoResponseBody
}

type BatchQueryVisionDeviceInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryVisionDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryVisionDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryVisionDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryVisionDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchQueryVisionDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchQueryVisionDeviceInfoResponse) GetBody() *BatchQueryVisionDeviceInfoResponseBody {
	return s.Body
}

func (s *BatchQueryVisionDeviceInfoResponse) SetHeaders(v map[string]*string) *BatchQueryVisionDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponse) SetStatusCode(v int32) *BatchQueryVisionDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponse) SetBody(v *BatchQueryVisionDeviceInfoResponseBody) *BatchQueryVisionDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *BatchQueryVisionDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
