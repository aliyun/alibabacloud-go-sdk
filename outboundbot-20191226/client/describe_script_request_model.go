// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeScriptRequest
	GetInstanceId() *string
	SetScriptId(v string) *DescribeScriptRequest
	GetScriptId() *string
}

type DescribeScriptRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 38d2e8ed-04e9-4dac-83b5-a8e57642ef13
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Script ID
	//
	// This parameter is required.
	//
	// example:
	//
	// bee7e5b1-5d9a-4389-aa7e-bbbee5353a16
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DescribeScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScriptRequest) GoString() string {
	return s.String()
}

func (s *DescribeScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeScriptRequest) SetInstanceId(v string) *DescribeScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeScriptRequest) SetScriptId(v string) *DescribeScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *DescribeScriptRequest) Validate() error {
	return dara.Validate(s)
}
