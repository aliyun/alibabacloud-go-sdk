// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirusScanMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirusScanMachineResponse
	GetStatusCode() *int32
	SetBody(v *ListVirusScanMachineResponseBody) *ListVirusScanMachineResponse
	GetBody() *ListVirusScanMachineResponseBody
}

type ListVirusScanMachineResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirusScanMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirusScanMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineResponse) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirusScanMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirusScanMachineResponse) GetBody() *ListVirusScanMachineResponseBody {
	return s.Body
}

func (s *ListVirusScanMachineResponse) SetHeaders(v map[string]*string) *ListVirusScanMachineResponse {
	s.Headers = v
	return s
}

func (s *ListVirusScanMachineResponse) SetStatusCode(v int32) *ListVirusScanMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirusScanMachineResponse) SetBody(v *ListVirusScanMachineResponseBody) *ListVirusScanMachineResponse {
	s.Body = v
	return s
}

func (s *ListVirusScanMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
