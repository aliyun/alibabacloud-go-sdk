// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrialAccessChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTrialAccessChannelResponseBody
	GetCode() *string
	SetData(v *GetTrialAccessChannelResponseBodyData) *GetTrialAccessChannelResponseBody
	GetData() *GetTrialAccessChannelResponseBodyData
	SetHttpStatusCode(v int32) *GetTrialAccessChannelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTrialAccessChannelResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTrialAccessChannelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTrialAccessChannelResponseBody
	GetSuccess() *bool
}

type GetTrialAccessChannelResponseBody struct {
	// The business processing status code. A value of OK indicates success. When an exception occurs, the corresponding business error code is returned. The error types and mapping rules follow the current API configuration. No error codes are added or modified in this update.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The channel configuration corresponding to the VisitorToken and its associated scenarios, used for trial page initialization.
	//
	// example:
	//
	// {"InstanceId":"instance-example-001","AccessChannelId":"66666666-7777-4888-8999-000000000000","AccessChannelType":"TEST","Name":"Product consultation trial","State":"Active","ChannelSetting":"{}","VisitorToken":"example-visitor-token","ExpirationTime":1794564000000,"CreatedTime":1794556800000,"UpdatedTime":1794560400000,"ChannelScripts":[{"ScriptId":"script-example-001","ScriptName":"Product consultation","VariableConfigs":"[]","LabelConfigs":"[]"}],"NluEngine":"BEEBOT","BuilderType":"FLOW"}
	Data *GetTrialAccessChannelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code recorded in the business response body. A normal return value is 200. To determine whether the business operation is successful, you must also check the Code field. The presence of this field alone does not indicate a successful call.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business response or error description. It is usually empty or not returned upon success, and a fixed text cannot be required for successful responses. Upon failure, it provides an error description to assist in troubleshooting. Combine it with Code and Params to understand the cause of the error. The example shows an error message when the trial link visitor token expires (corresponding to the Code Expired.VisitorToken) and is for format reference only.
	//
	// example:
	//
	// VisitorToken: example-visitor-token is expired.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID used to associate and troubleshoot this POP request. The value in the response corresponds to the current request. System request parameters are injected by the gateway, and callers do not need to fill in system fields manually.
	//
	// example:
	//
	// AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the business processing is successful. A value of true indicates success, and false indicates failure. The current general response implementation determines success based on whether the Code is OK.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTrialAccessChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrialAccessChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrialAccessChannelResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTrialAccessChannelResponseBody) GetData() *GetTrialAccessChannelResponseBodyData {
	return s.Data
}

func (s *GetTrialAccessChannelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTrialAccessChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTrialAccessChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrialAccessChannelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTrialAccessChannelResponseBody) SetCode(v string) *GetTrialAccessChannelResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrialAccessChannelResponseBody) SetData(v *GetTrialAccessChannelResponseBodyData) *GetTrialAccessChannelResponseBody {
	s.Data = v
	return s
}

func (s *GetTrialAccessChannelResponseBody) SetHttpStatusCode(v int32) *GetTrialAccessChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTrialAccessChannelResponseBody) SetMessage(v string) *GetTrialAccessChannelResponseBody {
	s.Message = &v
	return s
}

func (s *GetTrialAccessChannelResponseBody) SetRequestId(v string) *GetTrialAccessChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrialAccessChannelResponseBody) SetSuccess(v bool) *GetTrialAccessChannelResponseBody {
	s.Success = &v
	return s
}

