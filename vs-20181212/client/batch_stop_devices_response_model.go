// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStopDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStopDevicesResponse
	GetStatusCode() *int32
	SetBody(v *BatchStopDevicesResponseBody) *BatchStopDevicesResponse
	GetBody() *BatchStopDevicesResponseBody
}

type BatchStopDevicesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStopDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStopDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStopDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchStopDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStopDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStopDevicesResponse) GetBody() *BatchStopDevicesResponseBody {
	return s.Body
}

func (s *BatchStopDevicesResponse) SetHeaders(v map[string]*string) *BatchStopDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchStopDevicesResponse) SetStatusCode(v int32) *BatchStopDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStopDevicesResponse) SetBody(v *BatchStopDevicesResponseBody) *BatchStopDevicesResponse {
	s.Body = v
	return s
}

func (s *BatchStopDevicesResponse) Validate() error {
	return dara.Validate(s)
}
