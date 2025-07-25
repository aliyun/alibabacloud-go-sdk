// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlgorithmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmId(v string) *ListAlgorithmsRequest
	GetAlgorithmId() *string
	SetAlgorithmName(v string) *ListAlgorithmsRequest
	GetAlgorithmName() *string
	SetAlgorithmProvider(v string) *ListAlgorithmsRequest
	GetAlgorithmProvider() *string
	SetPageNumber(v int64) *ListAlgorithmsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlgorithmsRequest
	GetPageSize() *int64
	SetWorkspaceId(v string) *ListAlgorithmsRequest
	GetWorkspaceId() *string
}

type ListAlgorithmsRequest struct {
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// llm_training
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
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
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAlgorithmsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlgorithmsRequest) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsRequest) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *ListAlgorithmsRequest) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *ListAlgorithmsRequest) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *ListAlgorithmsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlgorithmsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlgorithmsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAlgorithmsRequest) SetAlgorithmId(v string) *ListAlgorithmsRequest {
	s.AlgorithmId = &v
	return s
}

func (s *ListAlgorithmsRequest) SetAlgorithmName(v string) *ListAlgorithmsRequest {
	s.AlgorithmName = &v
	return s
}

func (s *ListAlgorithmsRequest) SetAlgorithmProvider(v string) *ListAlgorithmsRequest {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListAlgorithmsRequest) SetPageNumber(v int64) *ListAlgorithmsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAlgorithmsRequest) SetPageSize(v int64) *ListAlgorithmsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlgorithmsRequest) SetWorkspaceId(v string) *ListAlgorithmsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAlgorithmsRequest) Validate() error {
	return dara.Validate(s)
}
