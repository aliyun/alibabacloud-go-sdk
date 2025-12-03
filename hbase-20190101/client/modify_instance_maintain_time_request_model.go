// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMaintainTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyInstanceMaintainTimeRequest
	GetClusterId() *string
	SetMaintainEndTime(v string) *ModifyInstanceMaintainTimeRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *ModifyInstanceMaintainTimeRequest
	GetMaintainStartTime() *string
}

type ModifyInstanceMaintainTimeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1b**6jco89****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 06:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
}

func (s ModifyInstanceMaintainTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyInstanceMaintainTimeRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *ModifyInstanceMaintainTimeRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *ModifyInstanceMaintainTimeRequest) SetClusterId(v string) *ModifyInstanceMaintainTimeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainEndTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainStartTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) Validate() error {
	return dara.Validate(s)
}
