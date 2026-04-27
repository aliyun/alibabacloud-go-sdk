// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListFreeAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListFreeAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListFreeAgentResponseBody
	GetCode() *string
	SetData(v *CloudListFreeAgentResponseBodyData) *CloudListFreeAgentResponseBody
	GetData() *CloudListFreeAgentResponseBodyData
	SetMessage(v string) *CloudListFreeAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListFreeAgentResponseBody
	GetRequestId() *string
}

type CloudListFreeAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListFreeAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListFreeAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListFreeAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListFreeAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListFreeAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListFreeAgentResponseBody) GetData() *CloudListFreeAgentResponseBodyData {
	return s.Data
}

func (s *CloudListFreeAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListFreeAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListFreeAgentResponseBody) SetAccessDeniedDetail(v string) *CloudListFreeAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListFreeAgentResponseBody) SetCode(v string) *CloudListFreeAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListFreeAgentResponseBody) SetData(v *CloudListFreeAgentResponseBodyData) *CloudListFreeAgentResponseBody {
	s.Data = v
	return s
}

func (s *CloudListFreeAgentResponseBody) SetMessage(v string) *CloudListFreeAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListFreeAgentResponseBody) SetRequestId(v string) *CloudListFreeAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListFreeAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListFreeAgentResponseBodyData struct {
	List []*CloudListFreeAgentResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudListFreeAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListFreeAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListFreeAgentResponseBodyData) GetList() []*CloudListFreeAgentResponseBodyDataList {
	return s.List
}

func (s *CloudListFreeAgentResponseBodyData) SetList(v []*CloudListFreeAgentResponseBodyDataList) *CloudListFreeAgentResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListFreeAgentResponseBodyData) Validate() error {
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

type CloudListFreeAgentResponseBodyDataList struct {
	// 座席编号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 座席名
	//
	// example:
	//
	// 示例值示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CloudListFreeAgentResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListFreeAgentResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListFreeAgentResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudListFreeAgentResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *CloudListFreeAgentResponseBodyDataList) SetCno(v string) *CloudListFreeAgentResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudListFreeAgentResponseBodyDataList) SetName(v string) *CloudListFreeAgentResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *CloudListFreeAgentResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
