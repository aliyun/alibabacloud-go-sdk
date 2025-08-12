// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridMonitorNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHybridMonitorNamespaceResponseBody
	GetCode() *string
	SetMessage(v string) *CreateHybridMonitorNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHybridMonitorNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateHybridMonitorNamespaceResponseBody
	GetSuccess() *string
}

type CreateHybridMonitorNamespaceResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Namespace.Existed
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

func (s CreateHybridMonitorNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHybridMonitorNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHybridMonitorNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHybridMonitorNamespaceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateHybridMonitorNamespaceResponseBody) SetCode(v string) *CreateHybridMonitorNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHybridMonitorNamespaceResponseBody) SetMessage(v string) *CreateHybridMonitorNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHybridMonitorNamespaceResponseBody) SetRequestId(v string) *CreateHybridMonitorNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridMonitorNamespaceResponseBody) SetSuccess(v string) *CreateHybridMonitorNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHybridMonitorNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
