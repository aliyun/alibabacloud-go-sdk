// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMachinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostname(v string) *DescribeEdgeMachinesRequest
	GetHostname() *string
	SetLifeState(v string) *DescribeEdgeMachinesRequest
	GetLifeState() *string
	SetModel(v string) *DescribeEdgeMachinesRequest
	GetModel() *string
	SetOnlineState(v string) *DescribeEdgeMachinesRequest
	GetOnlineState() *string
	SetPageNumber(v int64) *DescribeEdgeMachinesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeEdgeMachinesRequest
	GetPageSize() *int64
}

type DescribeEdgeMachinesRequest struct {
	// The `hostname` of the cloud-native box.
	//
	// example:
	//
	// ack-v-b010-ssdfw****
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The lifecycle status.
	//
	// example:
	//
	// activated/waitOnline/deleting
	LifeState *string `json:"life_state,omitempty" xml:"life_state,omitempty"`
	// The type of cloud-native box.
	//
	// example:
	//
	// ACK-V-B010
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The status of the cloud-native box. Valid values:
	//
	// 	- `offline`
	//
	// 	- `online`
	//
	// example:
	//
	// offline/online
	OnlineState *string `json:"online_state,omitempty" xml:"online_state,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s DescribeEdgeMachinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachinesRequest) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeEdgeMachinesRequest) GetLifeState() *string {
	return s.LifeState
}

func (s *DescribeEdgeMachinesRequest) GetModel() *string {
	return s.Model
}

func (s *DescribeEdgeMachinesRequest) GetOnlineState() *string {
	return s.OnlineState
}

func (s *DescribeEdgeMachinesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEdgeMachinesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEdgeMachinesRequest) SetHostname(v string) *DescribeEdgeMachinesRequest {
	s.Hostname = &v
	return s
}

func (s *DescribeEdgeMachinesRequest) SetLifeState(v string) *DescribeEdgeMachinesRequest {
	s.LifeState = &v
	return s
}

func (s *DescribeEdgeMachinesRequest) SetModel(v string) *DescribeEdgeMachinesRequest {
	s.Model = &v
	return s
}

func (s *DescribeEdgeMachinesRequest) SetOnlineState(v string) *DescribeEdgeMachinesRequest {
	s.OnlineState = &v
	return s
}

func (s *DescribeEdgeMachinesRequest) SetPageNumber(v int64) *DescribeEdgeMachinesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEdgeMachinesRequest) SetPageSize(v int64) *DescribeEdgeMachinesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEdgeMachinesRequest) Validate() error {
	return dara.Validate(s)
}
