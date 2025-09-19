// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmMachineCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlarmMachineCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlarmMachineCountResponse
	GetStatusCode() *int32
	SetBody(v *GetAlarmMachineCountResponseBody) *GetAlarmMachineCountResponse
	GetBody() *GetAlarmMachineCountResponseBody
}

type GetAlarmMachineCountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlarmMachineCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlarmMachineCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmMachineCountResponse) GoString() string {
	return s.String()
}

func (s *GetAlarmMachineCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlarmMachineCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlarmMachineCountResponse) GetBody() *GetAlarmMachineCountResponseBody {
	return s.Body
}

func (s *GetAlarmMachineCountResponse) SetHeaders(v map[string]*string) *GetAlarmMachineCountResponse {
	s.Headers = v
	return s
}

func (s *GetAlarmMachineCountResponse) SetStatusCode(v int32) *GetAlarmMachineCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlarmMachineCountResponse) SetBody(v *GetAlarmMachineCountResponseBody) *GetAlarmMachineCountResponse {
	s.Body = v
	return s
}

func (s *GetAlarmMachineCountResponse) Validate() error {
	return dara.Validate(s)
}
