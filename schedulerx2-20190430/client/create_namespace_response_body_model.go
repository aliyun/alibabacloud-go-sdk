// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateNamespaceResponseBody
	GetCode() *int32
	SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody
	GetData() *CreateNamespaceResponseBodyData
	SetMessage(v string) *CreateNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNamespaceResponseBody
	GetSuccess() *bool
}

type CreateNamespaceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the namespace.
	Data *CreateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// namespace=test3 is existed, noting update
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application was created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
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

func (s *CreateNamespaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateNamespaceResponseBody) GetData() *CreateNamespaceResponseBodyData {
	return s.Data
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

func (s *CreateNamespaceResponseBody) SetCode(v int32) *CreateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody {
	s.Data = v
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
	// The UID of the namespace.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	NamespaceUid *string `json:"NamespaceUid,omitempty" xml:"NamespaceUid,omitempty"`
}

func (s CreateNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBodyData) GetNamespaceUid() *string {
	return s.NamespaceUid
}

func (s *CreateNamespaceResponseBodyData) SetNamespaceUid(v string) *CreateNamespaceResponseBodyData {
	s.NamespaceUid = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
