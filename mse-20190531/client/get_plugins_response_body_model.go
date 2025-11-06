// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetPluginsResponseBody
	GetCode() *int32
	SetData(v []*GetPluginsResponseBodyData) *GetPluginsResponseBody
	GetData() []*GetPluginsResponseBodyData
	SetDynamicCode(v string) *GetPluginsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetPluginsResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *GetPluginsResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetPluginsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPluginsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPluginsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPluginsResponseBody
	GetSuccess() *bool
}

type GetPluginsResponseBody struct {
	// The returned code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data entries returned.
	Data []*GetPluginsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// message
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 500
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 03A3E2F4-6804-5663-9D5D-2EC47A1*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPluginsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPluginsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetPluginsResponseBody) GetData() []*GetPluginsResponseBodyData {
	return s.Data
}

func (s *GetPluginsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetPluginsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetPluginsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPluginsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPluginsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPluginsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPluginsResponseBody) SetCode(v int32) *GetPluginsResponseBody {
	s.Code = &v
	return s
}

func (s *GetPluginsResponseBody) SetData(v []*GetPluginsResponseBodyData) *GetPluginsResponseBody {
	s.Data = v
	return s
}

func (s *GetPluginsResponseBody) SetDynamicCode(v string) *GetPluginsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetPluginsResponseBody) SetDynamicMessage(v string) *GetPluginsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetPluginsResponseBody) SetErrorCode(v string) *GetPluginsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPluginsResponseBody) SetHttpStatusCode(v int32) *GetPluginsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPluginsResponseBody) SetMessage(v string) *GetPluginsResponseBody {
	s.Message = &v
	return s
}

func (s *GetPluginsResponseBody) SetRequestId(v string) *GetPluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPluginsResponseBody) SetSuccess(v bool) *GetPluginsResponseBody {
	s.Success = &v
	return s
}

func (s *GetPluginsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPluginsResponseBodyData struct {
	// The type of the plug-in. Valid values:
	//
	// 0: custom
	//
	// 1: permission authorization
	//
	// 2: security protection
	//
	// 3: transmission protocol
	//
	// 4: traffic control
	//
	// 5: traffic observation
	//
	// example:
	//
	// 0
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The information about the plug-in configuration used for checking.
	//
	// example:
	//
	// \\# The configuration includes the fields required for checking, such as name, age, and friends. Sample configuration: name: John age: 18 friends: - David - Anne
	ConfigCheck *string `json:"ConfigCheck,omitempty" xml:"ConfigCheck,omitempty"`
	// The ID of the plug-in.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1.0.0
	MaxVersion *string `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	// example:
	//
	// 0
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// key-auth
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	NewVersionPublishingFlag *bool `json:"NewVersionPublishingFlag,omitempty" xml:"NewVersionPublishingFlag,omitempty"`
	// The execution stage of the plug-in.
	//
	// 	- 0: default stage
	//
	// 	- 1: authorization stage
	//
	// 	- 2: authentication stage
	//
	// 	- 3: statistics stage
	//
	// example:
	//
	// 1
	Phase *int32 `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// 123
	PrimaryUser *string `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	// The execution priority of the plug-in. A larger value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The publish status.
	//
	// example:
	//
	// 1
	PublishState *int32 `json:"PublishState,omitempty" xml:"PublishState,omitempty"`
	// Indicates whether the plug-in is enabled.
	//
	// 	- 0: disabled
	//
	// 	- 1: enabled
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The summary of the plug-in.
	//
	// example:
	//
	// This is a plug-in.
	Summary   *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	SummaryEn *string `json:"SummaryEn,omitempty" xml:"SummaryEn,omitempty"`
	// The version of the plug-in.
	//
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The URL of the Object Storage Service (OSS) bucket that stores the WebAssembly plug-in.
	WasmFile *string `json:"WasmFile,omitempty" xml:"WasmFile,omitempty"`
	// The WebAssembly language. Valid values:
	//
	// 	- 0: C++
	//
	// 	- 1: TinyGo
	//
	// 	- 2: Rust
	//
	// 	- 3: AssemblyScript
	//
	// 	- 4: Zig
	//
	// example:
	//
	// 0
	WasmLang *int32 `json:"WasmLang,omitempty" xml:"WasmLang,omitempty"`
}

func (s GetPluginsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPluginsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPluginsResponseBodyData) GetCategory() *int32 {
	return s.Category
}

func (s *GetPluginsResponseBodyData) GetConfigCheck() *string {
	return s.ConfigCheck
}

func (s *GetPluginsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetPluginsResponseBodyData) GetMaxVersion() *string {
	return s.MaxVersion
}

func (s *GetPluginsResponseBodyData) GetMode() *int32 {
	return s.Mode
}

func (s *GetPluginsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetPluginsResponseBodyData) GetNewVersionPublishingFlag() *bool {
	return s.NewVersionPublishingFlag
}

func (s *GetPluginsResponseBodyData) GetPhase() *int32 {
	return s.Phase
}

func (s *GetPluginsResponseBodyData) GetPrimaryUser() *string {
	return s.PrimaryUser
}

func (s *GetPluginsResponseBodyData) GetPriority() *int32 {
	return s.Priority
}

func (s *GetPluginsResponseBodyData) GetPublishState() *int32 {
	return s.PublishState
}

func (s *GetPluginsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetPluginsResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *GetPluginsResponseBodyData) GetSummaryEn() *string {
	return s.SummaryEn
}

func (s *GetPluginsResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetPluginsResponseBodyData) GetWasmFile() *string {
	return s.WasmFile
}

func (s *GetPluginsResponseBodyData) GetWasmLang() *int32 {
	return s.WasmLang
}

func (s *GetPluginsResponseBodyData) SetCategory(v int32) *GetPluginsResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetConfigCheck(v string) *GetPluginsResponseBodyData {
	s.ConfigCheck = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetId(v int64) *GetPluginsResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetMaxVersion(v string) *GetPluginsResponseBodyData {
	s.MaxVersion = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetMode(v int32) *GetPluginsResponseBodyData {
	s.Mode = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetName(v string) *GetPluginsResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetNewVersionPublishingFlag(v bool) *GetPluginsResponseBodyData {
	s.NewVersionPublishingFlag = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetPhase(v int32) *GetPluginsResponseBodyData {
	s.Phase = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetPrimaryUser(v string) *GetPluginsResponseBodyData {
	s.PrimaryUser = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetPriority(v int32) *GetPluginsResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetPublishState(v int32) *GetPluginsResponseBodyData {
	s.PublishState = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetStatus(v string) *GetPluginsResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetSummary(v string) *GetPluginsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetSummaryEn(v string) *GetPluginsResponseBodyData {
	s.SummaryEn = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetVersion(v string) *GetPluginsResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetWasmFile(v string) *GetPluginsResponseBodyData {
	s.WasmFile = &v
	return s
}

func (s *GetPluginsResponseBodyData) SetWasmLang(v int32) *GetPluginsResponseBodyData {
	s.WasmLang = &v
	return s
}

func (s *GetPluginsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
