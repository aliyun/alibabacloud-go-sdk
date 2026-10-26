// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppPluginResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppPluginResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppPluginResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppPluginResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppPluginResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppPluginResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppPluginResponseBodyModule) *GetAppPluginResponseBody
	GetModule() *GetAppPluginResponseBodyModule
	SetRequestId(v string) *GetAppPluginResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppPluginResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppPluginResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppPluginResponseBody
	GetSynchro() *bool
}

type GetAppPluginResponseBody struct {
	// The detailed reason for the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// dewuApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace `%s` in the error message returned by the **ErrMessage*	- parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The returned error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The application module.
	Module *GetAppPluginResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The exception message.
	//
	// example:
	//
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppPluginResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppPluginResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppPluginResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppPluginResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppPluginResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppPluginResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppPluginResponseBody) GetModule() *GetAppPluginResponseBodyModule {
	return s.Module
}

func (s *GetAppPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppPluginResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppPluginResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppPluginResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppPluginResponseBody) SetAccessDeniedDetail(v string) *GetAppPluginResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppPluginResponseBody) SetAllowRetry(v bool) *GetAppPluginResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppPluginResponseBody) SetAppName(v string) *GetAppPluginResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppPluginResponseBody) SetDynamicCode(v string) *GetAppPluginResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppPluginResponseBody) SetDynamicMessage(v string) *GetAppPluginResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppPluginResponseBody) SetErrorArgs(v []interface{}) *GetAppPluginResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppPluginResponseBody) SetModule(v *GetAppPluginResponseBodyModule) *GetAppPluginResponseBody {
	s.Module = v
	return s
}

