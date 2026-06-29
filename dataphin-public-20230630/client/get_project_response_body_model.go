// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProjectResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetProjectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetProjectResponseBody
	GetMessage() *string
	SetProjectInfo(v *GetProjectResponseBodyProjectInfo) *GetProjectResponseBody
	GetProjectInfo() *GetProjectResponseBodyProjectInfo
	SetRequestId(v string) *GetProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProjectResponseBody
	GetSuccess() *bool
}

type GetProjectResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The project details.
	ProjectInfo *GetProjectResponseBodyProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProjectResponseBody) GetProjectInfo() *GetProjectResponseBodyProjectInfo {
	return s.ProjectInfo
}

func (s *GetProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectResponseBody) SetCode(v string) *GetProjectResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectResponseBody) SetHttpStatusCode(v int32) *GetProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetProjectResponseBody) SetMessage(v string) *GetProjectResponseBody {
	s.Message = &v
	return s
}

func (s *GetProjectResponseBody) SetProjectInfo(v *GetProjectResponseBodyProjectInfo) *GetProjectResponseBody {
	s.ProjectInfo = v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) SetSuccess(v bool) *GetProjectResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectResponseBody) Validate() error {
	if s.ProjectInfo != nil {
		if err := s.ProjectInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectResponseBodyProjectInfo struct {
	// The dependent data module display name.
	//
	// example:
	//
	// 业务板块显示名
	BizUnitDisplayName *string `json:"BizUnitDisplayName,omitempty" xml:"BizUnitDisplayName,omitempty"`
	// The dependent data module ID.
	//
	// example:
	//
	// 业务板块Id
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The compute source ID.
	//
	// example:
	//
	// 项目依赖计算源Id
	ComputeSourceId *int64 `json:"ComputeSourceId,omitempty" xml:"ComputeSourceId,omitempty"`
	// The compute source name.
	//
	// example:
	//
	// 项目依赖计算源名
	ComputeSourceName *string `json:"ComputeSourceName,omitempty" xml:"ComputeSourceName,omitempty"`
	// The project description.
	//
	// example:
	//
	// project for xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project display name.
	//
	// example:
	//
	// 项目显示名
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environment.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1703048484000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1703048484000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 12356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The project mode. Valid values:
	//
	// - BASIC
	//
	// - DEV_PROD
	//
	// example:
	//
	// BASIC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The project name.
	//
	// example:
	//
	// dp_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The project tag.
	//
	// example:
	//
	// GENERAL
	NameSpaceTag *string `json:"NameSpaceTag,omitempty" xml:"NameSpaceTag,omitempty"`
	// The project owner ID.
	//
	// example:
	//
	// 101111
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The project owner.
	//
	// example:
	//
	// 测试用户
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The stream compute source ID.
	//
	// example:
	//
	// 项目依赖实时计算源Id
	StreamComputeSourceId *int64 `json:"StreamComputeSourceId,omitempty" xml:"StreamComputeSourceId,omitempty"`
	// The stream compute source name.
	//
	// example:
	//
	// 项目依赖实时计算源名
	StreamComputeSourceName *string `json:"StreamComputeSourceName,omitempty" xml:"StreamComputeSourceName,omitempty"`
	// The project type.
	//
	// example:
	//
	// GENERAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The whitelists.
	WhiteLists []*GetProjectResponseBodyProjectInfoWhiteLists `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty" type:"Repeated"`
}

func (s GetProjectResponseBodyProjectInfo) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyProjectInfo) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProjectInfo) GetBizUnitDisplayName() *string {
	return s.BizUnitDisplayName
}

func (s *GetProjectResponseBodyProjectInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetProjectResponseBodyProjectInfo) GetComputeSourceId() *int64 {
	return s.ComputeSourceId
}

func (s *GetProjectResponseBodyProjectInfo) GetComputeSourceName() *string {
	return s.ComputeSourceName
}

func (s *GetProjectResponseBodyProjectInfo) GetDescription() *string {
	return s.Description
}

func (s *GetProjectResponseBodyProjectInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetProjectResponseBodyProjectInfo) GetEnv() *string {
	return s.Env
}

func (s *GetProjectResponseBodyProjectInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetProjectResponseBodyProjectInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetProjectResponseBodyProjectInfo) GetId() *int64 {
	return s.Id
}

