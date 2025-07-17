// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthorityTemplateItemsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemsShrink(v string) *AddAuthorityTemplateItemsShrinkRequest
	GetItemsShrink() *string
	SetTemplateId(v int64) *AddAuthorityTemplateItemsShrinkRequest
	GetTemplateId() *int64
	SetTid(v int64) *AddAuthorityTemplateItemsShrinkRequest
	GetTid() *int64
}

type AddAuthorityTemplateItemsShrinkRequest struct {
	// This parameter is required.
	ItemsShrink *string `json:"Items,omitempty" xml:"Items,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15***
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s AddAuthorityTemplateItemsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAuthorityTemplateItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddAuthorityTemplateItemsShrinkRequest) GetItemsShrink() *string {
	return s.ItemsShrink
}

func (s *AddAuthorityTemplateItemsShrinkRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *AddAuthorityTemplateItemsShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddAuthorityTemplateItemsShrinkRequest) SetItemsShrink(v string) *AddAuthorityTemplateItemsShrinkRequest {
	s.ItemsShrink = &v
	return s
}

func (s *AddAuthorityTemplateItemsShrinkRequest) SetTemplateId(v int64) *AddAuthorityTemplateItemsShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *AddAuthorityTemplateItemsShrinkRequest) SetTid(v int64) *AddAuthorityTemplateItemsShrinkRequest {
	s.Tid = &v
	return s
}

func (s *AddAuthorityTemplateItemsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
