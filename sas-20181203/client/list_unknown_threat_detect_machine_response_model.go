// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUnknownThreatDetectMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUnknownThreatDetectMachineResponse
	GetStatusCode() *int32
	SetBody(v *ListUnknownThreatDetectMachineResponseBody) *ListUnknownThreatDetectMachineResponse
	GetBody() *ListUnknownThreatDetectMachineResponseBody
}

type ListUnknownThreatDetectMachineResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUnknownThreatDetectMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUnknownThreatDetectMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectMachineResponse) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUnknownThreatDetectMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUnknownThreatDetectMachineResponse) GetBody() *ListUnknownThreatDetectMachineResponseBody {
	return s.Body
}

func (s *ListUnknownThreatDetectMachineResponse) SetHeaders(v map[string]*string) *ListUnknownThreatDetectMachineResponse {
	s.Headers = v
	return s
}

func (s *ListUnknownThreatDetectMachineResponse) SetStatusCode(v int32) *ListUnknownThreatDetectMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponse) SetBody(v *ListUnknownThreatDetectMachineResponseBody) *ListUnknownThreatDetectMachineResponse {
	s.Body = v
	return s
}

func (s *ListUnknownThreatDetectMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
