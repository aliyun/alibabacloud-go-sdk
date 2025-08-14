// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfCreateTime(v string) *ListWorkflowTasksRequest
	GetEndOfCreateTime() *string
	SetKeyText(v string) *ListWorkflowTasksRequest
	GetKeyText() *string
	SetMaxResults(v int32) *ListWorkflowTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkflowTasksRequest
	GetNextToken() *string
	SetStartOfCreateTime(v string) *ListWorkflowTasksRequest
	GetStartOfCreateTime() *string
	SetWorkflowId(v string) *ListWorkflowTasksRequest
	GetWorkflowId() *string
	SetWorkflowName(v string) *ListWorkflowTasksRequest
	GetWorkflowName() *string
}

type ListWorkflowTasksRequest struct {
	// example:
	//
	// 2025-07-15T00:00:00Z
	EndOfCreateTime *string `json:"EndOfCreateTime,omitempty" xml:"EndOfCreateTime,omitempty"`
	KeyText         *string `json:"KeyText,omitempty" xml:"KeyText,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// **************VRpbWUQARgBIpcBCgkA1bUtaAAAAAAKiQEDhAAAADFTMzg2NTY2NjU2MzM3NjU2NjYyMzkzMTYyMzI2MjYzNjY2**********
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 2025-07-12T00:00:00Z
	StartOfCreateTime *string `json:"StartOfCreateTime,omitempty" xml:"StartOfCreateTime,omitempty"`
	// example:
	//
	// ******b4fb044839815d4f2cd8******
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// example:
	//
	// example-workflow-****
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s ListWorkflowTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTasksRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowTasksRequest) GetEndOfCreateTime() *string {
	return s.EndOfCreateTime
}

func (s *ListWorkflowTasksRequest) GetKeyText() *string {
	return s.KeyText
}

func (s *ListWorkflowTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowTasksRequest) GetStartOfCreateTime() *string {
	return s.StartOfCreateTime
}

func (s *ListWorkflowTasksRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowTasksRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListWorkflowTasksRequest) SetEndOfCreateTime(v string) *ListWorkflowTasksRequest {
	s.EndOfCreateTime = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetKeyText(v string) *ListWorkflowTasksRequest {
	s.KeyText = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetMaxResults(v int32) *ListWorkflowTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetNextToken(v string) *ListWorkflowTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetStartOfCreateTime(v string) *ListWorkflowTasksRequest {
	s.StartOfCreateTime = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetWorkflowId(v string) *ListWorkflowTasksRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetWorkflowName(v string) *ListWorkflowTasksRequest {
	s.WorkflowName = &v
	return s
}

func (s *ListWorkflowTasksRequest) Validate() error {
	return dara.Validate(s)
}
