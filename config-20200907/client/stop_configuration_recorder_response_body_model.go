// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopConfigurationRecorderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopConfigurationRecorderResponseBody
	GetRequestId() *string
	SetStopConfigurationRecorderResult(v bool) *StopConfigurationRecorderResponseBody
	GetStopConfigurationRecorderResult() *bool
}

type StopConfigurationRecorderResponseBody struct {
	RequestId                       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StopConfigurationRecorderResult *bool   `json:"StopConfigurationRecorderResult,omitempty" xml:"StopConfigurationRecorderResult,omitempty"`
}

func (s StopConfigurationRecorderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopConfigurationRecorderResponseBody) GoString() string {
	return s.String()
}

func (s *StopConfigurationRecorderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopConfigurationRecorderResponseBody) GetStopConfigurationRecorderResult() *bool {
	return s.StopConfigurationRecorderResult
}

func (s *StopConfigurationRecorderResponseBody) SetRequestId(v string) *StopConfigurationRecorderResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopConfigurationRecorderResponseBody) SetStopConfigurationRecorderResult(v bool) *StopConfigurationRecorderResponseBody {
	s.StopConfigurationRecorderResult = &v
	return s
}

func (s *StopConfigurationRecorderResponseBody) Validate() error {
	return dara.Validate(s)
}
