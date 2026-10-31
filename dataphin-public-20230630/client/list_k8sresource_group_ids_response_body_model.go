// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8SResourceGroupIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListK8SResourceGroupIdsResponseBody
	GetCode() *string
	SetData(v []*ListK8SResourceGroupIdsResponseBodyData) *ListK8SResourceGroupIdsResponseBody
	GetData() []*ListK8SResourceGroupIdsResponseBodyData
	SetHttpStatusCode(v int32) *ListK8SResourceGroupIdsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListK8SResourceGroupIdsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListK8SResourceGroupIdsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListK8SResourceGroupIdsResponseBody
	GetSuccess() *bool
}

type ListK8SResourceGroupIdsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of custom Kubernetes resource groups.
	Data []*ListK8SResourceGroupIdsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend response error.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListK8SResourceGroupIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListK8SResourceGroupIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListK8SResourceGroupIdsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListK8SResourceGroupIdsResponseBody) GetData() []*ListK8SResourceGroupIdsResponseBodyData {
	return s.Data
}

func (s *ListK8SResourceGroupIdsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListK8SResourceGroupIdsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListK8SResourceGroupIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListK8SResourceGroupIdsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListK8SResourceGroupIdsResponseBody) SetCode(v string) *ListK8SResourceGroupIdsResponseBody {
	s.Code = &v
	return s
}

func (s *ListK8SResourceGroupIdsResponseBody) SetData(v []*ListK8SResourceGroupIdsResponseBodyData) *ListK8SResourceGroupIdsResponseBody {
	s.Data = v
	return s
}

func (s *ListK8SResourceGroupIdsResponseBody) SetHttpStatusCode(v int32) *ListK8SResourceGroupIdsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListK8SResourceGroupIdsResponseBody) SetMessage(v string) *ListK8SResourceGroupIdsResponseBody {
	s.Message = &v
	return s
}

func (s *ListK8SResourceGroupIdsResponseBody) SetRequestId(v string) *ListK8SResourceGroupIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListK8SResourceGroupIdsResponseBody) SetSuccess(v bool) *ListK8SResourceGroupIdsResponseBody {
	s.Success = &v
	return s
}

func (s *ListK8SResourceGroupIdsResponseBody) Validate() error {
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

type ListK8SResourceGroupIdsResponseBodyData struct {
	// The ID of the custom Kubernetes resource group.
	//
	// example:
	//
	// 1052061
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the custom Kubernetes resource group.
	//
	// example:
	//
	// Dedicated to real-time R&D
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListK8SResourceGroupIdsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListK8SResourceGroupIdsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListK8SResourceGroupIdsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListK8SResourceGroupIdsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListK8SResourceGroupIdsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListK8SResourceGroupIdsResponseBodyData) SetId(v string) *ListK8SResourceGroupIdsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListK8SResourceGroupIdsResponseBodyData) SetName(v string) *ListK8SResourceGroupIdsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListK8SResourceGroupIdsResponseBodyData) SetStatus(v string) *ListK8SResourceGroupIdsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListK8SResourceGroupIdsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
