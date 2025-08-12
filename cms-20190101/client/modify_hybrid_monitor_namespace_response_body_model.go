// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridMonitorNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyHybridMonitorNamespaceResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyHybridMonitorNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyHybridMonitorNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyHybridMonitorNamespaceResponseBody
	GetSuccess() *string
}

type ModifyHybridMonitorNamespaceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// %s
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E190AB2E-7BF9-59B7-9DDC-7CB1782C5ECD
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

func (s ModifyHybridMonitorNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyHybridMonitorNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyHybridMonitorNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridMonitorNamespaceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyHybridMonitorNamespaceResponseBody) SetCode(v string) *ModifyHybridMonitorNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyHybridMonitorNamespaceResponseBody) SetMessage(v string) *ModifyHybridMonitorNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyHybridMonitorNamespaceResponseBody) SetRequestId(v string) *ModifyHybridMonitorNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridMonitorNamespaceResponseBody) SetSuccess(v string) *ModifyHybridMonitorNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyHybridMonitorNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
