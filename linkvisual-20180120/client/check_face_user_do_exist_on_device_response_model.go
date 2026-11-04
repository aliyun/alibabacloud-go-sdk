// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFaceUserDoExistOnDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckFaceUserDoExistOnDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckFaceUserDoExistOnDeviceResponse
	GetStatusCode() *int32
	SetBody(v *CheckFaceUserDoExistOnDeviceResponseBody) *CheckFaceUserDoExistOnDeviceResponse
	GetBody() *CheckFaceUserDoExistOnDeviceResponseBody
}

type CheckFaceUserDoExistOnDeviceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckFaceUserDoExistOnDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckFaceUserDoExistOnDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckFaceUserDoExistOnDeviceResponse) GoString() string {
	return s.String()
}

func (s *CheckFaceUserDoExistOnDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckFaceUserDoExistOnDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckFaceUserDoExistOnDeviceResponse) GetBody() *CheckFaceUserDoExistOnDeviceResponseBody {
	return s.Body
}

func (s *CheckFaceUserDoExistOnDeviceResponse) SetHeaders(v map[string]*string) *CheckFaceUserDoExistOnDeviceResponse {
	s.Headers = v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponse) SetStatusCode(v int32) *CheckFaceUserDoExistOnDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponse) SetBody(v *CheckFaceUserDoExistOnDeviceResponseBody) *CheckFaceUserDoExistOnDeviceResponse {
	s.Body = v
	return s
}

func (s *CheckFaceUserDoExistOnDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
