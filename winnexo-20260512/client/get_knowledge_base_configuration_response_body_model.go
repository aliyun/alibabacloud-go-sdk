// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeRecordEnabled(v bool) *GetKnowledgeBaseConfigurationResponseBody
	GetChangeRecordEnabled() *bool
	SetCode(v string) *GetKnowledgeBaseConfigurationResponseBody
	GetCode() *string
	SetDirectoryId(v string) *GetKnowledgeBaseConfigurationResponseBody
	GetDirectoryId() *string
	SetMessage(v string) *GetKnowledgeBaseConfigurationResponseBody
	GetMessage() *string
	SetOoVisibilityMode(v string) *GetKnowledgeBaseConfigurationResponseBody
	GetOoVisibilityMode() *string
	SetOperatingObjectNames(v []*string) *GetKnowledgeBaseConfigurationResponseBody
	GetOperatingObjectNames() []*string
	SetRequestId(v string) *GetKnowledgeBaseConfigurationResponseBody
	GetRequestId() *string
}

type GetKnowledgeBaseConfigurationResponseBody struct {
	// 是否开启后续内容更新的 AI 版本摘要
	ChangeRecordEnabled *bool `json:"changeRecordEnabled,omitempty" xml:"changeRecordEnabled,omitempty"`
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 企业知识库根目录 ID
	//
	// example:
	//
	// kb-example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 错误描述，成功时为 success
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 数字员工可见模式；历史未显式设置 ALL 的配置按 SPECIFIED 返回
	//
	// example:
	//
	// SPECIFIED
	OoVisibilityMode *string `json:"ooVisibilityMode,omitempty" xml:"ooVisibilityMode,omitempty"`
	// 显式关联的数字员工名称；ALL 返回空数组，SPECIFIED 空数组表示没有关联员工
	OperatingObjectNames []*string `json:"operatingObjectNames,omitempty" xml:"operatingObjectNames,omitempty" type:"Repeated"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetKnowledgeBaseConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseConfigurationResponseBody) GetChangeRecordEnabled() *bool {
	return s.ChangeRecordEnabled
}

func (s *GetKnowledgeBaseConfigurationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeBaseConfigurationResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetKnowledgeBaseConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeBaseConfigurationResponseBody) GetOoVisibilityMode() *string {
	return s.OoVisibilityMode
}

func (s *GetKnowledgeBaseConfigurationResponseBody) GetOperatingObjectNames() []*string {
	return s.OperatingObjectNames
}

func (s *GetKnowledgeBaseConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseConfigurationResponseBody) SetChangeRecordEnabled(v bool) *GetKnowledgeBaseConfigurationResponseBody {
	s.ChangeRecordEnabled = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponseBody) SetCode(v string) *GetKnowledgeBaseConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponseBody) SetDirectoryId(v string) *GetKnowledgeBaseConfigurationResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponseBody) SetMessage(v string) *GetKnowledgeBaseConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponseBody) SetOoVisibilityMode(v string) *GetKnowledgeBaseConfigurationResponseBody {
	s.OoVisibilityMode = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponseBody) SetOperatingObjectNames(v []*string) *GetKnowledgeBaseConfigurationResponseBody {
	s.OperatingObjectNames = v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponseBody) SetRequestId(v string) *GetKnowledgeBaseConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
