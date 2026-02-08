// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumbers(v []*string) *SuspendCallRequest
	GetCalledNumbers() []*string
	SetGroupId(v string) *SuspendCallRequest
	GetGroupId() *string
	SetInstanceId(v string) *SuspendCallRequest
	GetInstanceId() *string
}

type SuspendCallRequest struct {
	// List of called numbers
	//
	// example:
	//
	// []
	CalledNumbers []*string `json:"CalledNumbers,omitempty" xml:"CalledNumbers,omitempty" type:"Repeated"`
	// Task group ID (Required)
	//
	// example:
	//
	// f745881b-343d-43e4-9c51-31b7b063031c
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Instance ID (Required)
	//
	// This parameter is required.
	//
	// example:
	//
	// a5fc6490-ef1e-4666-870a-07a4e586c414
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SuspendCallRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendCallRequest) GoString() string {
	return s.String()
}

func (s *SuspendCallRequest) GetCalledNumbers() []*string {
	return s.CalledNumbers
}

func (s *SuspendCallRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SuspendCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SuspendCallRequest) SetCalledNumbers(v []*string) *SuspendCallRequest {
	s.CalledNumbers = v
	return s
}

func (s *SuspendCallRequest) SetGroupId(v string) *SuspendCallRequest {
	s.GroupId = &v
	return s
}

func (s *SuspendCallRequest) SetInstanceId(v string) *SuspendCallRequest {
	s.InstanceId = &v
	return s
}

func (s *SuspendCallRequest) Validate() error {
	return dara.Validate(s)
}
