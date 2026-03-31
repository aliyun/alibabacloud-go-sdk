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
	// The request ID.
	//
	// example:
	//
	// AB2E892E-8A43-5B0F-8FE3-B53ADA53CB2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	StopConfigurationRecorderResult *bool `json:"StopConfigurationRecorderResult,omitempty" xml:"StopConfigurationRecorderResult,omitempty"`
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
