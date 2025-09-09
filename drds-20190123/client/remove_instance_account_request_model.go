// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *RemoveInstanceAccountRequest
	GetAccountName() *string
	SetDrdsInstanceId(v string) *RemoveInstanceAccountRequest
	GetDrdsInstanceId() *string
}

type RemoveInstanceAccountRequest struct {
	// The name of the member account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds***********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveInstanceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *RemoveInstanceAccountRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *RemoveInstanceAccountRequest) SetAccountName(v string) *RemoveInstanceAccountRequest {
	s.AccountName = &v
	return s
}

func (s *RemoveInstanceAccountRequest) SetDrdsInstanceId(v string) *RemoveInstanceAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveInstanceAccountRequest) Validate() error {
	return dara.Validate(s)
}
