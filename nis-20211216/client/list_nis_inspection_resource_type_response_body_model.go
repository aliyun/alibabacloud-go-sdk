// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNisInspectionResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListNisInspectionResourceTypeResponseBody
	GetRequestId() *string
	SetResourceTypeList(v []*ListNisInspectionResourceTypeResponseBodyResourceTypeList) *ListNisInspectionResourceTypeResponseBody
	GetResourceTypeList() []*ListNisInspectionResourceTypeResponseBodyResourceTypeList
}

type ListNisInspectionResourceTypeResponseBody struct {
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId        *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceTypeList []*ListNisInspectionResourceTypeResponseBodyResourceTypeList `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
}

func (s ListNisInspectionResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListNisInspectionResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNisInspectionResourceTypeResponseBody) GetResourceTypeList() []*ListNisInspectionResourceTypeResponseBodyResourceTypeList {
	return s.ResourceTypeList
}

func (s *ListNisInspectionResourceTypeResponseBody) SetRequestId(v string) *ListNisInspectionResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNisInspectionResourceTypeResponseBody) SetResourceTypeList(v []*ListNisInspectionResourceTypeResponseBodyResourceTypeList) *ListNisInspectionResourceTypeResponseBody {
	s.ResourceTypeList = v
	return s
}

func (s *ListNisInspectionResourceTypeResponseBody) Validate() error {
	if s.ResourceTypeList != nil {
		for _, item := range s.ResourceTypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNisInspectionResourceTypeResponseBodyResourceTypeList struct {
	// example:
	//
	// EIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListNisInspectionResourceTypeResponseBodyResourceTypeList) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionResourceTypeResponseBodyResourceTypeList) GoString() string {
	return s.String()
}

func (s *ListNisInspectionResourceTypeResponseBodyResourceTypeList) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListNisInspectionResourceTypeResponseBodyResourceTypeList) SetResourceType(v string) *ListNisInspectionResourceTypeResponseBodyResourceTypeList {
	s.ResourceType = &v
	return s
}

func (s *ListNisInspectionResourceTypeResponseBodyResourceTypeList) Validate() error {
	return dara.Validate(s)
}
