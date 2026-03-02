// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpConfigUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PdpConfigUpdateCmd
	GetAccountId() *string
	SetContext(v string) *PdpConfigUpdateCmd
	GetContext() *string
	SetId(v int64) *PdpConfigUpdateCmd
	GetId() *int64
}

type PdpConfigUpdateCmd struct {
	// example:
	//
	// 1223435435
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"key":"value"}
	Context *string `json:"context,omitempty" xml:"context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s PdpConfigUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpConfigUpdateCmd) GoString() string {
	return s.String()
}

func (s *PdpConfigUpdateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *PdpConfigUpdateCmd) GetContext() *string {
	return s.Context
}

func (s *PdpConfigUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *PdpConfigUpdateCmd) SetAccountId(v string) *PdpConfigUpdateCmd {
	s.AccountId = &v
	return s
}

func (s *PdpConfigUpdateCmd) SetContext(v string) *PdpConfigUpdateCmd {
	s.Context = &v
	return s
}

func (s *PdpConfigUpdateCmd) SetId(v int64) *PdpConfigUpdateCmd {
	s.Id = &v
	return s
}

func (s *PdpConfigUpdateCmd) Validate() error {
	return dara.Validate(s)
}
