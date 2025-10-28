// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListK8sNamespacesResponseBody
	GetCode() *int32
	SetData(v []*ListK8sNamespacesResponseBodyData) *ListK8sNamespacesResponseBody
	GetData() []*ListK8sNamespacesResponseBodyData
	SetMessage(v string) *ListK8sNamespacesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListK8sNamespacesResponseBody
	GetRequestId() *string
}

type ListK8sNamespacesResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data []*ListK8sNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 57F146F6-3C94-****-****-A66EF4B9*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListK8sNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListK8sNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListK8sNamespacesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListK8sNamespacesResponseBody) GetData() []*ListK8sNamespacesResponseBodyData {
	return s.Data
}

func (s *ListK8sNamespacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListK8sNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListK8sNamespacesResponseBody) SetCode(v int32) *ListK8sNamespacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListK8sNamespacesResponseBody) SetData(v []*ListK8sNamespacesResponseBodyData) *ListK8sNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListK8sNamespacesResponseBody) SetMessage(v string) *ListK8sNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListK8sNamespacesResponseBody) SetRequestId(v string) *ListK8sNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListK8sNamespacesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sNamespacesResponseBodyData struct {
	// The namespaces of the Kubernetes cluster.
	//
	// example:
	//
	// development
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListK8sNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListK8sNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListK8sNamespacesResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *ListK8sNamespacesResponseBodyData) SetNamespace(v string) *ListK8sNamespacesResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListK8sNamespacesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
