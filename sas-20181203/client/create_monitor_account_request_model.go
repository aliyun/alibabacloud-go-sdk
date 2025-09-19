// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v string) *CreateMonitorAccountRequest
	GetAccountIds() *string
}

type CreateMonitorAccountRequest struct {
	// The account IDs of members in the resource directory.
	//
	// >  You can call the [ListAccountsInResourceDirectory](~~ListAccountsInResourceDirectory~~) operation to obtain the account IDs. Separate multiple account IDs with commas (,). If you specify a value for this parameter, the existing list of members is replaced by the new list that you specify. Otherwise, the existing list is cleared.
	//
	// example:
	//
	// 1026780160******,1457515594******
	AccountIds *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
}

func (s CreateMonitorAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorAccountRequest) GetAccountIds() *string {
	return s.AccountIds
}

func (s *CreateMonitorAccountRequest) SetAccountIds(v string) *CreateMonitorAccountRequest {
	s.AccountIds = &v
	return s
}

func (s *CreateMonitorAccountRequest) Validate() error {
	return dara.Validate(s)
}
