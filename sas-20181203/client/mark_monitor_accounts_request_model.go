// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkMonitorAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v string) *MarkMonitorAccountsRequest
	GetAccountIds() *string
}

type MarkMonitorAccountsRequest struct {
	// The IDs of the members.
	//
	// This parameter is required.
	//
	// example:
	//
	// 131331822340XXXX,140649175187XXXX
	AccountIds *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
}

func (s MarkMonitorAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s MarkMonitorAccountsRequest) GoString() string {
	return s.String()
}

func (s *MarkMonitorAccountsRequest) GetAccountIds() *string {
	return s.AccountIds
}

func (s *MarkMonitorAccountsRequest) SetAccountIds(v string) *MarkMonitorAccountsRequest {
	s.AccountIds = &v
	return s
}

func (s *MarkMonitorAccountsRequest) Validate() error {
	return dara.Validate(s)
}
