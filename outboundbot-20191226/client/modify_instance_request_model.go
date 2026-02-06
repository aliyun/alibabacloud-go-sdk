// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumber(v []*string) *ModifyInstanceRequest
	GetCallingNumber() []*string
	SetInstanceDescription(v string) *ModifyInstanceRequest
	GetInstanceDescription() *string
	SetInstanceId(v string) *ModifyInstanceRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyInstanceRequest
	GetInstanceName() *string
	SetMaxConcurrentConversation(v int32) *ModifyInstanceRequest
	GetMaxConcurrentConversation() *int32
}

type ModifyInstanceRequest struct {
	// example:
	//
	// ["95187"]
	CallingNumber       []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	InstanceDescription *string   `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 90515b5-6115-4ccf-83e2-52d5bfaf2ddf
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	MaxConcurrentConversation *int32 `json:"MaxConcurrentConversation,omitempty" xml:"MaxConcurrentConversation,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *ModifyInstanceRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *ModifyInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceRequest) GetMaxConcurrentConversation() *int32 {
	return s.MaxConcurrentConversation
}

func (s *ModifyInstanceRequest) SetCallingNumber(v []*string) *ModifyInstanceRequest {
	s.CallingNumber = v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceDescription(v string) *ModifyInstanceRequest {
	s.InstanceDescription = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceName(v string) *ModifyInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceRequest) SetMaxConcurrentConversation(v int32) *ModifyInstanceRequest {
	s.MaxConcurrentConversation = &v
	return s
}

func (s *ModifyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
