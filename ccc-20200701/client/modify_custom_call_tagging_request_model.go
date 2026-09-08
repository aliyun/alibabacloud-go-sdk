// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomCallTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallTagNameList(v string) *ModifyCustomCallTaggingRequest
	GetCallTagNameList() *string
	SetDescription(v string) *ModifyCustomCallTaggingRequest
	GetDescription() *string
	SetInstanceId(v string) *ModifyCustomCallTaggingRequest
	GetInstanceId() *string
	SetNumber(v string) *ModifyCustomCallTaggingRequest
	GetNumber() *string
}

type ModifyCustomCallTaggingRequest struct {
	// A list of number tag names. You must provide the complete list of number tags to be modified, and ensure that these number tags have already been created.
	//
	// example:
	//
	// ["TagA","TagB"]
	CallTagNameList *string `json:"CallTagNameList,omitempty" xml:"CallTagNameList,omitempty"`
	// The new description for the inbound number mark. This parameter is optional. The default value is empty, which indicates that the description will not be modified.
	//
	// example:
	//
	// 王先生
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number associated with the inbound number mark. The system matches the inbound number mark to be modified based on this number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1312121****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s ModifyCustomCallTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomCallTaggingRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomCallTaggingRequest) GetCallTagNameList() *string {
	return s.CallTagNameList
}

func (s *ModifyCustomCallTaggingRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCustomCallTaggingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCustomCallTaggingRequest) GetNumber() *string {
	return s.Number
}

func (s *ModifyCustomCallTaggingRequest) SetCallTagNameList(v string) *ModifyCustomCallTaggingRequest {
	s.CallTagNameList = &v
	return s
}

func (s *ModifyCustomCallTaggingRequest) SetDescription(v string) *ModifyCustomCallTaggingRequest {
	s.Description = &v
	return s
}

func (s *ModifyCustomCallTaggingRequest) SetInstanceId(v string) *ModifyCustomCallTaggingRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCustomCallTaggingRequest) SetNumber(v string) *ModifyCustomCallTaggingRequest {
	s.Number = &v
	return s
}

func (s *ModifyCustomCallTaggingRequest) Validate() error {
	return dara.Validate(s)
}
