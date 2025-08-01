// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody
	GetData() *CreateNamespaceResponseBodyData
	SetErrorCode(v string) *CreateNamespaceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNamespaceResponseBody
	GetSuccess() *bool
}

type CreateNamespaceResponseBody struct {
	Data *CreateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D1F1A6F3-7E03-5EAD-B3F1-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) GetData() *CreateNamespaceResponseBodyData {
	return s.Data
}

func (s *CreateNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNamespaceResponseBody) SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateNamespaceResponseBody) SetErrorCode(v string) *CreateNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetMessage(v string) *CreateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetSuccess(v bool) *CreateNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateNamespaceResponseBodyData struct {
	// example:
	//
	// myNamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateNamespaceResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *CreateNamespaceResponseBodyData) SetNamespace(v string) *CreateNamespaceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) SetRegion(v string) *CreateNamespaceResponseBodyData {
	s.Region = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