func (s *GetProjectResponseBodyProjectInfo) GetMode() *string {
	return s.Mode
}

func (s *GetProjectResponseBodyProjectInfo) GetName() *string {
	return s.Name
}

func (s *GetProjectResponseBodyProjectInfo) GetNameSpaceTag() *string {
	return s.NameSpaceTag
}

func (s *GetProjectResponseBodyProjectInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetProjectResponseBodyProjectInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetProjectResponseBodyProjectInfo) GetStreamComputeSourceId() *int64 {
	return s.StreamComputeSourceId
}

func (s *GetProjectResponseBodyProjectInfo) GetStreamComputeSourceName() *string {
	return s.StreamComputeSourceName
}

func (s *GetProjectResponseBodyProjectInfo) GetType() *string {
	return s.Type
}

func (s *GetProjectResponseBodyProjectInfo) GetWhiteLists() []*GetProjectResponseBodyProjectInfoWhiteLists {
	return s.WhiteLists
}

func (s *GetProjectResponseBodyProjectInfo) SetBizUnitDisplayName(v string) *GetProjectResponseBodyProjectInfo {
	s.BizUnitDisplayName = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetBizUnitId(v int64) *GetProjectResponseBodyProjectInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetComputeSourceId(v int64) *GetProjectResponseBodyProjectInfo {
	s.ComputeSourceId = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetComputeSourceName(v string) *GetProjectResponseBodyProjectInfo {
	s.ComputeSourceName = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetDescription(v string) *GetProjectResponseBodyProjectInfo {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetDisplayName(v string) *GetProjectResponseBodyProjectInfo {
	s.DisplayName = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetEnv(v string) *GetProjectResponseBodyProjectInfo {
	s.Env = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetGmtCreate(v string) *GetProjectResponseBodyProjectInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetGmtModified(v string) *GetProjectResponseBodyProjectInfo {
	s.GmtModified = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetId(v int64) *GetProjectResponseBodyProjectInfo {
	s.Id = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetMode(v string) *GetProjectResponseBodyProjectInfo {
	s.Mode = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetName(v string) *GetProjectResponseBodyProjectInfo {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetNameSpaceTag(v string) *GetProjectResponseBodyProjectInfo {
	s.NameSpaceTag = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetOwner(v string) *GetProjectResponseBodyProjectInfo {
	s.Owner = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetOwnerName(v string) *GetProjectResponseBodyProjectInfo {
	s.OwnerName = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetStreamComputeSourceId(v int64) *GetProjectResponseBodyProjectInfo {
	s.StreamComputeSourceId = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetStreamComputeSourceName(v string) *GetProjectResponseBodyProjectInfo {
	s.StreamComputeSourceName = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetType(v string) *GetProjectResponseBodyProjectInfo {
	s.Type = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) SetWhiteLists(v []*GetProjectResponseBodyProjectInfoWhiteLists) *GetProjectResponseBodyProjectInfo {
	s.WhiteLists = v
	return s
}

func (s *GetProjectResponseBodyProjectInfo) Validate() error {
	if s.WhiteLists != nil {
		for _, item := range s.WhiteLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProjectResponseBodyProjectInfoWhiteLists struct {
	// The whitelist description.
	//
	// example:
	//
	// whitelist for xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 10.209.47.198
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetProjectResponseBodyProjectInfoWhiteLists) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyProjectInfoWhiteLists) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProjectInfoWhiteLists) GetDescription() *string {
	return s.Description
}

func (s *GetProjectResponseBodyProjectInfoWhiteLists) GetIp() *string {
	return s.Ip
}

func (s *GetProjectResponseBodyProjectInfoWhiteLists) GetPort() *string {
	return s.Port
}

func (s *GetProjectResponseBodyProjectInfoWhiteLists) SetDescription(v string) *GetProjectResponseBodyProjectInfoWhiteLists {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfoWhiteLists) SetIp(v string) *GetProjectResponseBodyProjectInfoWhiteLists {
	s.Ip = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfoWhiteLists) SetPort(v string) *GetProjectResponseBodyProjectInfoWhiteLists {
	s.Port = &v
	return s
}

func (s *GetProjectResponseBodyProjectInfoWhiteLists) Validate() error {
	return dara.Validate(s)
}
