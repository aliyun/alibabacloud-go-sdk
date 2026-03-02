// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceGroupForServiceCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *OpenServiceGroupForServiceCmd
	GetCompanyId() *int64
	SetLaneId(v int64) *OpenServiceGroupForServiceCmd
	GetLaneId() *int64
	SetServiceIds(v []*int64) *OpenServiceGroupForServiceCmd
	GetServiceIds() []*int64
}

type OpenServiceGroupForServiceCmd struct {
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LaneId     *int64   `json:"laneId,omitempty" xml:"laneId,omitempty"`
	ServiceIds []*int64 `json:"serviceIds,omitempty" xml:"serviceIds,omitempty" type:"Repeated"`
}

func (s OpenServiceGroupForServiceCmd) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceGroupForServiceCmd) GoString() string {
	return s.String()
}

func (s *OpenServiceGroupForServiceCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *OpenServiceGroupForServiceCmd) GetLaneId() *int64 {
	return s.LaneId
}

func (s *OpenServiceGroupForServiceCmd) GetServiceIds() []*int64 {
	return s.ServiceIds
}

func (s *OpenServiceGroupForServiceCmd) SetCompanyId(v int64) *OpenServiceGroupForServiceCmd {
	s.CompanyId = &v
	return s
}

func (s *OpenServiceGroupForServiceCmd) SetLaneId(v int64) *OpenServiceGroupForServiceCmd {
	s.LaneId = &v
	return s
}

func (s *OpenServiceGroupForServiceCmd) SetServiceIds(v []*int64) *OpenServiceGroupForServiceCmd {
	s.ServiceIds = v
	return s
}

func (s *OpenServiceGroupForServiceCmd) Validate() error {
	return dara.Validate(s)
}
