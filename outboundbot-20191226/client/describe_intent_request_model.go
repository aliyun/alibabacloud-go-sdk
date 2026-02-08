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
	// instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// intent ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 10717802
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// aa279896-64a6-4182-864c-4f2b04ec8d17
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
