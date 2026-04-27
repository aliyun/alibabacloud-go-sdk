// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentCnoAndNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryAgentCnoAndNameResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryAgentCnoAndNameResponseBody
	GetCode() *string
	SetData(v *CloudQueryAgentCnoAndNameResponseBodyData) *CloudQueryAgentCnoAndNameResponseBody
	GetData() *CloudQueryAgentCnoAndNameResponseBodyData
	SetMessage(v string) *CloudQueryAgentCnoAndNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryAgentCnoAndNameResponseBody
	GetRequestId() *string
}

type CloudQueryAgentCnoAndNameResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryAgentCnoAndNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryAgentCnoAndNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentCnoAndNameResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentCnoAndNameResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryAgentCnoAndNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryAgentCnoAndNameResponseBody) GetData() *CloudQueryAgentCnoAndNameResponseBodyData {
	return s.Data
}

func (s *CloudQueryAgentCnoAndNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryAgentCnoAndNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryAgentCnoAndNameResponseBody) SetAccessDeniedDetail(v string) *CloudQueryAgentCnoAndNameResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponseBody) SetCode(v string) *CloudQueryAgentCnoAndNameResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponseBody) SetData(v *CloudQueryAgentCnoAndNameResponseBodyData) *CloudQueryAgentCnoAndNameResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponseBody) SetMessage(v string) *CloudQueryAgentCnoAndNameResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponseBody) SetRequestId(v string) *CloudQueryAgentCnoAndNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryAgentCnoAndNameResponseBodyData struct {
	List []*CloudQueryAgentCnoAndNameResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudQueryAgentCnoAndNameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentCnoAndNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentCnoAndNameResponseBodyData) GetList() []*CloudQueryAgentCnoAndNameResponseBodyDataList {
	return s.List
}

func (s *CloudQueryAgentCnoAndNameResponseBodyData) SetList(v []*CloudQueryAgentCnoAndNameResponseBodyDataList) *CloudQueryAgentCnoAndNameResponseBodyData {
	s.List = v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponseBodyData) Validate() error {
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

type CloudQueryAgentCnoAndNameResponseBodyDataList struct {
	// 坐席工号
	//
	// example:
	//
	// 2000
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 座席名称
	//
	// example:
	//
	// 示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CloudQueryAgentCnoAndNameResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentCnoAndNameResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentCnoAndNameResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryAgentCnoAndNameResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *CloudQueryAgentCnoAndNameResponseBodyDataList) SetCno(v string) *CloudQueryAgentCnoAndNameResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponseBodyDataList) SetName(v string) *CloudQueryAgentCnoAndNameResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *CloudQueryAgentCnoAndNameResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
