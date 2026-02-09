// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCategoriesResponseBody
	GetCode() *int32
	SetData(v *ListCategoriesResponseBodyData) *ListCategoriesResponseBody
	GetData() *ListCategoriesResponseBodyData
	SetMessage(v string) *ListCategoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCategoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCategoriesResponseBody
	GetSuccess() *bool
}

type ListCategoriesResponseBody struct {
	// example:
	//
	// 200
	Code *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCategoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CA6204AC-6AA9-4CFA-9310-7DFD20C19EBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCategoriesResponseBody) GetData() *ListCategoriesResponseBodyData {
	return s.Data
}

func (s *ListCategoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCategoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCategoriesResponseBody) SetCode(v int32) *ListCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCategoriesResponseBody) SetData(v *ListCategoriesResponseBodyData) *ListCategoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListCategoriesResponseBody) SetMessage(v string) *ListCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCategoriesResponseBody) SetRequestId(v string) *ListCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoriesResponseBody) SetSuccess(v bool) *ListCategoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCategoriesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCategoriesResponseBodyData struct {
	List []*ListCategoriesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListCategoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBodyData) GetList() []*ListCategoriesResponseBodyDataList {
	return s.List
}

func (s *ListCategoriesResponseBodyData) SetList(v []*ListCategoriesResponseBodyDataList) *ListCategoriesResponseBodyData {
	s.List = v
	return s
}

func (s *ListCategoriesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCategoriesResponseBodyDataList struct {
	// example:
	//
	// 793
	Id   *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListCategoriesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBodyDataList) GetId() *int32 {
	return s.Id
}

func (s *ListCategoriesResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListCategoriesResponseBodyDataList) SetId(v int32) *ListCategoriesResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListCategoriesResponseBodyDataList) SetName(v string) *ListCategoriesResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListCategoriesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
