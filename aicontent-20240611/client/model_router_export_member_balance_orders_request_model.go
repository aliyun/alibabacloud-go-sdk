// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterExportMemberBalanceOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *ModelRouterExportMemberBalanceOrdersRequest
	GetBalanceType() *string
	SetDirection(v string) *ModelRouterExportMemberBalanceOrdersRequest
	GetDirection() *string
}

type ModelRouterExportMemberBalanceOrdersRequest struct {
	// The balance type filter. Valid values: permanent and monthly.
	//
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The change direction filter. Valid values: in and out.
	//
	// example:
	//
	// in
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
}

func (s ModelRouterExportMemberBalanceOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterExportMemberBalanceOrdersRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterExportMemberBalanceOrdersRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterExportMemberBalanceOrdersRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModelRouterExportMemberBalanceOrdersRequest) SetBalanceType(v string) *ModelRouterExportMemberBalanceOrdersRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersRequest) SetDirection(v string) *ModelRouterExportMemberBalanceOrdersRequest {
	s.Direction = &v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersRequest) Validate() error {
	return dara.Validate(s)
}
