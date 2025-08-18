// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaitingRoomRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DeleteWaitingRoomRuleRequest
	GetSiteId() *int64
	SetWaitingRoomRuleId(v int64) *DeleteWaitingRoomRuleRequest
	GetWaitingRoomRuleId() *int64
}

type DeleteWaitingRoomRuleRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The ID of the waiting room bypass rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3672886****
	WaitingRoomRuleId *int64 `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s DeleteWaitingRoomRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaitingRoomRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteWaitingRoomRuleRequest) GetWaitingRoomRuleId() *int64 {
	return s.WaitingRoomRuleId
}

func (s *DeleteWaitingRoomRuleRequest) SetSiteId(v int64) *DeleteWaitingRoomRuleRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteWaitingRoomRuleRequest) SetWaitingRoomRuleId(v int64) *DeleteWaitingRoomRuleRequest {
	s.WaitingRoomRuleId = &v
	return s
}

func (s *DeleteWaitingRoomRuleRequest) Validate() error {
	return dara.Validate(s)
}
