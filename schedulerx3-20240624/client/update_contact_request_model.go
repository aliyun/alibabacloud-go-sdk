// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateContactRequest
	GetConfig() *string
	SetContactName(v string) *UpdateContactRequest
	GetContactName() *string
	SetEnabled(v bool) *UpdateContactRequest
	GetEnabled() *bool
	SetType(v string) *UpdateContactRequest
	GetType() *string
}

type UpdateContactRequest struct {
	// 渠道参数配置 JSON 字符串（可选，传入则更新）
	//
	// example:
	//
	// {"channels":[{"channelType":"dingtalk","clientId":"xxx","clientSecret":"xxx","targetType":"group","targetId":"xxx","robotCode":"xxx"}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Tom
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// 是否启用（true/false，可选）
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// 渠道大类（可选，传入则更新）
	//
	// example:
	//
	// IM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateContactRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *UpdateContactRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateContactRequest) GetType() *string {
	return s.Type
}

func (s *UpdateContactRequest) SetConfig(v string) *UpdateContactRequest {
	s.Config = &v
	return s
}

func (s *UpdateContactRequest) SetContactName(v string) *UpdateContactRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateContactRequest) SetEnabled(v bool) *UpdateContactRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateContactRequest) SetType(v string) *UpdateContactRequest {
	s.Type = &v
	return s
}

func (s *UpdateContactRequest) Validate() error {
	return dara.Validate(s)
}
