// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanMachineEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirusScanMachineEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirusScanMachineEventResponse
	GetStatusCode() *int32
	SetBody(v *ListVirusScanMachineEventResponseBody) *ListVirusScanMachineEventResponse
	GetBody() *ListVirusScanMachineEventResponseBody
}

type ListVirusScanMachineEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirusScanMachineEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirusScanMachineEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponse) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirusScanMachineEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirusScanMachineEventResponse) GetBody() *ListVirusScanMachineEventResponseBody {
	return s.Body
}

func (s *ListVirusScanMachineEventResponse) SetHeaders(v map[string]*string) *ListVirusScanMachineEventResponse {
	s.Headers = v
	return s
}

func (s *ListVirusScanMachineEventResponse) SetStatusCode(v int32) *ListVirusScanMachineEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirusScanMachineEventResponse) SetBody(v *ListVirusScanMachineEventResponseBody) *ListVirusScanMachineEventResponse {
	s.Body = v
	return s
}

func (s *ListVirusScanMachineEventResponse) Validate() error {
	return dara.Validate(s)
}
