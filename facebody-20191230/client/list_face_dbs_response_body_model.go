// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaceDbsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListFaceDbsResponseBodyData) *ListFaceDbsResponseBody
	GetData() *ListFaceDbsResponseBodyData
	SetRequestId(v string) *ListFaceDbsResponseBody
	GetRequestId() *string
}

type ListFaceDbsResponseBody struct {
	Data *ListFaceDbsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2B93C43A-F824-40C8-AF79-844342B0F43A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFaceDbsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFaceDbsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponseBody) GetData() *ListFaceDbsResponseBodyData {
	return s.Data
}

func (s *ListFaceDbsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFaceDbsResponseBody) SetData(v *ListFaceDbsResponseBodyData) *ListFaceDbsResponseBody {
	s.Data = v
	return s
}

func (s *ListFaceDbsResponseBody) SetRequestId(v string) *ListFaceDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaceDbsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFaceDbsResponseBodyData struct {
	DbList []*ListFaceDbsResponseBodyDataDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
}

func (s ListFaceDbsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFaceDbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponseBodyData) GetDbList() []*ListFaceDbsResponseBodyDataDbList {
	return s.DbList
}

func (s *ListFaceDbsResponseBodyData) SetDbList(v []*ListFaceDbsResponseBodyDataDbList) *ListFaceDbsResponseBodyData {
	s.DbList = v
	return s
}

func (s *ListFaceDbsResponseBodyData) Validate() error {
	if s.DbList != nil {
		for _, item := range s.DbList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFaceDbsResponseBodyDataDbList struct {
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListFaceDbsResponseBodyDataDbList) String() string {
	return dara.Prettify(s)
}

func (s ListFaceDbsResponseBodyDataDbList) GoString() string {
	return s.String()
}

func (s *ListFaceDbsResponseBodyDataDbList) GetName() *string {
	return s.Name
}

func (s *ListFaceDbsResponseBodyDataDbList) SetName(v string) *ListFaceDbsResponseBodyDataDbList {
	s.Name = &v
	return s
}

func (s *ListFaceDbsResponseBodyDataDbList) Validate() error {
	return dara.Validate(s)
}
