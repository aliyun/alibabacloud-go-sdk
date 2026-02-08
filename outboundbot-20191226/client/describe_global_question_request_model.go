// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalQuestionId(v string) *DescribeGlobalQuestionRequest
	GetGlobalQuestionId() *string
	SetInstanceId(v string) *DescribeGlobalQuestionRequest
	GetInstanceId() *string
	SetScriptId(v string) *DescribeGlobalQuestionRequest
	GetScriptId() *string
}

type DescribeGlobalQuestionRequest struct {
	// Global question ID
	//
	// This parameter is required.
	//
	// example:
	//
	// e851e242-ad67-4507-96a2-d4114564dcec
	GlobalQuestionId *string `json:"GlobalQuestionId,omitempty" xml:"GlobalQuestionId,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 7cefbff0-8d50-4d6f-b93c-73cee23c1555
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// b7ee988b-2837-4bc1-9d56-f76e7c831f60
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DescribeGlobalQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalQuestionRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalQuestionRequest) GetGlobalQuestionId() *string {
	return s.GlobalQuestionId
}

func (s *DescribeGlobalQuestionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGlobalQuestionRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeGlobalQuestionRequest) SetGlobalQuestionId(v string) *DescribeGlobalQuestionRequest {
	s.GlobalQuestionId = &v
	return s
}

func (s *DescribeGlobalQuestionRequest) SetInstanceId(v string) *DescribeGlobalQuestionRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGlobalQuestionRequest) SetScriptId(v string) *DescribeGlobalQuestionRequest {
	s.ScriptId = &v
	return s
}

func (s *DescribeGlobalQuestionRequest) Validate() error {
	return dara.Validate(s)
}
