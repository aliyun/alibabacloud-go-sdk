// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridMonitorSLSGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHybridMonitorSLSGroupResponseBody
	GetCode() *string
	SetMessage(v string) *CreateHybridMonitorSLSGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHybridMonitorSLSGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateHybridMonitorSLSGroupResponseBody
	GetSuccess() *string
}

type CreateHybridMonitorSLSGroupResponseBody struct {
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
	// Duplicate.SLSGroup
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

func (s CreateHybridMonitorSLSGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorSLSGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorSLSGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHybridMonitorSLSGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHybridMonitorSLSGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHybridMonitorSLSGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateHybridMonitorSLSGroupResponseBody) SetCode(v string) *CreateHybridMonitorSLSGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupResponseBody) SetMessage(v string) *CreateHybridMonitorSLSGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupResponseBody) SetRequestId(v string) *CreateHybridMonitorSLSGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupResponseBody) SetSuccess(v string) *CreateHybridMonitorSLSGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHybridMonitorSLSGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
