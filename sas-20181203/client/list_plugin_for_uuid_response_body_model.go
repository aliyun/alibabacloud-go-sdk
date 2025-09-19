// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginForUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAegisUuidTargetPluginConfigList(v []*ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) *ListPluginForUuidResponseBody
	GetAegisUuidTargetPluginConfigList() []*ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList
	SetCode(v int32) *ListPluginForUuidResponseBody
	GetCode() *int32
	SetMessage(v string) *ListPluginForUuidResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPluginForUuidResponseBody
	GetRequestId() *string
}

type ListPluginForUuidResponseBody struct {
	// An array that consists of the information about the plug-ins.
	AegisUuidTargetPluginConfigList []*ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList `json:"AegisUuidTargetPluginConfigList,omitempty" xml:"AegisUuidTargetPluginConfigList,omitempty" type:"Repeated"`
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPluginForUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPluginForUuidResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginForUuidResponseBody) GetAegisUuidTargetPluginConfigList() []*ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList {
	return s.AegisUuidTargetPluginConfigList
}

func (s *ListPluginForUuidResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListPluginForUuidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPluginForUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPluginForUuidResponseBody) SetAegisUuidTargetPluginConfigList(v []*ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) *ListPluginForUuidResponseBody {
	s.AegisUuidTargetPluginConfigList = v
	return s
}

func (s *ListPluginForUuidResponseBody) SetCode(v int32) *ListPluginForUuidResponseBody {
	s.Code = &v
	return s
}

func (s *ListPluginForUuidResponseBody) SetMessage(v string) *ListPluginForUuidResponseBody {
	s.Message = &v
	return s
}

func (s *ListPluginForUuidResponseBody) SetRequestId(v string) *ListPluginForUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginForUuidResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList struct {
	// An array that consists of the configurations of plug-ins.
	AegisSuspiciousConfigList []*ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList `json:"AegisSuspiciousConfigList,omitempty" xml:"AegisSuspiciousConfigList,omitempty" type:"Repeated"`
	// The installation code of the plug-in.
	//
	// example:
	//
	// k5O5nd
	PluginInstallCode *string `json:"PluginInstallCode,omitempty" xml:"PluginInstallCode,omitempty"`
	// The name of the plug-in. Valid values:
	//
	// 	- **alihips**: trojan-specific prevention
	//
	// 	- **alisecguard**: attack-specific prevention
	//
	// 	- **alinet**: defense against attacks on servers
	//
	// example:
	//
	// alisecguard
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// Indicates whether the plug-in is installed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	PluginOnlineInstalled *bool `json:"PluginOnlineInstalled,omitempty" xml:"PluginOnlineInstalled,omitempty"`
	// Indicates whether the plug-in is online. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	PluginOnlineStatus *bool `json:"PluginOnlineStatus,omitempty" xml:"PluginOnlineStatus,omitempty"`
	// The version of the plug-in.
	//
	// example:
	//
	// 00_10
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
}

func (s ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) GoString() string {
	return s.String()
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) GetAegisSuspiciousConfigList() []*ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList {
	return s.AegisSuspiciousConfigList
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) GetPluginInstallCode() *string {
	return s.PluginInstallCode
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) GetPluginName() *string {
	return s.PluginName
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) GetPluginOnlineInstalled() *bool {
	return s.PluginOnlineInstalled
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) GetPluginOnlineStatus() *bool {
	return s.PluginOnlineStatus
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) SetAegisSuspiciousConfigList(v []*ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList {
	s.AegisSuspiciousConfigList = v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) SetPluginInstallCode(v string) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList {
	s.PluginInstallCode = &v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) SetPluginName(v string) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList {
	s.PluginName = &v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) SetPluginOnlineInstalled(v bool) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList {
	s.PluginOnlineInstalled = &v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) SetPluginOnlineStatus(v bool) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList {
	s.PluginOnlineStatus = &v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) SetPluginVersion(v string) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList {
	s.PluginVersion = &v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigList) Validate() error {
	return dara.Validate(s)
}

type ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList struct {
	// Indicates whether the plug-in is enabled. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Config *bool `json:"Config,omitempty" xml:"Config,omitempty"`
	// The message that indicates whether you are authorized to install the plug-in on your server or whether the plug-in is installed on your server. Valid values:
	//
	// 	- **authorized**: authorized
	//
	// 	- **unauthorized**: unauthorized
	//
	// 	- **unbind**: not installed
	//
	// 	- **nonsupport**: not supported
	//
	// example:
	//
	// authorized
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Indicates whether the plug-in is globally configured. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	OverallConfig *bool `json:"OverallConfig,omitempty" xml:"OverallConfig,omitempty"`
	// The name of the plug-in. Valid values:
	//
	// 	- **alihips**: trojan-specific prevention
	//
	// 	- **alisecguard**: attack-specific prevention
	//
	// 	- **alinet**: defense against attacks on servers
	//
	// example:
	//
	// alisecguard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) GoString() string {
	return s.String()
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) GetConfig() *bool {
	return s.Config
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) GetMsg() *string {
	return s.Msg
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) GetOverallConfig() *bool {
	return s.OverallConfig
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) GetType() *string {
	return s.Type
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) SetConfig(v bool) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList {
	s.Config = &v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) SetMsg(v string) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList {
	s.Msg = &v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) SetOverallConfig(v bool) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList {
	s.OverallConfig = &v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) SetType(v string) *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList {
	s.Type = &v
	return s
}

func (s *ListPluginForUuidResponseBodyAegisUuidTargetPluginConfigListAegisSuspiciousConfigList) Validate() error {
	return dara.Validate(s)
}
