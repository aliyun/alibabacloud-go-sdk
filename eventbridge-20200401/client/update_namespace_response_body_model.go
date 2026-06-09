// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateNamespaceResponseBody
	GetCode() *string
	SetData(v *UpdateNamespaceResponseBodyData) *UpdateNamespaceResponseBody
	GetData() *UpdateNamespaceResponseBodyData
	SetMessage(v string) *UpdateNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNamespaceResponseBody
	GetSuccess() *bool
}

type UpdateNamespaceResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateNamespaceResponseBody) GetData() *UpdateNamespaceResponseBodyData {
	return s.Data
}

func (s *UpdateNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNamespaceResponseBody) SetCode(v string) *UpdateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetData(v *UpdateNamespaceResponseBodyData) *UpdateNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateNamespaceResponseBody) SetMessage(v string) *UpdateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetRequestId(v string) *UpdateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetSuccess(v bool) *UpdateNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateNamespaceResponseBodyData struct {
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789:catalog/my_catalog/namespace/my_namespace
	NamespaceARN *string `json:"NamespaceARN,omitempty" xml:"NamespaceARN,omitempty"`
}

func (s UpdateNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBodyData) GetNamespaceARN() *string {
	return s.NamespaceARN
}

func (s *UpdateNamespaceResponseBodyData) SetNamespaceARN(v string) *UpdateNamespaceResponseBodyData {
	s.NamespaceARN = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
