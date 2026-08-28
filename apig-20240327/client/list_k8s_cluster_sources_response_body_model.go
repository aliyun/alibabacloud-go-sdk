// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sClusterSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListK8sClusterSourcesResponseBody
	GetCode() *string
	SetData(v *ListK8sClusterSourcesResponseBodyData) *ListK8sClusterSourcesResponseBody
	GetData() *ListK8sClusterSourcesResponseBodyData
	SetMessage(v string) *ListK8sClusterSourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListK8sClusterSourcesResponseBody
	GetRequestId() *string
}

type ListK8sClusterSourcesResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListK8sClusterSourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 043360DA-ED3B-5386-9B7A-D94DECF99A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListK8sClusterSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListK8sClusterSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListK8sClusterSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListK8sClusterSourcesResponseBody) GetData() *ListK8sClusterSourcesResponseBodyData {
	return s.Data
}

func (s *ListK8sClusterSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListK8sClusterSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListK8sClusterSourcesResponseBody) SetCode(v string) *ListK8sClusterSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListK8sClusterSourcesResponseBody) SetData(v *ListK8sClusterSourcesResponseBodyData) *ListK8sClusterSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListK8sClusterSourcesResponseBody) SetMessage(v string) *ListK8sClusterSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListK8sClusterSourcesResponseBody) SetRequestId(v string) *ListK8sClusterSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListK8sClusterSourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListK8sClusterSourcesResponseBodyData struct {
	Items []*ListK8sClusterSourcesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListK8sClusterSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListK8sClusterSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListK8sClusterSourcesResponseBodyData) GetItems() []*ListK8sClusterSourcesResponseBodyDataItems {
	return s.Items
}

func (s *ListK8sClusterSourcesResponseBodyData) SetItems(v []*ListK8sClusterSourcesResponseBodyDataItems) *ListK8sClusterSourcesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListK8sClusterSourcesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sClusterSourcesResponseBodyDataItems struct {
	// example:
	//
	// c4a21b3560fad4ec***
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// itemcenter-dev-cluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListK8sClusterSourcesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListK8sClusterSourcesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListK8sClusterSourcesResponseBodyDataItems) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sClusterSourcesResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListK8sClusterSourcesResponseBodyDataItems) SetClusterId(v string) *ListK8sClusterSourcesResponseBodyDataItems {
	s.ClusterId = &v
	return s
}

func (s *ListK8sClusterSourcesResponseBodyDataItems) SetName(v string) *ListK8sClusterSourcesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListK8sClusterSourcesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
