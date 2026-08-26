// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaseUserTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*ListSaseUserTagsResponseBodyDataList) *ListSaseUserTagsResponseBody
	GetDataList() []*ListSaseUserTagsResponseBodyDataList
	SetRequestId(v string) *ListSaseUserTagsResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListSaseUserTagsResponseBody
	GetTotalNum() *int32
}

type ListSaseUserTagsResponseBody struct {
	// The list of user labels returned.
	DataList []*ListSaseUserTagsResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of user labels.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListSaseUserTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSaseUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSaseUserTagsResponseBody) GetDataList() []*ListSaseUserTagsResponseBodyDataList {
	return s.DataList
}

func (s *ListSaseUserTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSaseUserTagsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListSaseUserTagsResponseBody) SetDataList(v []*ListSaseUserTagsResponseBodyDataList) *ListSaseUserTagsResponseBody {
	s.DataList = v
	return s
}

func (s *ListSaseUserTagsResponseBody) SetRequestId(v string) *ListSaseUserTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSaseUserTagsResponseBody) SetTotalNum(v int32) *ListSaseUserTagsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListSaseUserTagsResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSaseUserTagsResponseBodyDataList struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 141681795035****
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The number of users associated with the user label.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The description of the user label.
	//
	// example:
	//
	// These are the company\\"s employees
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the user label.
	//
	// example:
	//
	// boss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user label ID.
	//
	// example:
	//
	// su-tag-1ae52f66039fa0d4****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListSaseUserTagsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListSaseUserTagsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListSaseUserTagsResponseBodyDataList) GetAliuid() *string {
	return s.Aliuid
}

func (s *ListSaseUserTagsResponseBodyDataList) GetCount() *int32 {
	return s.Count
}

func (s *ListSaseUserTagsResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListSaseUserTagsResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListSaseUserTagsResponseBodyDataList) GetTagId() *string {
	return s.TagId
}

func (s *ListSaseUserTagsResponseBodyDataList) SetAliuid(v string) *ListSaseUserTagsResponseBodyDataList {
	s.Aliuid = &v
	return s
}

func (s *ListSaseUserTagsResponseBodyDataList) SetCount(v int32) *ListSaseUserTagsResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *ListSaseUserTagsResponseBodyDataList) SetDescription(v string) *ListSaseUserTagsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListSaseUserTagsResponseBodyDataList) SetName(v string) *ListSaseUserTagsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListSaseUserTagsResponseBodyDataList) SetTagId(v string) *ListSaseUserTagsResponseBodyDataList {
	s.TagId = &v
	return s
}

func (s *ListSaseUserTagsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
