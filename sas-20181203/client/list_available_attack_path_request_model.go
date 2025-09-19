// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableAttackPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListAvailableAttackPathRequest
	GetLang() *string
}

type ListAvailableAttackPathRequest struct {
	// The language type for requests and responses. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListAvailableAttackPathRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableAttackPathRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableAttackPathRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAvailableAttackPathRequest) SetLang(v string) *ListAvailableAttackPathRequest {
	s.Lang = &v
	return s
}

func (s *ListAvailableAttackPathRequest) Validate() error {
	return dara.Validate(s)
}
