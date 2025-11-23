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
	// The resources that you want to add to the permission template.
	//
	// This parameter is required.
	ItemsShrink *string `json:"Items,omitempty" xml:"Items,omitempty"`
	// The ID of the permission template. You can call the [CreateAuthorityTemplate](https://help.aliyun.com/document_detail/600705.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
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
