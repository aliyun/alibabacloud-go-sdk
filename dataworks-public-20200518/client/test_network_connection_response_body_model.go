// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestNetworkConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TestNetworkConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TestNetworkConnectionResponseBody
	GetSuccess() *bool
	SetTaskList(v *TestNetworkConnectionResponseBodyTaskList) *TestNetworkConnectionResponseBody
	GetTaskList() *TestNetworkConnectionResponseBodyTaskList
}

type TestNetworkConnectionResponseBody struct {
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about the connectivity test.
	TaskList *TestNetworkConnectionResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Struct"`
}

func (s TestNetworkConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestNetworkConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *TestNetworkConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestNetworkConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestNetworkConnectionResponseBody) GetTaskList() *TestNetworkConnectionResponseBodyTaskList {
	return s.TaskList
}

func (s *TestNetworkConnectionResponseBody) SetRequestId(v string) *TestNetworkConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestNetworkConnectionResponseBody) SetSuccess(v bool) *TestNetworkConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *TestNetworkConnectionResponseBody) SetTaskList(v *TestNetworkConnectionResponseBodyTaskList) *TestNetworkConnectionResponseBody {
	s.TaskList = v
	return s
}

func (s *TestNetworkConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}

type TestNetworkConnectionResponseBodyTaskList struct {
	// The reason why the data source and resource group failed the connectivity test. If data source and the resource group passed the connectivity test, this parameter is left empty.
	//
	// example:
	//
	// Connectable
	ConnectMessage *string `json:"ConnectMessage,omitempty" xml:"ConnectMessage,omitempty"`
	// The result of the connectivity test. Valid values:
	//
	// 	- true: The data source and the resource group passed the connectivity test.
	//
	// 	- false: The data source and the resource group failed the connectivity test. You can troubleshoot issues based on the ConnectMessage parameter.
	//
	// example:
	//
	// true
	ConnectStatus *bool `json:"ConnectStatus,omitempty" xml:"ConnectStatus,omitempty"`
}

func (s TestNetworkConnectionResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s TestNetworkConnectionResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *TestNetworkConnectionResponseBodyTaskList) GetConnectMessage() *string {
	return s.ConnectMessage
}

func (s *TestNetworkConnectionResponseBodyTaskList) GetConnectStatus() *bool {
	return s.ConnectStatus
}

func (s *TestNetworkConnectionResponseBodyTaskList) SetConnectMessage(v string) *TestNetworkConnectionResponseBodyTaskList {
	s.ConnectMessage = &v
	return s
}

func (s *TestNetworkConnectionResponseBodyTaskList) SetConnectStatus(v bool) *TestNetworkConnectionResponseBodyTaskList {
	s.ConnectStatus = &v
	return s
}

func (s *TestNetworkConnectionResponseBodyTaskList) Validate() error {
	return dara.Validate(s)
}
