// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDagComplementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizBeginTime(v string) *CreateDagComplementRequest
	GetBizBeginTime() *string
	SetBizEndTime(v string) *CreateDagComplementRequest
	GetBizEndTime() *string
	SetEndBizDate(v string) *CreateDagComplementRequest
	GetEndBizDate() *string
	SetExcludeNodeIds(v string) *CreateDagComplementRequest
	GetExcludeNodeIds() *string
	SetIncludeNodeIds(v string) *CreateDagComplementRequest
	GetIncludeNodeIds() *string
	SetName(v string) *CreateDagComplementRequest
	GetName() *string
	SetNodeParams(v string) *CreateDagComplementRequest
	GetNodeParams() *string
	SetParallelism(v bool) *CreateDagComplementRequest
	GetParallelism() *bool
	SetProjectEnv(v string) *CreateDagComplementRequest
	GetProjectEnv() *string
	SetRootNodeId(v int64) *CreateDagComplementRequest
	GetRootNodeId() *int64
	SetStartBizDate(v string) *CreateDagComplementRequest
	GetStartBizDate() *string
}

type CreateDagComplementRequest struct {
	// example:
	//
	// 00:00:00
	BizBeginTime *string `json:"BizBeginTime,omitempty" xml:"BizBeginTime,omitempty"`
	// example:
	//
	// 23:00:00
	BizEndTime *string `json:"BizEndTime,omitempty" xml:"BizEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-21 00:00:00
	EndBizDate *string `json:"EndBizDate,omitempty" xml:"EndBizDate,omitempty"`
	// example:
	//
	// 1234
	ExcludeNodeIds *string `json:"ExcludeNodeIds,omitempty" xml:"ExcludeNodeIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74324,74325
	IncludeNodeIds *string `json:"IncludeNodeIds,omitempty" xml:"IncludeNodeIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xm_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {74324:"key1=val1 key2=val"}
	NodeParams *string `json:"NodeParams,omitempty" xml:"NodeParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Parallelism *bool `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	RootNodeId *int64 `json:"RootNodeId,omitempty" xml:"RootNodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-20 00:00:00
	StartBizDate *string `json:"StartBizDate,omitempty" xml:"StartBizDate,omitempty"`
}

func (s CreateDagComplementRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDagComplementRequest) GoString() string {
	return s.String()
}

func (s *CreateDagComplementRequest) GetBizBeginTime() *string {
	return s.BizBeginTime
}

func (s *CreateDagComplementRequest) GetBizEndTime() *string {
	return s.BizEndTime
}

func (s *CreateDagComplementRequest) GetEndBizDate() *string {
	return s.EndBizDate
}

func (s *CreateDagComplementRequest) GetExcludeNodeIds() *string {
	return s.ExcludeNodeIds
}

func (s *CreateDagComplementRequest) GetIncludeNodeIds() *string {
	return s.IncludeNodeIds
}

func (s *CreateDagComplementRequest) GetName() *string {
	return s.Name
}

func (s *CreateDagComplementRequest) GetNodeParams() *string {
	return s.NodeParams
}

func (s *CreateDagComplementRequest) GetParallelism() *bool {
	return s.Parallelism
}

func (s *CreateDagComplementRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *CreateDagComplementRequest) GetRootNodeId() *int64 {
	return s.RootNodeId
}

func (s *CreateDagComplementRequest) GetStartBizDate() *string {
	return s.StartBizDate
}

func (s *CreateDagComplementRequest) SetBizBeginTime(v string) *CreateDagComplementRequest {
	s.BizBeginTime = &v
	return s
}

func (s *CreateDagComplementRequest) SetBizEndTime(v string) *CreateDagComplementRequest {
	s.BizEndTime = &v
	return s
}

func (s *CreateDagComplementRequest) SetEndBizDate(v string) *CreateDagComplementRequest {
	s.EndBizDate = &v
	return s
}

func (s *CreateDagComplementRequest) SetExcludeNodeIds(v string) *CreateDagComplementRequest {
	s.ExcludeNodeIds = &v
	return s
}

func (s *CreateDagComplementRequest) SetIncludeNodeIds(v string) *CreateDagComplementRequest {
	s.IncludeNodeIds = &v
	return s
}

func (s *CreateDagComplementRequest) SetName(v string) *CreateDagComplementRequest {
	s.Name = &v
	return s
}

func (s *CreateDagComplementRequest) SetNodeParams(v string) *CreateDagComplementRequest {
	s.NodeParams = &v
	return s
}

func (s *CreateDagComplementRequest) SetParallelism(v bool) *CreateDagComplementRequest {
	s.Parallelism = &v
	return s
}

func (s *CreateDagComplementRequest) SetProjectEnv(v string) *CreateDagComplementRequest {
	s.ProjectEnv = &v
	return s
}

func (s *CreateDagComplementRequest) SetRootNodeId(v int64) *CreateDagComplementRequest {
	s.RootNodeId = &v
	return s
}

func (s *CreateDagComplementRequest) SetStartBizDate(v string) *CreateDagComplementRequest {
	s.StartBizDate = &v
	return s
}

func (s *CreateDagComplementRequest) Validate() error {
	return dara.Validate(s)
}
