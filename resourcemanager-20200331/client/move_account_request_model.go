// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *MoveAccountRequest
	GetAccountId() *string
	SetDestinationFolderId(v string) *MoveAccountRequest
	GetDestinationFolderId() *string
}

type MoveAccountRequest struct {
	// The ID of the account you want to move.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the destination folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-bVaRIG****
	DestinationFolderId *string `json:"DestinationFolderId,omitempty" xml:"DestinationFolderId,omitempty"`
}

func (s MoveAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveAccountRequest) GoString() string {
	return s.String()
}

func (s *MoveAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *MoveAccountRequest) GetDestinationFolderId() *string {
	return s.DestinationFolderId
}

func (s *MoveAccountRequest) SetAccountId(v string) *MoveAccountRequest {
	s.AccountId = &v
	return s
}

func (s *MoveAccountRequest) SetDestinationFolderId(v string) *MoveAccountRequest {
	s.DestinationFolderId = &v
	return s
}

func (s *MoveAccountRequest) Validate() error {
	return dara.Validate(s)
}
