// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateContactRequest
	GetConfig() *string
	SetContactName(v string) *CreateContactRequest
	GetContactName() *string
	SetType(v string) *CreateContactRequest
	GetType() *string
}

type CreateContactRequest struct {
	// 渠道参数配置 JSON 字符串。IM 类型示例：{"channels":[{"channelType":"dingtalk","clientId":"xxx","clientSecret":"xxx","targetType":"group","targetId":"xxx","robotCode":"xxx"}]}
	//
	// This parameter is required.
	//
	// example:
	//
	// {"channels":[{"channelType":"dingtalk","clientId":"xxx","clientSecret":"xxx","targetType":"group","targetId":"xxx","robotCode":"xxx"}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 联系人名称（用户自定义，用于展示），同一用户下不可重名
	//
	// This parameter is required.
	//
	// example:
	//
	// 陈睿
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// 渠道大类，当前支持 IM
	//
	// This parameter is required.
	//
	// example:
	//
	// IM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateContactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContactRequest) GoString() string {
	return s.String()
}

func (s *CreateContactRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CreateContactRequest) GetType() *string {
	return s.Type
}

func (s *CreateContactRequest) SetConfig(v string) *CreateContactRequest {
	s.Config = &v
	return s
}

func (s *CreateContactRequest) SetContactName(v string) *CreateContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateContactRequest) SetType(v string) *CreateContactRequest {
	s.Type = &v
	return s
}

func (s *CreateContactRequest) Validate() error {
	return dara.Validate(s)
}
