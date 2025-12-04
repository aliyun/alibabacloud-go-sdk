// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMonitorAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameter(v string) *CheckMonitorAlertResponseBody
	GetParameter() *string
	SetRequestId(v string) *CheckMonitorAlertResponseBody
	GetRequestId() *string
	SetState(v string) *CheckMonitorAlertResponseBody
	GetState() *string
}

type CheckMonitorAlertResponseBody struct {
	// The parameters that are used to configure the monitoring and alerting feature.
	//
	// example:
	//
	// {   "monitor":{     "key1":"value1",     "key2":"value2"   },   "alert":{     "key1":"value1",     "key2":"value2"   } }
	Parameter *string `json:"Parameter,omitempty" xml:"Parameter,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 94F92113-FF63-5E57-8401-6FE123AD11DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the monitoring and alerting feature is enabled. Valid values:
	//
	// 	- **enable**: The monitoring and alerting feature is enabled.
	//
	// 	- **disable**: The monitoring and alerting feature is disabled.
	//
	// example:
	//
	// enable
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CheckMonitorAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckMonitorAlertResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMonitorAlertResponseBody) GetParameter() *string {
	return s.Parameter
}

func (s *CheckMonitorAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckMonitorAlertResponseBody) GetState() *string {
	return s.State
}

func (s *CheckMonitorAlertResponseBody) SetParameter(v string) *CheckMonitorAlertResponseBody {
	s.Parameter = &v
	return s
}

func (s *CheckMonitorAlertResponseBody) SetRequestId(v string) *CheckMonitorAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMonitorAlertResponseBody) SetState(v string) *CheckMonitorAlertResponseBody {
	s.State = &v
	return s
}

func (s *CheckMonitorAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
