// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceCaptureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDeviceCaptureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDeviceCaptureResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDeviceCaptureResponseBody) *ModifyDeviceCaptureResponse
	GetBody() *ModifyDeviceCaptureResponseBody
}

type ModifyDeviceCaptureResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDeviceCaptureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDeviceCaptureResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceCaptureResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceCaptureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDeviceCaptureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDeviceCaptureResponse) GetBody() *ModifyDeviceCaptureResponseBody {
	return s.Body
}

func (s *ModifyDeviceCaptureResponse) SetHeaders(v map[string]*string) *ModifyDeviceCaptureResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceCaptureResponse) SetStatusCode(v int32) *ModifyDeviceCaptureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeviceCaptureResponse) SetBody(v *ModifyDeviceCaptureResponseBody) *ModifyDeviceCaptureResponse {
	s.Body = v
	return s
}

func (s *ModifyDeviceCaptureResponse) Validate() error {
	return dara.Validate(s)
}
