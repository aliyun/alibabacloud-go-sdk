// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatConfigurations(v []*ListChatConfigurationsResponseBodyChatConfigurations) *ListChatConfigurationsResponseBody
	GetChatConfigurations() []*ListChatConfigurationsResponseBodyChatConfigurations
	SetRequestId(v string) *ListChatConfigurationsResponseBody
	GetRequestId() *string
}

type ListChatConfigurationsResponseBody struct {
	// example:
	//
	// [{"Name": "test-wechat", "CreatedDate": "2025-12-10T10:09:32+00:00", "UpdatedDate": "2025-12-28T13:59:48+00:00", "RamRole": "OOSServiceRole", "Type": "WeChat", "Outputs": "{\\"URL\\": \\"https://1407907063606569.appflow.aliyunnes\\"}]}
	ChatConfigurations []*ListChatConfigurationsResponseBodyChatConfigurations `json:"ChatConfigurations,omitempty" xml:"ChatConfigurations,omitempty" type:"Repeated"`
	// example:
	//
	// 4DB0****1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListChatConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatConfigurationsResponseBody) GetChatConfigurations() []*ListChatConfigurationsResponseBodyChatConfigurations {
	return s.ChatConfigurations
}

func (s *ListChatConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatConfigurationsResponseBody) SetChatConfigurations(v []*ListChatConfigurationsResponseBodyChatConfigurations) *ListChatConfigurationsResponseBody {
	s.ChatConfigurations = v
	return s
}

func (s *ListChatConfigurationsResponseBody) SetRequestId(v string) *ListChatConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatConfigurationsResponseBody) Validate() error {
	if s.ChatConfigurations != nil {
		for _, item := range s.ChatConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChatConfigurationsResponseBodyChatConfigurations struct {
	// example:
	//
	// 2025-12-10T10:09:32+00:00
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// example:
	//
	// test-wechat
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"URL": "https://1407907063606569.appflow.aliyunnes", "AesKey": "UKAxyRabaJlaDxS3XZ6st18grPteemRtGeDVVgF"}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// example:
	//
	// WeChat
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2025-12-28T13:59:48+00:00
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListChatConfigurationsResponseBodyChatConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ListChatConfigurationsResponseBodyChatConfigurations) GoString() string {
	return s.String()
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) GetName() *string {
	return s.Name
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) GetOutputs() *string {
	return s.Outputs
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) GetRamRole() *string {
	return s.RamRole
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) GetType() *string {
	return s.Type
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) SetCreatedDate(v string) *ListChatConfigurationsResponseBodyChatConfigurations {
	s.CreatedDate = &v
	return s
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) SetName(v string) *ListChatConfigurationsResponseBodyChatConfigurations {
	s.Name = &v
	return s
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) SetOutputs(v string) *ListChatConfigurationsResponseBodyChatConfigurations {
	s.Outputs = &v
	return s
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) SetRamRole(v string) *ListChatConfigurationsResponseBodyChatConfigurations {
	s.RamRole = &v
	return s
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) SetType(v string) *ListChatConfigurationsResponseBodyChatConfigurations {
	s.Type = &v
	return s
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) SetUpdatedDate(v string) *ListChatConfigurationsResponseBodyChatConfigurations {
	s.UpdatedDate = &v
	return s
}

func (s *ListChatConfigurationsResponseBodyChatConfigurations) Validate() error {
	return dara.Validate(s)
}
