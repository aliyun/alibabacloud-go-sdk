// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridMonitorNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHybridMonitorNamespaceResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteHybridMonitorNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHybridMonitorNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteHybridMonitorNamespaceResponseBody
	GetSuccess() *string
}

type DeleteHybridMonitorNamespaceResponseBody struct {
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
	// Namespace.NotFound
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3843D23A-FB9E-5DC8-BCCC-458768B79296
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

func (s DeleteHybridMonitorNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridMonitorNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHybridMonitorNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHybridMonitorNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHybridMonitorNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHybridMonitorNamespaceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteHybridMonitorNamespaceResponseBody) SetCode(v string) *DeleteHybridMonitorNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHybridMonitorNamespaceResponseBody) SetMessage(v string) *DeleteHybridMonitorNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHybridMonitorNamespaceResponseBody) SetRequestId(v string) *DeleteHybridMonitorNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHybridMonitorNamespaceResponseBody) SetSuccess(v string) *DeleteHybridMonitorNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHybridMonitorNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
