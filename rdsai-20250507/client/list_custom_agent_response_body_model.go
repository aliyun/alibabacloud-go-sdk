// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListCustomAgentResponseBodyData) *ListCustomAgentResponseBody
	GetData() []*ListCustomAgentResponseBodyData
	SetPageNumber(v int64) *ListCustomAgentResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCustomAgentResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCustomAgentResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCustomAgentResponseBody
	GetTotalCount() *int32
}

type ListCustomAgentResponseBody struct {
	Data []*ListCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponseBody) GetData() []*ListCustomAgentResponseBodyData {
	return s.Data
}

func (s *ListCustomAgentResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCustomAgentResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomAgentResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomAgentResponseBody) SetData(v []*ListCustomAgentResponseBodyData) *ListCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomAgentResponseBody) SetPageNumber(v int64) *ListCustomAgentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomAgentResponseBody) SetPageSize(v int64) *ListCustomAgentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomAgentResponseBody) SetRequestId(v string) *ListCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomAgentResponseBody) SetTotalCount(v int32) *ListCustomAgentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomAgentResponseBody) Validate() error {
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

type ListCustomAgentResponseBodyData struct {
	// example:
	//
	// 2020-11-27 16:01:28
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// true
	EnableTools *bool `json:"EnableTools,omitempty" xml:"EnableTools,omitempty"`
	// AgentId。
	//
	// example:
	//
	// d1b7d639-f34e-44c7-8231-987da14d****
	Id           *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	SystemPrompt *string   `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	Tools        []*string `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-11-27 16:02:28
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListCustomAgentResponseBodyData) GetEnableTools() *bool {
	return s.EnableTools
}

func (s *ListCustomAgentResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListCustomAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListCustomAgentResponseBodyData) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *ListCustomAgentResponseBodyData) GetTools() []*string {
	return s.Tools
}

func (s *ListCustomAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListCustomAgentResponseBodyData) SetCreatedAt(v string) *ListCustomAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetEnableTools(v bool) *ListCustomAgentResponseBodyData {
	s.EnableTools = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetId(v string) *ListCustomAgentResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetName(v string) *ListCustomAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetSystemPrompt(v string) *ListCustomAgentResponseBodyData {
	s.SystemPrompt = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetTools(v []*string) *ListCustomAgentResponseBodyData {
	s.Tools = v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetUpdatedAt(v string) *ListCustomAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
