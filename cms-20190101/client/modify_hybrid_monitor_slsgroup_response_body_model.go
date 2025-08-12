// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridMonitorSLSGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyHybridMonitorSLSGroupResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyHybridMonitorSLSGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyHybridMonitorSLSGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyHybridMonitorSLSGroupResponseBody
	GetSuccess() *string
}

type ModifyHybridMonitorSLSGroupResponseBody struct {
	// The HTTP status code.
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
	// NotFound.SLSGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 66683237-7126-50F8-BBF8-D67ACC919A17
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

func (s ModifyHybridMonitorSLSGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorSLSGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorSLSGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyHybridMonitorSLSGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyHybridMonitorSLSGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridMonitorSLSGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyHybridMonitorSLSGroupResponseBody) SetCode(v string) *ModifyHybridMonitorSLSGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupResponseBody) SetMessage(v string) *ModifyHybridMonitorSLSGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupResponseBody) SetRequestId(v string) *ModifyHybridMonitorSLSGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupResponseBody) SetSuccess(v string) *ModifyHybridMonitorSLSGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
