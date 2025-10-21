// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWordGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *GetWordGroupRequest
	GetGroupId() *int64
	SetRegionId(v string) *GetWordGroupRequest
	GetRegionId() *string
}

type GetWordGroupRequest struct {
	// Keyword group ID.
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetWordGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWordGroupRequest) GoString() string {
	return s.String()
}

func (s *GetWordGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetWordGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWordGroupRequest) SetGroupId(v int64) *GetWordGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetWordGroupRequest) SetRegionId(v string) *GetWordGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetWordGroupRequest) Validate() error {
	return dara.Validate(s)
}
