// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCreditLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreditLine(v string) *SetCreditLineRequest
	GetCreditLine() *string
	SetUid(v int64) *SetCreditLineRequest
	GetUid() *int64
}

type SetCreditLineRequest struct {
	// New Credit Line
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	CreditLine *string `json:"CreditLine,omitempty" xml:"CreditLine,omitempty"`
	// The UID of Sub Account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1263644979775567
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SetCreditLineRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCreditLineRequest) GoString() string {
	return s.String()
}

func (s *SetCreditLineRequest) GetCreditLine() *string {
	return s.CreditLine
}

func (s *SetCreditLineRequest) GetUid() *int64 {
	return s.Uid
}

func (s *SetCreditLineRequest) SetCreditLine(v string) *SetCreditLineRequest {
	s.CreditLine = &v
	return s
}

func (s *SetCreditLineRequest) SetUid(v int64) *SetCreditLineRequest {
	s.Uid = &v
	return s
}

func (s *SetCreditLineRequest) Validate() error {
	return dara.Validate(s)
}
