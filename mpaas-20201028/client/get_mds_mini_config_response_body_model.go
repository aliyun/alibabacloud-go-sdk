// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMdsMiniConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMdsMiniConfigResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetMdsMiniConfigResponseBody
	GetResultCode() *string
	SetResultContent(v *GetMdsMiniConfigResponseBodyResultContent) *GetMdsMiniConfigResponseBody
	GetResultContent() *GetMdsMiniConfigResponseBodyResultContent
	SetResultMessage(v string) *GetMdsMiniConfigResponseBody
	GetResultMessage() *string
}

type GetMdsMiniConfigResponseBody struct {
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *GetMdsMiniConfigResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMdsMiniConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMdsMiniConfigResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetMdsMiniConfigResponseBody) GetResultContent() *GetMdsMiniConfigResponseBodyResultContent {
	return s.ResultContent
}

func (s *GetMdsMiniConfigResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetMdsMiniConfigResponseBody) SetRequestId(v string) *GetMdsMiniConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMdsMiniConfigResponseBody) SetResultCode(v string) *GetMdsMiniConfigResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMdsMiniConfigResponseBody) SetResultContent(v *GetMdsMiniConfigResponseBodyResultContent) *GetMdsMiniConfigResponseBody {
	s.ResultContent = v
	return s
}

func (s *GetMdsMiniConfigResponseBody) SetResultMessage(v string) *GetMdsMiniConfigResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetMdsMiniConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMdsMiniConfigResponseBodyResultContent struct {
	Data      *GetMdsMiniConfigResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMdsMiniConfigResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigResponseBodyResultContent) GetData() *GetMdsMiniConfigResponseBodyResultContentData {
	return s.Data
}

func (s *GetMdsMiniConfigResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMdsMiniConfigResponseBodyResultContent) SetData(v *GetMdsMiniConfigResponseBodyResultContentData) *GetMdsMiniConfigResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContent) SetRequestId(v string) *GetMdsMiniConfigResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type GetMdsMiniConfigResponseBodyResultContentData struct {
	Content   *GetMdsMiniConfigResponseBodyResultContentDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string                                               `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMdsMiniConfigResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigResponseBodyResultContentData) GetContent() *GetMdsMiniConfigResponseBodyResultContentDataContent {
	return s.Content
}

func (s *GetMdsMiniConfigResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMdsMiniConfigResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *GetMdsMiniConfigResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *GetMdsMiniConfigResponseBodyResultContentData) SetContent(v *GetMdsMiniConfigResponseBodyResultContentDataContent) *GetMdsMiniConfigResponseBodyResultContentData {
	s.Content = v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentData) SetRequestId(v string) *GetMdsMiniConfigResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentData) SetResultMsg(v string) *GetMdsMiniConfigResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentData) SetSuccess(v bool) *GetMdsMiniConfigResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}

type GetMdsMiniConfigResponseBodyResultContentDataContent struct {
	ApiConfigList           []*GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList           `json:"ApiConfigList,omitempty" xml:"ApiConfigList,omitempty" type:"Repeated"`
	AppCode                 *string                                                                        `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	EnableServerDomainCount *string                                                                        `json:"EnableServerDomainCount,omitempty" xml:"EnableServerDomainCount,omitempty"`
	H5Id                    *string                                                                        `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name                  *string                                                                        `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	PrivilegeSwitch         *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch           `json:"PrivilegeSwitch,omitempty" xml:"PrivilegeSwitch,omitempty" type:"Struct"`
	ServerDomainConfigList  []*GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList  `json:"ServerDomainConfigList,omitempty" xml:"ServerDomainConfigList,omitempty" type:"Repeated"`
	WebviewDomainConfigList []*GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList `json:"WebviewDomainConfigList,omitempty" xml:"WebviewDomainConfigList,omitempty" type:"Repeated"`
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContent) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContent) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) GetApiConfigList() []*GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	return s.ApiConfigList
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) GetEnableServerDomainCount() *string {
	return s.EnableServerDomainCount
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) GetH5Id() *string {
	return s.H5Id
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) GetH5Name() *string {
	return s.H5Name
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) GetPrivilegeSwitch() *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	return s.PrivilegeSwitch
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) GetServerDomainConfigList() []*GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	return s.ServerDomainConfigList
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) GetWebviewDomainConfigList() []*GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	return s.WebviewDomainConfigList
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) SetApiConfigList(v []*GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) *GetMdsMiniConfigResponseBodyResultContentDataContent {
	s.ApiConfigList = v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) SetAppCode(v string) *GetMdsMiniConfigResponseBodyResultContentDataContent {
	s.AppCode = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) SetEnableServerDomainCount(v string) *GetMdsMiniConfigResponseBodyResultContentDataContent {
	s.EnableServerDomainCount = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) SetH5Id(v string) *GetMdsMiniConfigResponseBodyResultContentDataContent {
	s.H5Id = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) SetH5Name(v string) *GetMdsMiniConfigResponseBodyResultContentDataContent {
	s.H5Name = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) SetPrivilegeSwitch(v *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) *GetMdsMiniConfigResponseBodyResultContentDataContent {
	s.PrivilegeSwitch = v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) SetServerDomainConfigList(v []*GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) *GetMdsMiniConfigResponseBodyResultContentDataContent {
	s.ServerDomainConfigList = v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) SetWebviewDomainConfigList(v []*GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) *GetMdsMiniConfigResponseBodyResultContentDataContent {
	s.WebviewDomainConfigList = v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContent) Validate() error {
	return dara.Validate(s)
}

type GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList struct {
	AppCode      *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	ConfigStatus *int64  `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	ConfigType   *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	ConfigValue  *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	H5Id         *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name       *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetConfigStatus() *int64 {
	return s.ConfigStatus
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetDescription() *string {
	return s.Description
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetH5Id() *string {
	return s.H5Id
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetH5Name() *string {
	return s.H5Name
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) GetId() *int64 {
	return s.Id
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetAppCode(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.AppCode = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetConfigStatus(v int64) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.ConfigStatus = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetConfigType(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.ConfigType = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetConfigValue(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.ConfigValue = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetDescription(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.Description = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetGmtCreate(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.GmtCreate = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetGmtModified(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.GmtModified = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetH5Id(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.H5Id = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetH5Name(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.H5Name = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) SetId(v int64) *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList {
	s.Id = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentApiConfigList) Validate() error {
	return dara.Validate(s)
}

type GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch struct {
	AppCode      *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	ConfigStatus *int64  `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	ConfigType   *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	ConfigValue  *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	H5Id         *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name       *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetConfigStatus() *int64 {
	return s.ConfigStatus
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetDescription() *string {
	return s.Description
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetH5Id() *string {
	return s.H5Id
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetH5Name() *string {
	return s.H5Name
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) GetId() *int64 {
	return s.Id
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetAppCode(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.AppCode = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetConfigStatus(v int64) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.ConfigStatus = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetConfigType(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.ConfigType = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetConfigValue(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.ConfigValue = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetDescription(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.Description = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetGmtCreate(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.GmtCreate = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetGmtModified(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.GmtModified = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetH5Id(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.H5Id = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetH5Name(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.H5Name = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) SetId(v int64) *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch {
	s.Id = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentPrivilegeSwitch) Validate() error {
	return dara.Validate(s)
}

type GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList struct {
	AppCode      *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	ConfigStatus *int64  `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	ConfigType   *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	ConfigValue  *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	H5Id         *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name       *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetConfigStatus() *int64 {
	return s.ConfigStatus
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetDescription() *string {
	return s.Description
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetH5Id() *string {
	return s.H5Id
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetH5Name() *string {
	return s.H5Name
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) GetId() *int64 {
	return s.Id
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetAppCode(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.AppCode = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetConfigStatus(v int64) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.ConfigStatus = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetConfigType(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.ConfigType = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetConfigValue(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.ConfigValue = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetDescription(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.Description = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetGmtCreate(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.GmtCreate = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetGmtModified(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.GmtModified = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetH5Id(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.H5Id = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetH5Name(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.H5Name = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) SetId(v int64) *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList {
	s.Id = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentServerDomainConfigList) Validate() error {
	return dara.Validate(s)
}

type GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList struct {
	AppCode      *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	ConfigStatus *int64  `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	ConfigType   *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	ConfigValue  *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	H5Id         *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name       *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GoString() string {
	return s.String()
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetConfigStatus() *int64 {
	return s.ConfigStatus
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetDescription() *string {
	return s.Description
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetH5Id() *string {
	return s.H5Id
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetH5Name() *string {
	return s.H5Name
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) GetId() *int64 {
	return s.Id
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetAppCode(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.AppCode = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetConfigStatus(v int64) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.ConfigStatus = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetConfigType(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.ConfigType = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetConfigValue(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.ConfigValue = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetDescription(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.Description = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetGmtCreate(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.GmtCreate = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetGmtModified(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.GmtModified = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetH5Id(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.H5Id = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetH5Name(v string) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.H5Name = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) SetId(v int64) *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList {
	s.Id = &v
	return s
}

func (s *GetMdsMiniConfigResponseBodyResultContentDataContentWebviewDomainConfigList) Validate() error {
	return dara.Validate(s)
}