func (s *GetTrialAccessChannelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrialAccessChannelResponseBodyData struct {
	// The ID of the channel associated with the trial link or specific access configuration. When a channel needs to be associated, use the ID actually returned by the channel service. It is not the ChannelId returned by a call, and the two must not be mixed. The configuration layer uses a UUID as the channel ID when creating a new channel.
	//
	// example:
	//
	// 66666666-7777-4888-8999-000000000000
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// Identifies the access source of the current interaction. The configuration layer enumerations are TEST (voice test window or real number outbound call test), TEXT_TEST (text test window), TRIAL (web voice trial link), and TEXT_TRIAL (text trial link). The existing default values and visibility of each API remain unchanged. If the current definition uses the historical value Text, it should be processed according to the existing definition and is not modified in this documentation update.
	//
	// example:
	//
	// TEST
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	// The scenario building method. The configuration layer enumerations are FLOW and LITE. This field is not declared in the current business layer channel response object. Therefore, whether it is actually returned depends on the deployment version.
	//
	// example:
	//
	// FLOW
	BuilderType *string `json:"BuilderType,omitempty" xml:"BuilderType,omitempty"`
	// The scenarios associated with this channel and their display, variable, and label configurations. The trial page should use the ScriptId actually returned in the list and must not specify scenarios outside the channel.
	//
	// example:
	//
	// [{"ScriptId":"script-example-001","ScriptName":"Product consultation","VariableConfigs":"[]","LabelConfigs":"[]"}]
	ChannelScripts []*GetTrialAccessChannelResponseBodyDataChannelScripts `json:"ChannelScripts,omitempty" xml:"ChannelScripts,omitempty" type:"Repeated"`
	// The display configuration of the trial channel, which is a JSON object string. The content is configured during channel creation and agreed upon with the trial page, and may include display information such as Images. The current code does not provide a complete fixed structure, and the empty object in the example only illustrates the serialization format.
	//
	// example:
	//
	// {}
	ChannelSetting *string `json:"ChannelSetting,omitempty" xml:"ChannelSetting,omitempty"`
	// The creation time of the channel, in Unix timestamp format and measured in milliseconds. It is converted from the creation time of the configuration layer entity.
	//
	// example:
	//
	// 1794556800000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The expiration time of the trial channel and the visitor token, in Unix timestamp format and measured in milliseconds. The service compares the validity period with the current system time in milliseconds. Expired tokens cannot continue the trial.
	//
	// example:
	//
	// 1794564000000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The instance identity of the product to which it belongs. You must use an instance that the invoker has permission to access and that actually exists. It must be consistent with the instance to which the scenario, channel, and session belong. The maximum length is subject to the current POP parameter definition.
	//
	// example:
	//
	// instance-example-001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the trial channel, used for display and identification on the trial page. The content is configured by the channel creator.
	//
	// example:
	//
	// Product consultation trial
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The natural language understanding (NLU) engine type defined in the configuration layer. Valid values: BEEBOT, PROMPTS, and FUNCTION. This field is not declared in the current business layer channel response object. Therefore, whether it can be returned through this operation depends on the deployment version and must not be assumed to always exist on the client.
	//
	// example:
	//
	// BEEBOT
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// The channel status. Valid values: Active (effective) and Expired (invalid). Trial access still requires server-side token and validity period verification. You cannot bypass the verification based solely on the status value saved on the client.
	//
	// example:
	//
	// Active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The last update time of the channel, in Unix timestamp format and measured in milliseconds. It is converted from the update time of the configuration layer entity.
	//
	// example:
	//
	// 1794560400000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The visitor token obtained from the trial link. It is used to locate the trial channel and verify access permissions and the validity period. The token is generated by the channel service and must not be constructed manually. It cannot be empty and must be within its validity period. The example is a placeholder value that cannot be used for authentication. Do not disclose actual tokens to unauthorized personnel.
	//
	// example:
	//
	// example-visitor-token
	VisitorToken *string `json:"VisitorToken,omitempty" xml:"VisitorToken,omitempty"`
}

func (s GetTrialAccessChannelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTrialAccessChannelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrialAccessChannelResponseBodyData) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *GetTrialAccessChannelResponseBodyData) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *GetTrialAccessChannelResponseBodyData) GetBuilderType() *string {
	return s.BuilderType
}

func (s *GetTrialAccessChannelResponseBodyData) GetChannelScripts() []*GetTrialAccessChannelResponseBodyDataChannelScripts {
	return s.ChannelScripts
}

func (s *GetTrialAccessChannelResponseBodyData) GetChannelSetting() *string {
	return s.ChannelSetting
}

func (s *GetTrialAccessChannelResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetTrialAccessChannelResponseBodyData) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GetTrialAccessChannelResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTrialAccessChannelResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTrialAccessChannelResponseBodyData) GetNluEngine() *string {
	return s.NluEngine
}

func (s *GetTrialAccessChannelResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetTrialAccessChannelResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetTrialAccessChannelResponseBodyData) GetVisitorToken() *string {
	return s.VisitorToken
}

