// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListIvrNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListIvrNodesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListIvrNodesResponseBody
	GetCode() *string
	SetData(v *ClinkListIvrNodesResponseBodyData) *ClinkListIvrNodesResponseBody
	GetData() *ClinkListIvrNodesResponseBodyData
	SetMessage(v string) *ClinkListIvrNodesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListIvrNodesResponseBody
	GetRequestId() *string
}

type ClinkListIvrNodesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListIvrNodesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListIvrNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListIvrNodesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListIvrNodesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListIvrNodesResponseBody) GetData() *ClinkListIvrNodesResponseBodyData {
	return s.Data
}

func (s *ClinkListIvrNodesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListIvrNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListIvrNodesResponseBody) SetAccessDeniedDetail(v string) *ClinkListIvrNodesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListIvrNodesResponseBody) SetCode(v string) *ClinkListIvrNodesResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListIvrNodesResponseBody) SetData(v *ClinkListIvrNodesResponseBodyData) *ClinkListIvrNodesResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListIvrNodesResponseBody) SetMessage(v string) *ClinkListIvrNodesResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListIvrNodesResponseBody) SetRequestId(v string) *ClinkListIvrNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListIvrNodesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListIvrNodesResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 语音导航节点列表
	IvrNodes []*ClinkListIvrNodesResponseBodyDataIvrNodes `json:"IvrNodes,omitempty" xml:"IvrNodes,omitempty" type:"Repeated"`
}

func (s ClinkListIvrNodesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrNodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListIvrNodesResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListIvrNodesResponseBodyData) GetIvrNodes() []*ClinkListIvrNodesResponseBodyDataIvrNodes {
	return s.IvrNodes
}

func (s *ClinkListIvrNodesResponseBodyData) SetClinkRequestId(v string) *ClinkListIvrNodesResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListIvrNodesResponseBodyData) SetIvrNodes(v []*ClinkListIvrNodesResponseBodyDataIvrNodes) *ClinkListIvrNodesResponseBodyData {
	s.IvrNodes = v
	return s
}

func (s *ClinkListIvrNodesResponseBodyData) Validate() error {
	if s.IvrNodes != nil {
		for _, item := range s.IvrNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListIvrNodesResponseBodyDataIvrNodes struct {
	// 语音导航节点
	//
	// example:
	//
	// E91_rp0D1
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// 语音导航常用节点
	//
	// example:
	//
	// null
	FrequentlyPath *string `json:"FrequentlyPath,omitempty" xml:"FrequentlyPath,omitempty"`
	// 语音导航节点id
	//
	// example:
	//
	// 7
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 语音导航id
	//
	// example:
	//
	// 53
	IvrId *int64 `json:"IvrId,omitempty" xml:"IvrId,omitempty"`
	// 语音导航节点名称
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

func (s ClinkListIvrNodesResponseBodyDataIvrNodes) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrNodesResponseBodyDataIvrNodes) GoString() string {
	return s.String()
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) GetFrequentlyPath() *string {
	return s.FrequentlyPath
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) GetId() *int64 {
	return s.Id
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) GetIvrId() *int64 {
	return s.IvrId
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) GetName() *string {
	return s.Name
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) GetStatistic() *int64 {
	return s.Statistic
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) SetEndpoint(v string) *ClinkListIvrNodesResponseBodyDataIvrNodes {
	s.Endpoint = &v
	return s
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) SetFrequentlyPath(v string) *ClinkListIvrNodesResponseBodyDataIvrNodes {
	s.FrequentlyPath = &v
	return s
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) SetId(v int64) *ClinkListIvrNodesResponseBodyDataIvrNodes {
	s.Id = &v
	return s
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) SetIvrId(v int64) *ClinkListIvrNodesResponseBodyDataIvrNodes {
	s.IvrId = &v
	return s
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) SetName(v string) *ClinkListIvrNodesResponseBodyDataIvrNodes {
	s.Name = &v
	return s
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) SetStatistic(v int64) *ClinkListIvrNodesResponseBodyDataIvrNodes {
	s.Statistic = &v
	return s
}

func (s *ClinkListIvrNodesResponseBodyDataIvrNodes) Validate() error {
	return dara.Validate(s)
}
