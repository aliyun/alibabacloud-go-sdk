// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupName(v string) *PutContactGroupRequest
	GetContactGroupName() *string
	SetContactNames(v []*string) *PutContactGroupRequest
	GetContactNames() []*string
	SetDescribe(v string) *PutContactGroupRequest
	GetDescribe() *string
	SetEnableSubscribed(v bool) *PutContactGroupRequest
	GetEnableSubscribed() *bool
}

type PutContactGroupRequest struct {
	// The name of the alert contact group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Attendance system group
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The name of the alert contact.
	ContactNames []*string `json:"ContactNames,omitempty" xml:"ContactNames,omitempty" type:"Repeated"`
	// The description of the alert contact group.
	//
	// example:
	//
	// Alert test
	Describe *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// Specifies whether to enable the subscription feature. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	EnableSubscribed *bool `json:"EnableSubscribed,omitempty" xml:"EnableSubscribed,omitempty"`
}

func (s PutContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s PutContactGroupRequest) GoString() string {
	return s.String()
}

func (s *PutContactGroupRequest) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *PutContactGroupRequest) GetContactNames() []*string {
	return s.ContactNames
}

func (s *PutContactGroupRequest) GetDescribe() *string {
	return s.Describe
}

func (s *PutContactGroupRequest) GetEnableSubscribed() *bool {
	return s.EnableSubscribed
}

func (s *PutContactGroupRequest) SetContactGroupName(v string) *PutContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *PutContactGroupRequest) SetContactNames(v []*string) *PutContactGroupRequest {
	s.ContactNames = v
	return s
}

func (s *PutContactGroupRequest) SetDescribe(v string) *PutContactGroupRequest {
	s.Describe = &v
	return s
}

func (s *PutContactGroupRequest) SetEnableSubscribed(v bool) *PutContactGroupRequest {
	s.EnableSubscribed = &v
	return s
}

func (s *PutContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
