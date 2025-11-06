// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsCubeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMdsCubeTaskRequest
	GetAppId() *string
	SetGreyConfigInfo(v string) *CreateMdsCubeTaskRequest
	GetGreyConfigInfo() *string
	SetGreyEndtimeData(v string) *CreateMdsCubeTaskRequest
	GetGreyEndtimeData() *string
	SetGreyNum(v int32) *CreateMdsCubeTaskRequest
	GetGreyNum() *int32
	SetPublishMode(v int32) *CreateMdsCubeTaskRequest
	GetPublishMode() *int32
	SetPublishType(v int32) *CreateMdsCubeTaskRequest
	GetPublishType() *int32
	SetTaskDesc(v string) *CreateMdsCubeTaskRequest
	GetTaskDesc() *string
	SetTemplateResourceId(v int64) *CreateMdsCubeTaskRequest
	GetTemplateResourceId() *int64
	SetTenantId(v string) *CreateMdsCubeTaskRequest
	GetTenantId() *string
	SetWhitelistIds(v string) *CreateMdsCubeTaskRequest
	GetWhitelistIds() *string
	SetWorkspaceId(v string) *CreateMdsCubeTaskRequest
	GetWorkspaceId() *string
}

type CreateMdsCubeTaskRequest struct {
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// ""
	GreyConfigInfo *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	// example:
	//
	// ""
	GreyEndtimeData *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	// example:
	//
	// 1
	GreyNum *int32 `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	// example:
	//
	// 0
	PublishMode *int32 `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	// example:
	//
	// 3
	PublishType *int32 `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// example:
	//
	// task_test
	TaskDesc *string `json:"TaskDesc,omitempty" xml:"TaskDesc,omitempty"`
	// example:
	//
	// 1
	TemplateResourceId *int64 `json:"TemplateResourceId,omitempty" xml:"TemplateResourceId,omitempty"`
	// example:
	//
	// ZXCXMAHQ-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// test1,test2
	WhitelistIds *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMdsCubeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMdsCubeTaskRequest) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *CreateMdsCubeTaskRequest) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *CreateMdsCubeTaskRequest) GetGreyNum() *int32 {
	return s.GreyNum
}

func (s *CreateMdsCubeTaskRequest) GetPublishMode() *int32 {
	return s.PublishMode
}

func (s *CreateMdsCubeTaskRequest) GetPublishType() *int32 {
	return s.PublishType
}

func (s *CreateMdsCubeTaskRequest) GetTaskDesc() *string {
	return s.TaskDesc
}

func (s *CreateMdsCubeTaskRequest) GetTemplateResourceId() *int64 {
	return s.TemplateResourceId
}

func (s *CreateMdsCubeTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMdsCubeTaskRequest) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *CreateMdsCubeTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMdsCubeTaskRequest) SetAppId(v string) *CreateMdsCubeTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetGreyConfigInfo(v string) *CreateMdsCubeTaskRequest {
	s.GreyConfigInfo = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetGreyEndtimeData(v string) *CreateMdsCubeTaskRequest {
	s.GreyEndtimeData = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetGreyNum(v int32) *CreateMdsCubeTaskRequest {
	s.GreyNum = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetPublishMode(v int32) *CreateMdsCubeTaskRequest {
	s.PublishMode = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetPublishType(v int32) *CreateMdsCubeTaskRequest {
	s.PublishType = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetTaskDesc(v string) *CreateMdsCubeTaskRequest {
	s.TaskDesc = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetTemplateResourceId(v int64) *CreateMdsCubeTaskRequest {
	s.TemplateResourceId = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetTenantId(v string) *CreateMdsCubeTaskRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetWhitelistIds(v string) *CreateMdsCubeTaskRequest {
	s.WhitelistIds = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) SetWorkspaceId(v string) *CreateMdsCubeTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMdsCubeTaskRequest) Validate() error {
	return dara.Validate(s)
}
