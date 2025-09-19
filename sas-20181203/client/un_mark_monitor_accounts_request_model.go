// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnMarkMonitorAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v string) *UnMarkMonitorAccountsRequest
	GetAccountIds() *string
}

type UnMarkMonitorAccountsRequest struct {
	// The IDs of the members.
	//
	// This parameter is required.
	//
	// example:
	//
	// 125267953644XXXX,125807832682XXXX
	AccountIds *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
}

func (s UnMarkMonitorAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s UnMarkMonitorAccountsRequest) GoString() string {
	return s.String()
}

func (s *UnMarkMonitorAccountsRequest) GetAccountIds() *string {
	return s.AccountIds
}

func (s *UnMarkMonitorAccountsRequest) SetAccountIds(v string) *UnMarkMonitorAccountsRequest {
	s.AccountIds = &v
	return s
}

func (s *UnMarkMonitorAccountsRequest) Validate() error {
	return dara.Validate(s)
}
