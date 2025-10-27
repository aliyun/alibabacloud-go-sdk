// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListNamespacesResponseBody
	GetCode() *int32
	SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody
	GetData() *ListNamespacesResponseBodyData
	SetMessage(v string) *ListNamespacesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListNamespacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNamespacesResponseBody
	GetSuccess() *bool
}

type ListNamespacesResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the namespaces.
	Data *ListNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB58C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s ListNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListNamespacesResponseBody) GetData() *ListNamespacesResponseBodyData {
	return s.Data
}

func (s *ListNamespacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamespacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNamespacesResponseBody) SetCode(v int32) *ListNamespacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListNamespacesResponseBody) SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespacesResponseBody) SetMessage(v string) *ListNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacesResponseBody) SetSuccess(v bool) *ListNamespacesResponseBody {
	s.Success = &v
	return s
}

func (s *ListNamespacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNamespacesResponseBodyData struct {
	// The namespaces and their details.
	Namespaces []*ListNamespacesResponseBodyDataNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
}

func (s ListNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyData) GetNamespaces() []*ListNamespacesResponseBodyDataNamespaces {
	return s.Namespaces
}

func (s *ListNamespacesResponseBodyData) SetNamespaces(v []*ListNamespacesResponseBodyDataNamespaces) *ListNamespacesResponseBodyData {
	s.Namespaces = v
	return s
}

func (s *ListNamespacesResponseBodyData) Validate() error {
	if s.Namespaces != nil {
		for _, item := range s.Namespaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNamespacesResponseBodyDataNamespaces struct {
	// The description of the namespace.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// doc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// 1a72ecb1-b4cc-400a-a71b-20cdec9b****
	UId *string `json:"UId,omitempty" xml:"UId,omitempty"`
}

func (s ListNamespacesResponseBodyDataNamespaces) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyDataNamespaces) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyDataNamespaces) GetDescription() *string {
	return s.Description
}

func (s *ListNamespacesResponseBodyDataNamespaces) GetName() *string {
	return s.Name
}

func (s *ListNamespacesResponseBodyDataNamespaces) GetUId() *string {
	return s.UId
}

func (s *ListNamespacesResponseBodyDataNamespaces) SetDescription(v string) *ListNamespacesResponseBodyDataNamespaces {
	s.Description = &v
	return s
}

func (s *ListNamespacesResponseBodyDataNamespaces) SetName(v string) *ListNamespacesResponseBodyDataNamespaces {
	s.Name = &v
	return s
}

func (s *ListNamespacesResponseBodyDataNamespaces) SetUId(v string) *ListNamespacesResponseBodyDataNamespaces {
	s.UId = &v
	return s
}

func (s *ListNamespacesResponseBodyDataNamespaces) Validate() error {
	return dara.Validate(s)
}
