// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBlockStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *ModifyBlockStatusRequest
	GetDuration() *int32
	SetInstanceId(v string) *ModifyBlockStatusRequest
	GetInstanceId() *string
	SetLines(v []*string) *ModifyBlockStatusRequest
	GetLines() []*string
	SetStatus(v string) *ModifyBlockStatusRequest
	GetStatus() *string
}

type ModifyBlockStatusRequest struct {
	// The blocking period. Valid values: **15*	- to **43200**. Unit: minutes.
	//
	// > If you set **Status*	- to **do**, you must also specify this parameter.
	//
	// example:
	//
	// 60
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the Anti-DDoS Proxy (Chinese Mainland) instance to manage.
	//
	// >  You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all Anti-DDoS Proxy instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// An array consisting of the Internet service provider (ISP) lines from which traffic is blocked.
	//
	// This parameter is required.
	//
	// example:
	//
	// ct
	Lines []*string `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
	// Specifies the status of the Diversion from Origin Server policy. Valid values:
	//
	// 	- **do**: enables the policy.
	//
	// 	- **undo**: disables the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// do
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyBlockStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlockStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyBlockStatusRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ModifyBlockStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBlockStatusRequest) GetLines() []*string {
	return s.Lines
}

func (s *ModifyBlockStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyBlockStatusRequest) SetDuration(v int32) *ModifyBlockStatusRequest {
	s.Duration = &v
	return s
}

func (s *ModifyBlockStatusRequest) SetInstanceId(v string) *ModifyBlockStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBlockStatusRequest) SetLines(v []*string) *ModifyBlockStatusRequest {
	s.Lines = v
	return s
}

func (s *ModifyBlockStatusRequest) SetStatus(v string) *ModifyBlockStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyBlockStatusRequest) Validate() error {
	return dara.Validate(s)
}
