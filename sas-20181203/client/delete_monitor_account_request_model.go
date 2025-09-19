// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeleteMonitorAccountRequest
	GetAccountId() *string
}

type DeleteMonitorAccountRequest struct {
	// The ID of the member that you want to delete.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the IDs of the members in the Security Center console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1840517068******
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteMonitorAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteMonitorAccountRequest) SetAccountId(v string) *DeleteMonitorAccountRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteMonitorAccountRequest) Validate() error {
	return dara.Validate(s)
}
