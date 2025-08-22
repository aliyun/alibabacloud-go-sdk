// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnWafGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateDcdnWafGroupRequest
	GetName() *string
	SetSubscribe(v string) *CreateDcdnWafGroupRequest
	GetSubscribe() *string
	SetTemplateId(v int64) *CreateDcdnWafGroupRequest
	GetTemplateId() *int64
}

type CreateDcdnWafGroupRequest struct {
	// The name of the WAF rule group. The name can be up to 128 characters in length. This parameter is required when you create a custom WAF rule group.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable subscription. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// When you replicate a custom rule group, do not specify this parameter.
	//
	// example:
	//
	// on
	Subscribe *string `json:"Subscribe,omitempty" xml:"Subscribe,omitempty"`
	// The ID of the rule group to be replicated. This parameter is required when you replicate a custom WAF rule group. You can call the [DescribeDcdnWafGroups](~~DescribeDcdnWafGroups~~) operation to query the ID of the rule group. If no template is used, set the value to 0 or do not specify this parameter.
	//
	// example:
	//
	// 0
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDcdnWafGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnWafGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDcdnWafGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateDcdnWafGroupRequest) GetSubscribe() *string {
	return s.Subscribe
}

func (s *CreateDcdnWafGroupRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateDcdnWafGroupRequest) SetName(v string) *CreateDcdnWafGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateDcdnWafGroupRequest) SetSubscribe(v string) *CreateDcdnWafGroupRequest {
	s.Subscribe = &v
	return s
}

func (s *CreateDcdnWafGroupRequest) SetTemplateId(v int64) *CreateDcdnWafGroupRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateDcdnWafGroupRequest) Validate() error {
	return dara.Validate(s)
}
