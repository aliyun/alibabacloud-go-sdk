// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProjectByNameResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetProjectByNameResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetProjectByNameResponseBody
	GetMessage() *string
	SetProjectInfo(v *GetProjectByNameResponseBodyProjectInfo) *GetProjectByNameResponseBody
	GetProjectInfo() *GetProjectByNameResponseBodyProjectInfo
	SetRequestId(v string) *GetProjectByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProjectByNameResponseBody
	GetSuccess() *bool
}

type GetProjectByNameResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message     *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	ProjectInfo *GetProjectByNameResponseBodyProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProjectByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectByNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProjectByNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetProjectByNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProjectByNameResponseBody) GetProjectInfo() *GetProjectByNameResponseBodyProjectInfo {
	return s.ProjectInfo
}

func (s *GetProjectByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectByNameResponseBody) SetCode(v string) *GetProjectByNameResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectByNameResponseBody) SetHttpStatusCode(v int32) *GetProjectByNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetProjectByNameResponseBody) SetMessage(v string) *GetProjectByNameResponseBody {
	s.Message = &v
	return s
}

func (s *GetProjectByNameResponseBody) SetProjectInfo(v *GetProjectByNameResponseBodyProjectInfo) *GetProjectByNameResponseBody {
	s.ProjectInfo = v
	return s
}

func (s *GetProjectByNameResponseBody) SetRequestId(v string) *GetProjectByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectByNameResponseBody) SetSuccess(v bool) *GetProjectByNameResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectByNameResponseBody) Validate() error {
	if s.ProjectInfo != nil {
		if err := s.ProjectInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectByNameResponseBodyProjectInfo struct {
	// example:
	//
	// 测试
	BizUnitDisplayName *string `json:"BizUnitDisplayName,omitempty" xml:"BizUnitDisplayName,omitempty"`
	// example:
	//
	// 101131
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// 101711
	ComputeSourceId *int64 `json:"ComputeSourceId,omitempty" xml:"ComputeSourceId,omitempty"`
	// example:
	//
	// ds1
	ComputeSourceName *string `json:"ComputeSourceName,omitempty" xml:"ComputeSourceName,omitempty"`
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// xx test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 2025-06-10 10:01:01
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-06-10 10:01:01
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1030111021
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BASIC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BASE
	NameSpaceTag *string `json:"NameSpaceTag,omitempty" xml:"NameSpaceTag,omitempty"`
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 201711
	StreamComputeSourceId *int64 `json:"StreamComputeSourceId,omitempty" xml:"StreamComputeSourceId,omitempty"`
	// example:
	//
	// ds2
	StreamComputeSourceName *string `json:"StreamComputeSourceName,omitempty" xml:"StreamComputeSourceName,omitempty"`
	// example:
	//
	// GENERAL
	Type       *string                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	WhiteLists []*GetProjectByNameResponseBodyProjectInfoWhiteLists `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty" type:"Repeated"`
}

func (s GetProjectByNameResponseBodyProjectInfo) String() string {
	return dara.Prettify(s)
}

func (s GetProjectByNameResponseBodyProjectInfo) GoString() string {
	return s.String()
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetBizUnitDisplayName() *string {
	return s.BizUnitDisplayName
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetComputeSourceId() *int64 {
	return s.ComputeSourceId
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetComputeSourceName() *string {
	return s.ComputeSourceName
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetDescription() *string {
	return s.Description
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetEnv() *string {
	return s.Env
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetId() *int64 {
	return s.Id
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetMode() *string {
	return s.Mode
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetName() *string {
	return s.Name
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetNameSpaceTag() *string {
	return s.NameSpaceTag
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetStreamComputeSourceId() *int64 {
	return s.StreamComputeSourceId
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetStreamComputeSourceName() *string {
	return s.StreamComputeSourceName
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetType() *string {
	return s.Type
}

func (s *GetProjectByNameResponseBodyProjectInfo) GetWhiteLists() []*GetProjectByNameResponseBodyProjectInfoWhiteLists {
	return s.WhiteLists
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetBizUnitDisplayName(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.BizUnitDisplayName = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetBizUnitId(v int64) *GetProjectByNameResponseBodyProjectInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetComputeSourceId(v int64) *GetProjectByNameResponseBodyProjectInfo {
	s.ComputeSourceId = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetComputeSourceName(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.ComputeSourceName = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetDescription(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.Description = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetDisplayName(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.DisplayName = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetEnv(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.Env = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetGmtCreate(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetGmtModified(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.GmtModified = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetId(v int64) *GetProjectByNameResponseBodyProjectInfo {
	s.Id = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetMode(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.Mode = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetName(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.Name = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetNameSpaceTag(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.NameSpaceTag = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetOwner(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.Owner = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetOwnerName(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.OwnerName = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetStreamComputeSourceId(v int64) *GetProjectByNameResponseBodyProjectInfo {
	s.StreamComputeSourceId = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetStreamComputeSourceName(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.StreamComputeSourceName = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetType(v string) *GetProjectByNameResponseBodyProjectInfo {
	s.Type = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) SetWhiteLists(v []*GetProjectByNameResponseBodyProjectInfoWhiteLists) *GetProjectByNameResponseBodyProjectInfo {
	s.WhiteLists = v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfo) Validate() error {
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

type GetProjectByNameResponseBodyProjectInfoWhiteLists struct {
	// example:
	//
	// xx 白名单
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Ip
	//
	// example:
	//
	// 10.11.1.21
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetProjectByNameResponseBodyProjectInfoWhiteLists) String() string {
	return dara.Prettify(s)
}

func (s GetProjectByNameResponseBodyProjectInfoWhiteLists) GoString() string {
	return s.String()
}

func (s *GetProjectByNameResponseBodyProjectInfoWhiteLists) GetDescription() *string {
	return s.Description
}

func (s *GetProjectByNameResponseBodyProjectInfoWhiteLists) GetIp() *string {
	return s.Ip
}

func (s *GetProjectByNameResponseBodyProjectInfoWhiteLists) GetPort() *string {
	return s.Port
}

func (s *GetProjectByNameResponseBodyProjectInfoWhiteLists) SetDescription(v string) *GetProjectByNameResponseBodyProjectInfoWhiteLists {
	s.Description = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfoWhiteLists) SetIp(v string) *GetProjectByNameResponseBodyProjectInfoWhiteLists {
	s.Ip = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfoWhiteLists) SetPort(v string) *GetProjectByNameResponseBodyProjectInfoWhiteLists {
	s.Port = &v
	return s
}

func (s *GetProjectByNameResponseBodyProjectInfoWhiteLists) Validate() error {
	return dara.Validate(s)
}
