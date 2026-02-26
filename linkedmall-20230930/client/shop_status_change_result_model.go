// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShopStatusChangeResult interface {
	dara.Model
	String() string
	GoString() string
	SetOperate(v bool) *ShopStatusChangeResult
	GetOperate() *bool
}

type ShopStatusChangeResult struct {
	// example:
	//
	// true
	Operate *bool `json:"operate,omitempty" xml:"operate,omitempty"`
}

func (s ShopStatusChangeResult) String() string {
	return dara.Prettify(s)
}

func (s ShopStatusChangeResult) GoString() string {
	return s.String()
}

func (s *ShopStatusChangeResult) GetOperate() *bool {
	return s.Operate
}

func (s *ShopStatusChangeResult) SetOperate(v bool) *ShopStatusChangeResult {
	s.Operate = &v
	return s
}

func (s *ShopStatusChangeResult) Validate() error {
	return dara.Validate(s)
}
