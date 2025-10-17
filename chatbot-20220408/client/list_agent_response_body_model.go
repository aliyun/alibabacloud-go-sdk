// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAgentResponseBodyData) *ListAgentResponseBody
	GetData() []*ListAgentResponseBodyData
	SetPageNumber(v int32) *ListAgentResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAgentResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAgentResponseBody
	GetTotalCount() *int32
}

type ListAgentResponseBody struct {
	Data []*ListAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F55D90C1-31BE-4B2A-AA3F-25EFC36F9419
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentResponseBody) GetData() []*ListAgentResponseBodyData {
	return s.Data
}

func (s *ListAgentResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentResponseBody) SetData(v []*ListAgentResponseBodyData) *ListAgentResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentResponseBody) SetPageNumber(v int32) *ListAgentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAgentResponseBody) SetPageSize(v int32) *ListAgentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAgentResponseBody) SetRequestId(v string) *ListAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentResponseBody) SetTotalCount(v int32) *ListAgentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgentResponseBody) Validate() error {
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

type ListAgentResponseBodyData struct {
	// example:
	//
	// 881
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 4e7400028e6f4a7393ed3acf6a7b8927_p_beebot_public
	AgentKey      *string                `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentName     *string                `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	InstanceInfos map[string]interface{} `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty"`
}

func (s ListAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentResponseBodyData) GetAgentId() *int64 {
	return s.AgentId
}

func (s *ListAgentResponseBodyData) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListAgentResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAgentResponseBodyData) GetInstanceInfos() map[string]interface{} {
	return s.InstanceInfos
}

func (s *ListAgentResponseBodyData) SetAgentId(v int64) *ListAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListAgentResponseBodyData) SetAgentKey(v string) *ListAgentResponseBodyData {
	s.AgentKey = &v
	return s
}

func (s *ListAgentResponseBodyData) SetAgentName(v string) *ListAgentResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *ListAgentResponseBodyData) SetInstanceInfos(v map[string]interface{}) *ListAgentResponseBodyData {
	s.InstanceInfos = v
	return s
}

func (s *ListAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
