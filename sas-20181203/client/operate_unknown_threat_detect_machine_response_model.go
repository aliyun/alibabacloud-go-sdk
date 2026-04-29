// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnknownThreatDetectMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateUnknownThreatDetectMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateUnknownThreatDetectMachineResponse
	GetStatusCode() *int32
	SetBody(v *OperateUnknownThreatDetectMachineResponseBody) *OperateUnknownThreatDetectMachineResponse
	GetBody() *OperateUnknownThreatDetectMachineResponseBody
}

type OperateUnknownThreatDetectMachineResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateUnknownThreatDetectMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateUnknownThreatDetectMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateUnknownThreatDetectMachineResponse) GoString() string {
	return s.String()
}

func (s *OperateUnknownThreatDetectMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateUnknownThreatDetectMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateUnknownThreatDetectMachineResponse) GetBody() *OperateUnknownThreatDetectMachineResponseBody {
	return s.Body
}

func (s *OperateUnknownThreatDetectMachineResponse) SetHeaders(v map[string]*string) *OperateUnknownThreatDetectMachineResponse {
	s.Headers = v
	return s
}

func (s *OperateUnknownThreatDetectMachineResponse) SetStatusCode(v int32) *OperateUnknownThreatDetectMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateUnknownThreatDetectMachineResponse) SetBody(v *OperateUnknownThreatDetectMachineResponseBody) *OperateUnknownThreatDetectMachineResponse {
	s.Body = v
	return s
}

func (s *OperateUnknownThreatDetectMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
