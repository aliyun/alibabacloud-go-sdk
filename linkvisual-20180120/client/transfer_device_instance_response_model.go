// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferDeviceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferDeviceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferDeviceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *TransferDeviceInstanceResponseBody) *TransferDeviceInstanceResponse
	GetBody() *TransferDeviceInstanceResponseBody
}

type TransferDeviceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferDeviceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferDeviceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferDeviceInstanceResponse) GoString() string {
	return s.String()
}

func (s *TransferDeviceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferDeviceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferDeviceInstanceResponse) GetBody() *TransferDeviceInstanceResponseBody {
	return s.Body
}

func (s *TransferDeviceInstanceResponse) SetHeaders(v map[string]*string) *TransferDeviceInstanceResponse {
	s.Headers = v
	return s
}

func (s *TransferDeviceInstanceResponse) SetStatusCode(v int32) *TransferDeviceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferDeviceInstanceResponse) SetBody(v *TransferDeviceInstanceResponseBody) *TransferDeviceInstanceResponse {
	s.Body = v
	return s
}

func (s *TransferDeviceInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
