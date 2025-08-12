// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridMonitorSLSGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHybridMonitorSLSGroupResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteHybridMonitorSLSGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHybridMonitorSLSGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteHybridMonitorSLSGroupResponseBody
	GetSuccess() *string
}

type DeleteHybridMonitorSLSGroupResponseBody struct {
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
	// RESOURCE_NOT_FOUND
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 85198BFF-1DE6-556E-B6A4-C77C51C62B8D
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

func (s DeleteHybridMonitorSLSGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridMonitorSLSGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHybridMonitorSLSGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHybridMonitorSLSGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHybridMonitorSLSGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHybridMonitorSLSGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteHybridMonitorSLSGroupResponseBody) SetCode(v string) *DeleteHybridMonitorSLSGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHybridMonitorSLSGroupResponseBody) SetMessage(v string) *DeleteHybridMonitorSLSGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHybridMonitorSLSGroupResponseBody) SetRequestId(v string) *DeleteHybridMonitorSLSGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHybridMonitorSLSGroupResponseBody) SetSuccess(v string) *DeleteHybridMonitorSLSGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHybridMonitorSLSGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
