// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAttackEventDetailRequest
	GetId() *string
	SetLang(v string) *GetAttackEventDetailRequest
	GetLang() *string
}

type GetAttackEventDetailRequest struct {
	// The unique identifier ID for the alert event.
	//
	// example:
	//
	// 18825544674********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language type for requesting and receiving messages. Values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetAttackEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttackEventDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAttackEventDetailRequest) GetId() *string {
	return s.Id
}

func (s *GetAttackEventDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAttackEventDetailRequest) SetId(v string) *GetAttackEventDetailRequest {
	s.Id = &v
	return s
}

func (s *GetAttackEventDetailRequest) SetLang(v string) *GetAttackEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *GetAttackEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
