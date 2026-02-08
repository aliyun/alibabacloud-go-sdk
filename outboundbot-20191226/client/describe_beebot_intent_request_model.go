// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBeebotIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeBeebotIntentRequest
	GetInstanceId() *string
	SetIntentId(v int64) *DescribeBeebotIntentRequest
	GetIntentId() *int64
	SetScriptId(v string) *DescribeBeebotIntentRequest
	GetScriptId() *string
}

type DescribeBeebotIntentRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The intent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DescribeBeebotIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBeebotIntentRequest) GoString() string {
	return s.String()
}

func (s *DescribeBeebotIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBeebotIntentRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DescribeBeebotIntentRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeBeebotIntentRequest) SetInstanceId(v string) *DescribeBeebotIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBeebotIntentRequest) SetIntentId(v int64) *DescribeBeebotIntentRequest {
	s.IntentId = &v
	return s
}

func (s *DescribeBeebotIntentRequest) SetScriptId(v string) *DescribeBeebotIntentRequest {
	s.ScriptId = &v
	return s
}

func (s *DescribeBeebotIntentRequest) Validate() error {
	return dara.Validate(s)
}
