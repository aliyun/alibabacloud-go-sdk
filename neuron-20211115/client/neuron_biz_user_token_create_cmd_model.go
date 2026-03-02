// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronBizUserTokenCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *NeuronBizUserTokenCreateCmd
	GetAccountId() *int64
	SetBizId(v string) *NeuronBizUserTokenCreateCmd
	GetBizId() *string
	SetPermission(v string) *NeuronBizUserTokenCreateCmd
	GetPermission() *string
}

type NeuronBizUserTokenCreateCmd struct {
	AccountId *int64 `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 233131
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// example:
	//
	// ADMIN
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
}

func (s NeuronBizUserTokenCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronBizUserTokenCreateCmd) GoString() string {
	return s.String()
}

func (s *NeuronBizUserTokenCreateCmd) GetAccountId() *int64 {
	return s.AccountId
}

func (s *NeuronBizUserTokenCreateCmd) GetBizId() *string {
	return s.BizId
}

func (s *NeuronBizUserTokenCreateCmd) GetPermission() *string {
	return s.Permission
}

func (s *NeuronBizUserTokenCreateCmd) SetAccountId(v int64) *NeuronBizUserTokenCreateCmd {
	s.AccountId = &v
	return s
}

func (s *NeuronBizUserTokenCreateCmd) SetBizId(v string) *NeuronBizUserTokenCreateCmd {
	s.BizId = &v
	return s
}

func (s *NeuronBizUserTokenCreateCmd) SetPermission(v string) *NeuronBizUserTokenCreateCmd {
	s.Permission = &v
	return s
}

func (s *NeuronBizUserTokenCreateCmd) Validate() error {
	return dara.Validate(s)
}