func (s *GetTrialAccessChannelResponseBodyData) SetAccessChannelId(v string) *GetTrialAccessChannelResponseBodyData {
	s.AccessChannelId = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetAccessChannelType(v string) *GetTrialAccessChannelResponseBodyData {
	s.AccessChannelType = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetBuilderType(v string) *GetTrialAccessChannelResponseBodyData {
	s.BuilderType = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetChannelScripts(v []*GetTrialAccessChannelResponseBodyDataChannelScripts) *GetTrialAccessChannelResponseBodyData {
	s.ChannelScripts = v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetChannelSetting(v string) *GetTrialAccessChannelResponseBodyData {
	s.ChannelSetting = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetCreatedTime(v int64) *GetTrialAccessChannelResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetExpirationTime(v int64) *GetTrialAccessChannelResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetInstanceId(v string) *GetTrialAccessChannelResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetName(v string) *GetTrialAccessChannelResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetNluEngine(v string) *GetTrialAccessChannelResponseBodyData {
	s.NluEngine = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetState(v string) *GetTrialAccessChannelResponseBodyData {
	s.State = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetUpdatedTime(v int64) *GetTrialAccessChannelResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) SetVisitorToken(v string) *GetTrialAccessChannelResponseBodyData {
	s.VisitorToken = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyData) Validate() error {
	if s.ChannelScripts != nil {
		for _, item := range s.ChannelScripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTrialAccessChannelResponseBodyDataChannelScripts struct {
	// The JSON string of the label configuration exported from the scenario version. The specific fields are determined by the scenario version. The empty JSON array in the example only indicates that there are no configuration items and must not be confused with the call parameter Tags.
	//
	// example:
	//
	// []
	LabelConfigs *string `json:"LabelConfigs,omitempty" xml:"LabelConfigs,omitempty"`
	// The ID of the dialog scenario under the instance. Specify an existing scenario. Subsequent dialogs and end requests for the same text session should use the scenario specified when the session is started.
	//
	// example:
	//
	// script-example-001
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The name of the scenario associated with the channel, used for display on the trial page. Use the actual value returned by the service.
	//
	// example:
	//
	// Product consultation
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The JSON string of the variable configuration exported from the scenario version. It is not the variable values passed in this interaction. The specific configuration items are determined by the scenario version. The empty JSON array in the example only indicates that there are no configuration items and does not declare unverified internal fields.
	//
	// example:
	//
	// []
	VariableConfigs *string `json:"VariableConfigs,omitempty" xml:"VariableConfigs,omitempty"`
}

func (s GetTrialAccessChannelResponseBodyDataChannelScripts) String() string {
	return dara.Prettify(s)
}

func (s GetTrialAccessChannelResponseBodyDataChannelScripts) GoString() string {
	return s.String()
}

func (s *GetTrialAccessChannelResponseBodyDataChannelScripts) GetLabelConfigs() *string {
	return s.LabelConfigs
}

func (s *GetTrialAccessChannelResponseBodyDataChannelScripts) GetScriptId() *string {
	return s.ScriptId
}

func (s *GetTrialAccessChannelResponseBodyDataChannelScripts) GetScriptName() *string {
	return s.ScriptName
}

func (s *GetTrialAccessChannelResponseBodyDataChannelScripts) GetVariableConfigs() *string {
	return s.VariableConfigs
}

func (s *GetTrialAccessChannelResponseBodyDataChannelScripts) SetLabelConfigs(v string) *GetTrialAccessChannelResponseBodyDataChannelScripts {
	s.LabelConfigs = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyDataChannelScripts) SetScriptId(v string) *GetTrialAccessChannelResponseBodyDataChannelScripts {
	s.ScriptId = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyDataChannelScripts) SetScriptName(v string) *GetTrialAccessChannelResponseBodyDataChannelScripts {
	s.ScriptName = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyDataChannelScripts) SetVariableConfigs(v string) *GetTrialAccessChannelResponseBodyDataChannelScripts {
	s.VariableConfigs = &v
	return s
}

func (s *GetTrialAccessChannelResponseBodyDataChannelScripts) Validate() error {
	return dara.Validate(s)
}
