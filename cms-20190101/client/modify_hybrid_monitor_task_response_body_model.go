// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridMonitorTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyHybridMonitorTaskResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyHybridMonitorTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyHybridMonitorTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyHybridMonitorTaskResponseBody
	GetSuccess() *string
}

type ModifyHybridMonitorTaskResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid.SLSConfig
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 11145B76-566A-5D80-A8A3-FAD98D310079
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyHybridMonitorTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyHybridMonitorTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyHybridMonitorTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridMonitorTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyHybridMonitorTaskResponseBody) SetCode(v string) *ModifyHybridMonitorTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyHybridMonitorTaskResponseBody) SetMessage(v string) *ModifyHybridMonitorTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyHybridMonitorTaskResponseBody) SetRequestId(v string) *ModifyHybridMonitorTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridMonitorTaskResponseBody) SetSuccess(v string) *ModifyHybridMonitorTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyHybridMonitorTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
