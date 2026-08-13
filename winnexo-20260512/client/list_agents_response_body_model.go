// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentsResponseBody
	GetCode() *string
	SetItems(v []*ListAgentsResponseBodyItems) *ListAgentsResponseBody
	GetItems() []*ListAgentsResponseBodyItems
	SetMessage(v string) *ListAgentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAgentsResponseBody
	GetRequestId() *string
}

type ListAgentsResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 租户全量数字员工列表（含停用，按名称字母序）
	Items []*ListAgentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentsResponseBody) GetItems() []*ListAgentsResponseBodyItems {
	return s.Items
}

func (s *ListAgentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentsResponseBody) SetCode(v string) *ListAgentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentsResponseBody) SetItems(v []*ListAgentsResponseBodyItems) *ListAgentsResponseBody {
	s.Items = v
	return s
}

func (s *ListAgentsResponseBody) SetMessage(v string) *ListAgentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentsResponseBody) SetRequestId(v string) *ListAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentsResponseBodyItems struct {
	// 使用权限授权模式：SPECIFIED_USERS=指定用户 / ALL_USERS=所有用户；未设置时为 null
	//
	// example:
	//
	// string_value
	AuthMode *string `json:"authMode,omitempty" xml:"authMode,omitempty"`
	// 数字员工显示名称
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 启用/停用状态
	//
	// example:
	//
	// true
	IsActive *bool `json:"isActive,omitempty" xml:"isActive,omitempty"`
	// 数字员工名称（唯一标识）
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
}

func (s ListAgentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyItems) GetAuthMode() *string {
	return s.AuthMode
}

func (s *ListAgentsResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAgentsResponseBodyItems) GetIsActive() *bool {
	return s.IsActive
}

func (s *ListAgentsResponseBodyItems) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListAgentsResponseBodyItems) SetAuthMode(v string) *ListAgentsResponseBodyItems {
	s.AuthMode = &v
	return s
}

func (s *ListAgentsResponseBodyItems) SetDisplayName(v string) *ListAgentsResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListAgentsResponseBodyItems) SetIsActive(v bool) *ListAgentsResponseBodyItems {
	s.IsActive = &v
	return s
}

func (s *ListAgentsResponseBodyItems) SetOperatingObjectName(v string) *ListAgentsResponseBodyItems {
	s.OperatingObjectName = &v
	return s
}

func (s *ListAgentsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
