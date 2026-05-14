// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListIvrsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListIvrsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListIvrsResponseBody
	GetCode() *string
	SetData(v *ClinkListIvrsResponseBodyData) *ClinkListIvrsResponseBody
	GetData() *ClinkListIvrsResponseBodyData
	SetMessage(v string) *ClinkListIvrsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListIvrsResponseBody
	GetRequestId() *string
}

type ClinkListIvrsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListIvrsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListIvrsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrsResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListIvrsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListIvrsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListIvrsResponseBody) GetData() *ClinkListIvrsResponseBodyData {
	return s.Data
}

func (s *ClinkListIvrsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListIvrsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListIvrsResponseBody) SetAccessDeniedDetail(v string) *ClinkListIvrsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListIvrsResponseBody) SetCode(v string) *ClinkListIvrsResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListIvrsResponseBody) SetData(v *ClinkListIvrsResponseBodyData) *ClinkListIvrsResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListIvrsResponseBody) SetMessage(v string) *ClinkListIvrsResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListIvrsResponseBody) SetRequestId(v string) *ClinkListIvrsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListIvrsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListIvrsResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 语音导航列表
	Ivrs []*ClinkListIvrsResponseBodyDataIvrs `json:"Ivrs,omitempty" xml:"Ivrs,omitempty" type:"Repeated"`
}

func (s ClinkListIvrsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListIvrsResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListIvrsResponseBodyData) GetIvrs() []*ClinkListIvrsResponseBodyDataIvrs {
	return s.Ivrs
}

func (s *ClinkListIvrsResponseBodyData) SetClinkRequestId(v string) *ClinkListIvrsResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListIvrsResponseBodyData) SetIvrs(v []*ClinkListIvrsResponseBodyDataIvrs) *ClinkListIvrsResponseBodyData {
	s.Ivrs = v
	return s
}

func (s *ClinkListIvrsResponseBodyData) Validate() error {
	if s.Ivrs != nil {
		for _, item := range s.Ivrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListIvrsResponseBodyDataIvrs struct {
	// 语音导航描述
	//
	// example:
	//
	// 示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 语音导航id
	//
	// example:
	//
	// 7
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 语音导航名称
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 是否开启节点统计，0：不开启；1：开启
	//
	// example:
	//
	// 0
	Statistic *int64 `json:"Statistic,omitempty" xml:"Statistic,omitempty"`
}

func (s ClinkListIvrsResponseBodyDataIvrs) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrsResponseBodyDataIvrs) GoString() string {
	return s.String()
}

func (s *ClinkListIvrsResponseBodyDataIvrs) GetDescription() *string {
	return s.Description
}

func (s *ClinkListIvrsResponseBodyDataIvrs) GetId() *int64 {
	return s.Id
}

func (s *ClinkListIvrsResponseBodyDataIvrs) GetName() *string {
	return s.Name
}

func (s *ClinkListIvrsResponseBodyDataIvrs) GetStatistic() *int64 {
	return s.Statistic
}

func (s *ClinkListIvrsResponseBodyDataIvrs) SetDescription(v string) *ClinkListIvrsResponseBodyDataIvrs {
	s.Description = &v
	return s
}

func (s *ClinkListIvrsResponseBodyDataIvrs) SetId(v int64) *ClinkListIvrsResponseBodyDataIvrs {
	s.Id = &v
	return s
}

func (s *ClinkListIvrsResponseBodyDataIvrs) SetName(v string) *ClinkListIvrsResponseBodyDataIvrs {
	s.Name = &v
	return s
}

func (s *ClinkListIvrsResponseBodyDataIvrs) SetStatistic(v int64) *ClinkListIvrsResponseBodyDataIvrs {
	s.Statistic = &v
	return s
}

func (s *ClinkListIvrsResponseBodyDataIvrs) Validate() error {
	return dara.Validate(s)
}