func (s *GetAppPluginResponseBody) SetRequestId(v string) *GetAppPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppPluginResponseBody) SetRootErrorCode(v string) *GetAppPluginResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppPluginResponseBody) SetRootErrorMsg(v string) *GetAppPluginResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppPluginResponseBody) SetSynchro(v bool) *GetAppPluginResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppPluginResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppPluginResponseBodyModule struct {
	// The category.
	//
	// example:
	//
	// Archive
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The configuration form in React JSON Schema format.
	//
	// example:
	//
	// ***
	ConfigItems *string `json:"ConfigItems,omitempty" xml:"ConfigItems,omitempty"`
	// The creator.
	//
	// example:
	//
	// adminSubuser(264859988988056208)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The application description.
	//
	// example:
	//
	// /bak->serverless.handler(2020091300200279)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image display mode. Valid values:
	//
	// - **0*	- (None): The image is not displayed.
	//
	// - **1*	- (Always): The image is always displayed.
	//
	// example:
	//
	// always
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// Indicates whether scheduled delivery of resource snapshots is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// True
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {"deliveryNodeName":"Bind domain name and publish","deliveryNodeStatus":"Finish","deliveryOperatorRole":"Provider"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The version number.
	//
	// example:
	//
	// v0.1.0
	GitRef *string `json:"GitRef,omitempty" xml:"GitRef,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-08-28T02:25:41Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The hook definitions.
	//
	// example:
	//
	// []
	Hooks *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// The popularity.
	//
	// example:
	//
	// 1
	HotCount *int32 `json:"HotCount,omitempty" xml:"HotCount,omitempty"`
	// The plugin description.
	//
	// example:
	//
	// @lALPM2AwTOg9IUHNAUDNAUA
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The ID of the associated API gateway plugin.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The name of the plugin. The name must be 4 to 50 characters in length and can contain uppercase and lowercase letters, Chinese characters, digits, and underscores (_). It cannot start with an underscore.
	//
	// example:
	//
	// tf_testaccapigatewayplugin29311
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The URL of the image repository.
	//
	// example:
	//
	// xxx
	RepositoryUrl *string `json:"RepositoryUrl,omitempty" xml:"RepositoryUrl,omitempty"`
	// The skill header information for model selection.
	//
	// example:
	//
	// header
	SkillHeader *string `json:"SkillHeader,omitempty" xml:"SkillHeader,omitempty"`
	// trial,draft,live,refunded,expired,released
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The category tags.
	//
	// example:
	//
	// [{\\"Key\\": \\"kubernetes.io/cluster-id\\", \\"Value\\": \\"cc67198b13db948c9848599654da5586e\\"}, {\\"Key\\": \\"created-by\\", \\"Value\\": \\"alibabacloud-imagecache-controller\\"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The version of the application instance.
	//
	// example:
	//
	// 2019-04-02
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// **The visibility level.**
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s GetAppPluginResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppPluginResponseBodyModule) GetCategory() *string {
	return s.Category
}

func (s *GetAppPluginResponseBodyModule) GetConfigItems() *string {
	return s.ConfigItems
}

func (s *GetAppPluginResponseBodyModule) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetAppPluginResponseBodyModule) GetDescription() *string {
	return s.Description
}

func (s *GetAppPluginResponseBodyModule) GetDisplay() *int32 {
	return s.Display
}

func (s *GetAppPluginResponseBodyModule) GetEnabled() *int32 {
	return s.Enabled
}

func (s *GetAppPluginResponseBodyModule) GetExtend() *string {
	return s.Extend
}

func (s *GetAppPluginResponseBodyModule) GetGitRef() *string {
	return s.GitRef
}

func (s *GetAppPluginResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAppPluginResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAppPluginResponseBodyModule) GetHooks() *string {
	return s.Hooks
}

func (s *GetAppPluginResponseBodyModule) GetHotCount() *int32 {
	return s.HotCount
}

func (s *GetAppPluginResponseBodyModule) GetIcon() *string {
	return s.Icon
}

func (s *GetAppPluginResponseBodyModule) GetPluginId() *string {
	return s.PluginId
}

func (s *GetAppPluginResponseBodyModule) GetPluginName() *string {
	return s.PluginName
}

func (s *GetAppPluginResponseBodyModule) GetRepositoryUrl() *string {
	return s.RepositoryUrl
}

func (s *GetAppPluginResponseBodyModule) GetSkillHeader() *string {
	return s.SkillHeader
}

func (s *GetAppPluginResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *GetAppPluginResponseBodyModule) GetTags() *string {
	return s.Tags
}

func (s *GetAppPluginResponseBodyModule) GetVersion() *string {
	return s.Version
}

func (s *GetAppPluginResponseBodyModule) GetVisibility() *string {
	return s.Visibility
}

func (s *GetAppPluginResponseBodyModule) SetCategory(v string) *GetAppPluginResponseBodyModule {
	s.Category = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetConfigItems(v string) *GetAppPluginResponseBodyModule {
	s.ConfigItems = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetCreatedBy(v string) *GetAppPluginResponseBodyModule {
	s.CreatedBy = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetDescription(v string) *GetAppPluginResponseBodyModule {
	s.Description = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetDisplay(v int32) *GetAppPluginResponseBodyModule {
	s.Display = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetEnabled(v int32) *GetAppPluginResponseBodyModule {
	s.Enabled = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetExtend(v string) *GetAppPluginResponseBodyModule {
	s.Extend = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetGitRef(v string) *GetAppPluginResponseBodyModule {
	s.GitRef = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetGmtCreate(v string) *GetAppPluginResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetGmtModified(v string) *GetAppPluginResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetHooks(v string) *GetAppPluginResponseBodyModule {
	s.Hooks = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetHotCount(v int32) *GetAppPluginResponseBodyModule {
	s.HotCount = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetIcon(v string) *GetAppPluginResponseBodyModule {
	s.Icon = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetPluginId(v string) *GetAppPluginResponseBodyModule {
	s.PluginId = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetPluginName(v string) *GetAppPluginResponseBodyModule {
	s.PluginName = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetRepositoryUrl(v string) *GetAppPluginResponseBodyModule {
	s.RepositoryUrl = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetSkillHeader(v string) *GetAppPluginResponseBodyModule {
	s.SkillHeader = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetStatus(v string) *GetAppPluginResponseBodyModule {
	s.Status = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetTags(v string) *GetAppPluginResponseBodyModule {
	s.Tags = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetVersion(v string) *GetAppPluginResponseBodyModule {
	s.Version = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) SetVisibility(v string) *GetAppPluginResponseBodyModule {
	s.Visibility = &v
	return s
}

func (s *GetAppPluginResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
