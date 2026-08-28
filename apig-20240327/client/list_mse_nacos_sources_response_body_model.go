// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMseNacosSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMseNacosSourcesResponseBody
	GetCode() *string
	SetData(v *ListMseNacosSourcesResponseBodyData) *ListMseNacosSourcesResponseBody
	GetData() *ListMseNacosSourcesResponseBodyData
	SetMessage(v string) *ListMseNacosSourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMseNacosSourcesResponseBody
	GetRequestId() *string
}

type ListMseNacosSourcesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListMseNacosSourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1994B10-C6A8-58FA-8347-6A08B0D4EFDE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMseNacosSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMseNacosSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMseNacosSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMseNacosSourcesResponseBody) GetData() *ListMseNacosSourcesResponseBodyData {
	return s.Data
}

func (s *ListMseNacosSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMseNacosSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMseNacosSourcesResponseBody) SetCode(v string) *ListMseNacosSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListMseNacosSourcesResponseBody) SetData(v *ListMseNacosSourcesResponseBodyData) *ListMseNacosSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListMseNacosSourcesResponseBody) SetMessage(v string) *ListMseNacosSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListMseNacosSourcesResponseBody) SetRequestId(v string) *ListMseNacosSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMseNacosSourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMseNacosSourcesResponseBodyData struct {
	// The list of Nacos instances that can be added.
	Items []*ListMseNacosSourcesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListMseNacosSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMseNacosSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMseNacosSourcesResponseBodyData) GetItems() []*ListMseNacosSourcesResponseBodyDataItems {
	return s.Items
}

func (s *ListMseNacosSourcesResponseBodyData) SetItems(v []*ListMseNacosSourcesResponseBodyDataItems) *ListMseNacosSourcesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListMseNacosSourcesResponseBodyData) Validate() error {
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

type ListMseNacosSourcesResponseBodyDataItems struct {
	// The Nacos instance ID.
	//
	// example:
	//
	// mse-cn-84a***
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The Nacos instance name.
	//
	// example:
	//
	// 商品中心Nacos注册中心
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListMseNacosSourcesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListMseNacosSourcesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListMseNacosSourcesResponseBodyDataItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMseNacosSourcesResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListMseNacosSourcesResponseBodyDataItems) SetInstanceId(v string) *ListMseNacosSourcesResponseBodyDataItems {
	s.InstanceId = &v
	return s
}

func (s *ListMseNacosSourcesResponseBodyDataItems) SetName(v string) *ListMseNacosSourcesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListMseNacosSourcesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
