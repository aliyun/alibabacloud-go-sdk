// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppBizUser interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *NeuronAppBizUser
	GetAccountId() *int64
	SetBizId(v string) *NeuronAppBizUser
	GetBizId() *string
	SetPermission(v string) *NeuronAppBizUser
	GetPermission() *string
}

type NeuronAppBizUser struct {
	// example:
	//
	// 1235254534
	AccountId *int64 `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 233131
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// example:
	//
	// ADMIN
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
}

func (s NeuronAppBizUser) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppBizUser) GoString() string {
	return s.String()
}

func (s *NeuronAppBizUser) GetAccountId() *int64 {
	return s.AccountId
}

func (s *NeuronAppBizUser) GetBizId() *string {
	return s.BizId
}

func (s *NeuronAppBizUser) GetPermission() *string {
	return s.Permission
}

func (s *NeuronAppBizUser) SetAccountId(v int64) *NeuronAppBizUser {
	s.AccountId = &v
	return s
}

func (s *NeuronAppBizUser) SetBizId(v string) *NeuronAppBizUser {
	s.BizId = &v
	return s
}

func (s *NeuronAppBizUser) SetPermission(v string) *NeuronAppBizUser {
	s.Permission = &v
	return s
}

func (s *NeuronAppBizUser) Validate() error {
	return dara.Validate(s)
}
