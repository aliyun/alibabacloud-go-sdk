// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestDataSourceConnectivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectivity(v *TestDataSourceConnectivityResponseBodyConnectivity) *TestDataSourceConnectivityResponseBody
	GetConnectivity() *TestDataSourceConnectivityResponseBodyConnectivity
	SetRequestId(v string) *TestDataSourceConnectivityResponseBody
	GetRequestId() *string
}

type TestDataSourceConnectivityResponseBody struct {
	// The details of the connectivity test.
	Connectivity *TestDataSourceConnectivityResponseBodyConnectivity `json:"Connectivity,omitempty" xml:"Connectivity,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA81****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TestDataSourceConnectivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestDataSourceConnectivityResponseBody) GoString() string {
	return s.String()
}

func (s *TestDataSourceConnectivityResponseBody) GetConnectivity() *TestDataSourceConnectivityResponseBodyConnectivity {
	return s.Connectivity
}

func (s *TestDataSourceConnectivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestDataSourceConnectivityResponseBody) SetConnectivity(v *TestDataSourceConnectivityResponseBodyConnectivity) *TestDataSourceConnectivityResponseBody {
	s.Connectivity = v
	return s
}

func (s *TestDataSourceConnectivityResponseBody) SetRequestId(v string) *TestDataSourceConnectivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestDataSourceConnectivityResponseBody) Validate() error {
	return dara.Validate(s)
}

type TestDataSourceConnectivityResponseBodyConnectivity struct {
	// The error message returned if the connectivity test fails. No such a message is returned if the connectivity test is successful.
	ConnectMessage *string `json:"ConnectMessage,omitempty" xml:"ConnectMessage,omitempty"`
	// The result of the connectivity test. Valid values: Connectable: The network can be connected. ConfigError: The network can be connected, but the configurations are incorrect. Unreachable: The network cannot be connected. Unsupport: An error is reported due to other causes. For example, the desired resource group is being initialized.
	//
	// example:
	//
	// Connectable
	ConnectState *string `json:"ConnectState,omitempty" xml:"ConnectState,omitempty"`
	// The detailed logs of each step in the connectivity test.
	DetailLogs []*TestDataSourceConnectivityResponseBodyConnectivityDetailLogs `json:"DetailLogs,omitempty" xml:"DetailLogs,omitempty" type:"Repeated"`
}

func (s TestDataSourceConnectivityResponseBodyConnectivity) String() string {
	return dara.Prettify(s)
}

func (s TestDataSourceConnectivityResponseBodyConnectivity) GoString() string {
	return s.String()
}

func (s *TestDataSourceConnectivityResponseBodyConnectivity) GetConnectMessage() *string {
	return s.ConnectMessage
}

func (s *TestDataSourceConnectivityResponseBodyConnectivity) GetConnectState() *string {
	return s.ConnectState
}

func (s *TestDataSourceConnectivityResponseBodyConnectivity) GetDetailLogs() []*TestDataSourceConnectivityResponseBodyConnectivityDetailLogs {
	return s.DetailLogs
}

func (s *TestDataSourceConnectivityResponseBodyConnectivity) SetConnectMessage(v string) *TestDataSourceConnectivityResponseBodyConnectivity {
	s.ConnectMessage = &v
	return s
}

func (s *TestDataSourceConnectivityResponseBodyConnectivity) SetConnectState(v string) *TestDataSourceConnectivityResponseBodyConnectivity {
	s.ConnectState = &v
	return s
}

func (s *TestDataSourceConnectivityResponseBodyConnectivity) SetDetailLogs(v []*TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) *TestDataSourceConnectivityResponseBodyConnectivity {
	s.DetailLogs = v
	return s
}

func (s *TestDataSourceConnectivityResponseBodyConnectivity) Validate() error {
	return dara.Validate(s)
}

type TestDataSourceConnectivityResponseBodyConnectivityDetailLogs struct {
	// The code of the test item.
	//
	// example:
	//
	// validate_input_parameters
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The end time of a step.
	//
	// example:
	//
	// 1730217604002
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the step.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The start time of a step.
	//
	// example:
	//
	// 1730217600001
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) String() string {
	return dara.Prettify(s)
}

func (s TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) GoString() string {
	return s.String()
}

func (s *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) GetCode() *string {
	return s.Code
}

func (s *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) GetEndTime() *int64 {
	return s.EndTime
}

func (s *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) GetMessage() *string {
	return s.Message
}

func (s *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) SetCode(v string) *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs {
	s.Code = &v
	return s
}

func (s *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) SetEndTime(v int64) *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs {
	s.EndTime = &v
	return s
}

func (s *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) SetMessage(v string) *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs {
	s.Message = &v
	return s
}

func (s *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) SetStartTime(v int64) *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs {
	s.StartTime = &v
	return s
}

func (s *TestDataSourceConnectivityResponseBodyConnectivityDetailLogs) Validate() error {
	return dara.Validate(s)
}
