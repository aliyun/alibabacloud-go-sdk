// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckout interface {
	dara.Model
	String() string
	GoString() string
	SetRef(v string) *Checkout
	GetRef() *string
	SetRemote(v string) *Checkout
	GetRemote() *string
}

type Checkout struct {
	// example:
	//
	// +001691d0768ca49e9550beeb59fbc163f33b7e88:refs/remotes/origin/master
	Ref *string `json:"ref,omitempty" xml:"ref,omitempty"`
	// example:
	//
	// https:/your_token/@github.com/buptwzj/test-initRepo4.git
	Remote *string `json:"remote,omitempty" xml:"remote,omitempty"`
}

func (s Checkout) String() string {
	return dara.Prettify(s)
}

func (s Checkout) GoString() string {
	return s.String()
}

func (s *Checkout) GetRef() *string {
	return s.Ref
}

func (s *Checkout) GetRemote() *string {
	return s.Remote
}

func (s *Checkout) SetRef(v string) *Checkout {
	s.Ref = &v
	return s
}

func (s *Checkout) SetRemote(v string) *Checkout {
	s.Remote = &v
	return s
}

func (s *Checkout) Validate() error {
	return dara.Validate(s)
}
