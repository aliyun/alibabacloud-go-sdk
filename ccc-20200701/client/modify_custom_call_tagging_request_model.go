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
	// example:
	//
	// ["TagA","TagB"]
	CallTagNameList *string `json:"CallTagNameList,omitempty" xml:"CallTagNameList,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
