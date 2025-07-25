// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlgorithmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithms(v []*ListAlgorithmsResponseBodyAlgorithms) *ListAlgorithmsResponseBody
	GetAlgorithms() []*ListAlgorithmsResponseBodyAlgorithms
	SetRequestId(v string) *ListAlgorithmsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAlgorithmsResponseBody
	GetTotalCount() *int64
}

type ListAlgorithmsResponseBody struct {
	Algorithms []*ListAlgorithmsResponseBodyAlgorithms `json:"Algorithms,omitempty" xml:"Algorithms,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlgorithmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlgorithmsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponseBody) GetAlgorithms() []*ListAlgorithmsResponseBodyAlgorithms {
	return s.Algorithms
}

func (s *ListAlgorithmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlgorithmsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAlgorithmsResponseBody) SetAlgorithms(v []*ListAlgorithmsResponseBodyAlgorithms) *ListAlgorithmsResponseBody {
	s.Algorithms = v
	return s
}

func (s *ListAlgorithmsResponseBody) SetRequestId(v string) *ListAlgorithmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlgorithmsResponseBody) SetTotalCount(v int64) *ListAlgorithmsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlgorithmsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlgorithmsResponseBodyAlgorithms struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	// example:
	//
	// algo-sidjc8134hv
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_train
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// LLM Train
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2023-07-21T03:35:24Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-07-25T02:15:40Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAlgorithmsResponseBodyAlgorithms) String() string {
	return dara.Prettify(s)
}

func (s ListAlgorithmsResponseBodyAlgorithms) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponseBodyAlgorithms) GetAlgorithmDescription() *string {
	return s.AlgorithmDescription
}

func (s *ListAlgorithmsResponseBodyAlgorithms) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *ListAlgorithmsResponseBodyAlgorithms) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *ListAlgorithmsResponseBodyAlgorithms) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *ListAlgorithmsResponseBodyAlgorithms) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAlgorithmsResponseBodyAlgorithms) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListAlgorithmsResponseBodyAlgorithms) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListAlgorithmsResponseBodyAlgorithms) GetUserId() *string {
	return s.UserId
}

func (s *ListAlgorithmsResponseBodyAlgorithms) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetAlgorithmDescription(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.AlgorithmDescription = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetAlgorithmId(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.AlgorithmId = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetAlgorithmName(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.AlgorithmName = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetAlgorithmProvider(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetDisplayName(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.DisplayName = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetGmtCreateTime(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.GmtCreateTime = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetGmtModifiedTime(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetUserId(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.UserId = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) SetWorkspaceId(v string) *ListAlgorithmsResponseBodyAlgorithms {
	s.WorkspaceId = &v
	return s
}

func (s *ListAlgorithmsResponseBodyAlgorithms) Validate() error {
	return dara.Validate(s)
}
