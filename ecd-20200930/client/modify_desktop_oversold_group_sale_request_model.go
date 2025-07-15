// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopOversoldGroupSaleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrenceCount(v int32) *ModifyDesktopOversoldGroupSaleRequest
	GetConcurrenceCount() *int32
	SetOversoldGroupId(v string) *ModifyDesktopOversoldGroupSaleRequest
	GetOversoldGroupId() *string
	SetOversoldUserCount(v int32) *ModifyDesktopOversoldGroupSaleRequest
	GetOversoldUserCount() *int32
}

type ModifyDesktopOversoldGroupSaleRequest struct {
	ConcurrenceCount  *int32  `json:"ConcurrenceCount,omitempty" xml:"ConcurrenceCount,omitempty"`
	OversoldGroupId   *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	OversoldUserCount *int32  `json:"OversoldUserCount,omitempty" xml:"OversoldUserCount,omitempty"`
}

func (s ModifyDesktopOversoldGroupSaleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldGroupSaleRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldGroupSaleRequest) GetConcurrenceCount() *int32 {
	return s.ConcurrenceCount
}

func (s *ModifyDesktopOversoldGroupSaleRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *ModifyDesktopOversoldGroupSaleRequest) GetOversoldUserCount() *int32 {
	return s.OversoldUserCount
}

func (s *ModifyDesktopOversoldGroupSaleRequest) SetConcurrenceCount(v int32) *ModifyDesktopOversoldGroupSaleRequest {
	s.ConcurrenceCount = &v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleRequest) SetOversoldGroupId(v string) *ModifyDesktopOversoldGroupSaleRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleRequest) SetOversoldUserCount(v int32) *ModifyDesktopOversoldGroupSaleRequest {
	s.OversoldUserCount = &v
	return s
}

func (s *ModifyDesktopOversoldGroupSaleRequest) Validate() error {
	return dara.Validate(s)
}
