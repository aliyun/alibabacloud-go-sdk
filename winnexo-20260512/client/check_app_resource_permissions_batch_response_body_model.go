// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppResourcePermissionsBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckAppResourcePermissionsBatchResponseBody
	GetCode() *string
	SetMessage(v string) *CheckAppResourcePermissionsBatchResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckAppResourcePermissionsBatchResponseBody
	GetRequestId() *string
	SetResults(v []*CheckAppResourcePermissionsBatchResponseBodyResults) *CheckAppResourcePermissionsBatchResponseBody
	GetResults() []*CheckAppResourcePermissionsBatchResponseBodyResults
}

type CheckAppResourcePermissionsBatchResponseBody struct {
	// 业务状态码；成功为 200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述；成功时为空
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 与请求顺序一致的结果
	Results []*CheckAppResourcePermissionsBatchResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s CheckAppResourcePermissionsBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionsBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionsBatchResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckAppResourcePermissionsBatchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAppResourcePermissionsBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAppResourcePermissionsBatchResponseBody) GetResults() []*CheckAppResourcePermissionsBatchResponseBodyResults {
	return s.Results
}

func (s *CheckAppResourcePermissionsBatchResponseBody) SetCode(v string) *CheckAppResourcePermissionsBatchResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBody) SetMessage(v string) *CheckAppResourcePermissionsBatchResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBody) SetRequestId(v string) *CheckAppResourcePermissionsBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBody) SetResults(v []*CheckAppResourcePermissionsBatchResponseBodyResults) *CheckAppResourcePermissionsBatchResponseBody {
	s.Results = v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckAppResourcePermissionsBatchResponseBodyResults struct {
	// 是否有权限
	//
	// example:
	//
	// true
	Allowed *bool `json:"allowed,omitempty" xml:"allowed,omitempty"`
	// 目标权限点
	//
	// example:
	//
	// APP/VIEW
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// 首段为 APP 应用键的点分资源键
	//
	// example:
	//
	// app.demo.resource_1
	ResourceKey *string `json:"resourceKey,omitempty" xml:"resourceKey,omitempty"`
	// 固定为 APP_RESOURCE
	//
	// example:
	//
	// APP_RESOURCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CheckAppResourcePermissionsBatchResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s CheckAppResourcePermissionsBatchResponseBodyResults) GoString() string {
	return s.String()
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) GetAllowed() *bool {
	return s.Allowed
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) GetPermission() *string {
	return s.Permission
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) GetResourceType() *string {
	return s.ResourceType
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) SetAllowed(v bool) *CheckAppResourcePermissionsBatchResponseBodyResults {
	s.Allowed = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) SetPermission(v string) *CheckAppResourcePermissionsBatchResponseBodyResults {
	s.Permission = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) SetResourceKey(v string) *CheckAppResourcePermissionsBatchResponseBodyResults {
	s.ResourceKey = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) SetResourceType(v string) *CheckAppResourcePermissionsBatchResponseBodyResults {
	s.ResourceType = &v
	return s
}

func (s *CheckAppResourcePermissionsBatchResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
