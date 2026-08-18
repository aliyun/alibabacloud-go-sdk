// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetMemberBalanceLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeType(v string) *ModelRouterGetMemberBalanceLogsRequest
	GetChangeType() *string
	SetPage(v int32) *ModelRouterGetMemberBalanceLogsRequest
	GetPage() *int32
	SetSize(v int32) *ModelRouterGetMemberBalanceLogsRequest
	GetSize() *int32
	SetSkipTotal(v bool) *ModelRouterGetMemberBalanceLogsRequest
	GetSkipTotal() *bool
}

type ModelRouterGetMemberBalanceLogsRequest struct {
	// The change type filter.
	//
	// example:
	//
	// recharge
	ChangeType *string `json:"changeType,omitempty" xml:"changeType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// Specifies whether to skip the total count calculation.
	//
	// example:
	//
	// false
	SkipTotal *bool `json:"skipTotal,omitempty" xml:"skipTotal,omitempty"`
}

func (s ModelRouterGetMemberBalanceLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberBalanceLogsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberBalanceLogsRequest) GetChangeType() *string {
	return s.ChangeType
}

func (s *ModelRouterGetMemberBalanceLogsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterGetMemberBalanceLogsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ModelRouterGetMemberBalanceLogsRequest) GetSkipTotal() *bool {
	return s.SkipTotal
}

func (s *ModelRouterGetMemberBalanceLogsRequest) SetChangeType(v string) *ModelRouterGetMemberBalanceLogsRequest {
	s.ChangeType = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsRequest) SetPage(v int32) *ModelRouterGetMemberBalanceLogsRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsRequest) SetSize(v int32) *ModelRouterGetMemberBalanceLogsRequest {
	s.Size = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsRequest) SetSkipTotal(v bool) *ModelRouterGetMemberBalanceLogsRequest {
	s.SkipTotal = &v
	return s
}

func (s *ModelRouterGetMemberBalanceLogsRequest) Validate() error {
	return dara.Validate(s)
}
