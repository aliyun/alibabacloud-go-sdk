// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryStoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMemoryStoresResponseBody
	GetMaxResults() *int32
	SetMemoryStores(v []*ListMemoryStoresResponseBodyMemoryStores) *ListMemoryStoresResponseBody
	GetMemoryStores() []*ListMemoryStoresResponseBodyMemoryStores
	SetNextToken(v string) *ListMemoryStoresResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMemoryStoresResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListMemoryStoresResponseBody
	GetTotal() *int32
}

type ListMemoryStoresResponseBody struct {
	// example:
	//
	// 1
	MaxResults   *int32                                      `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	MemoryStores []*ListMemoryStoresResponseBodyMemoryStores `json:"memoryStores,omitempty" xml:"memoryStores,omitempty" type:"Repeated"`
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 16C0A6D6-C3E7-511D-A60B-A87FD85F5BA7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Total     *int32  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMemoryStoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoryStoresResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoryStoresResponseBody) GetMemoryStores() []*ListMemoryStoresResponseBodyMemoryStores {
	return s.MemoryStores
}

func (s *ListMemoryStoresResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryStoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoryStoresResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListMemoryStoresResponseBody) SetMaxResults(v int32) *ListMemoryStoresResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetMemoryStores(v []*ListMemoryStoresResponseBodyMemoryStores) *ListMemoryStoresResponseBody {
	s.MemoryStores = v
	return s
}

func (s *ListMemoryStoresResponseBody) SetNextToken(v string) *ListMemoryStoresResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetRequestId(v string) *ListMemoryStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoryStoresResponseBody) SetTotal(v int32) *ListMemoryStoresResponseBody {
	s.Total = &v
	return s
}

func (s *ListMemoryStoresResponseBody) Validate() error {
	if s.MemoryStores != nil {
		for _, item := range s.MemoryStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMemoryStoresResponseBodyMemoryStores struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1764556182850
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// Created by taishan-module-recovery
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// qianyi_test_1
	MemoryStoreName *string `json:"memoryStoreName,omitempty" xml:"memoryStoreName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 1764556182850
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// default-cms-1646467597142798-cn-shenzhen
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListMemoryStoresResponseBodyMemoryStores) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryStoresResponseBodyMemoryStores) GoString() string {
	return s.String()
}

func (s *ListMemoryStoresResponseBodyMemoryStores) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMemoryStoresResponseBodyMemoryStores) GetDescription() *string {
	return s.Description
}

func (s *ListMemoryStoresResponseBodyMemoryStores) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *ListMemoryStoresResponseBodyMemoryStores) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMemoryStoresResponseBodyMemoryStores) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListMemoryStoresResponseBodyMemoryStores) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListMemoryStoresResponseBodyMemoryStores) SetCreateTime(v string) *ListMemoryStoresResponseBodyMemoryStores {
	s.CreateTime = &v
	return s
}

func (s *ListMemoryStoresResponseBodyMemoryStores) SetDescription(v string) *ListMemoryStoresResponseBodyMemoryStores {
	s.Description = &v
	return s
}

func (s *ListMemoryStoresResponseBodyMemoryStores) SetMemoryStoreName(v string) *ListMemoryStoresResponseBodyMemoryStores {
	s.MemoryStoreName = &v
	return s
}

func (s *ListMemoryStoresResponseBodyMemoryStores) SetRegionId(v string) *ListMemoryStoresResponseBodyMemoryStores {
	s.RegionId = &v
	return s
}

func (s *ListMemoryStoresResponseBodyMemoryStores) SetUpdateTime(v string) *ListMemoryStoresResponseBodyMemoryStores {
	s.UpdateTime = &v
	return s
}

func (s *ListMemoryStoresResponseBodyMemoryStores) SetWorkspace(v string) *ListMemoryStoresResponseBodyMemoryStores {
	s.Workspace = &v
	return s
}

func (s *ListMemoryStoresResponseBodyMemoryStores) Validate() error {
	return dara.Validate(s)
}
