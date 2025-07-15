// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInvocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *StopInvocationRequest
	GetDesktopId() []*string
	SetInvokeId(v string) *StopInvocationRequest
	GetInvokeId() *string
	SetRegionId(v string) *StopInvocationRequest
	GetRegionId() *string
}

type StopInvocationRequest struct {
	// The ID of cloud desktop N. Valid values of N: 1 to 50.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-7d2a745b412b4601b2d47f6a768d****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInvocationRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *StopInvocationRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *StopInvocationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopInvocationRequest) SetDesktopId(v []*string) *StopInvocationRequest {
	s.DesktopId = v
	return s
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetRegionId(v string) *StopInvocationRequest {
	s.RegionId = &v
	return s
}

func (s *StopInvocationRequest) Validate() error {
	return dara.Validate(s)
}
