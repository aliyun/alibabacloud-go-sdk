// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeIntentRequest
	GetInstanceId() *string
	SetIntentId(v string) *DescribeIntentRequest
	GetIntentId() *string
	SetScriptId(v string) *DescribeIntentRequest
	GetScriptId() *string
}

type DescribeIntentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0f8a0059-dc9c-4151-8378-4734bbadf3cc
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b9ff4e88-65f9-4eb3-987c-11ba51f3f24d
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DescribeIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentRequest) GoString() string {
	return s.String()
}

func (s *DescribeIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIntentRequest) GetIntentId() *string {
	return s.IntentId
}

func (s *DescribeIntentRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeIntentRequest) SetInstanceId(v string) *DescribeIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntentRequest) SetIntentId(v string) *DescribeIntentRequest {
	s.IntentId = &v
	return s
}

func (s *DescribeIntentRequest) SetScriptId(v string) *DescribeIntentRequest {
	s.ScriptId = &v
	return s
}

func (s *DescribeIntentRequest) Validate() error {
	return dara.Validate(s)
}
