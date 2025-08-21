// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteDevicesResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteDevicesResponseBody) *BatchDeleteDevicesResponse
	GetBody() *BatchDeleteDevicesResponseBody
}

type BatchDeleteDevicesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteDevicesResponse) GetBody() *BatchDeleteDevicesResponseBody {
	return s.Body
}

func (s *BatchDeleteDevicesResponse) SetHeaders(v map[string]*string) *BatchDeleteDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteDevicesResponse) SetStatusCode(v int32) *BatchDeleteDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteDevicesResponse) SetBody(v *BatchDeleteDevicesResponseBody) *BatchDeleteDevicesResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteDevicesResponse) Validate() error {
	return dara.Validate(s)
}
