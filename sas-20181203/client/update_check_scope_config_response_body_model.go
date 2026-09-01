// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckScopeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCheckScopeConfigResponseBody
	GetCode() *string
	SetData(v *UpdateCheckScopeConfigResponseBodyData) *UpdateCheckScopeConfigResponseBody
	GetData() *UpdateCheckScopeConfigResponseBodyData
	SetMessage(v string) *UpdateCheckScopeConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCheckScopeConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCheckScopeConfigResponseBody
	GetSuccess() *bool
}

type UpdateCheckScopeConfigResponseBody struct {
	// The result code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UpdateCheckScopeConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 676F80E3-4B3F-43DA-9CBB-5FF79F202AA2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCheckScopeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckScopeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCheckScopeConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCheckScopeConfigResponseBody) GetData() *UpdateCheckScopeConfigResponseBodyData {
	return s.Data
}

func (s *UpdateCheckScopeConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCheckScopeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCheckScopeConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCheckScopeConfigResponseBody) SetCode(v string) *UpdateCheckScopeConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBody) SetData(v *UpdateCheckScopeConfigResponseBodyData) *UpdateCheckScopeConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCheckScopeConfigResponseBody) SetMessage(v string) *UpdateCheckScopeConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBody) SetRequestId(v string) *UpdateCheckScopeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBody) SetSuccess(v bool) *UpdateCheckScopeConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCheckScopeConfigResponseBodyData struct {
	// The automatic scan configuration as a JSON string. The following fields are included:
	//
	// - **autoInclude**: specifies whether to enable automatic scan. Valid values: **true**: enabled. **false**: disabled.
	//
	// - **autoRule**: the enablement configuration.
	//
	// - **ruleOperator**: the enablement configuration rule. Set the value to **include**.
	//
	// - **operator**: the logical operator. Set the value to **or**.
	//
	// - **rule**: the rule.
	//
	// - **condition**: the rule condition. Valid values: **vendor**: vendor. **assetType**: level-1 asset type. **assetSubType**: level-2 asset type.
	//
	// > For more information, refer to the [GetCloudAssetCriteria](~~GetCloudAssetCriteria~~) operation.
	//
	// example:
	//
	// "{\\"autoInclude\\":true,\\"autoRule\\":{\\"ruleOperator\\":\\"include\\",\\"operator\\":\\"or\\",\\"rule\\":[{\\"condition\\":\\"assetSubType\\",\\"ruleOperator\\":\\"include\\",\\"value\\":[{\\"vendor\\":\\"0\\",\\"assetType\\":\\"0\\",\\"assetSubType\\":\\"100\\"}]}]}}"
	AutoConfig *string `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	// The automatic scan configuration type. Valid values:
	//
	// - **0**: disable automatic scan
	//
	// - **1**: automatically scan newly added cloud assets
	//
	// example:
	//
	// 1
	AutoType *int32 `json:"AutoType,omitempty" xml:"AutoType,omitempty"`
	// The ID of the configuration.
	//
	// example:
	//
	// 97a1fed216908e417407344e1505xxxx
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-10-16 18:17:16
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-04-09 18:56:15
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The scan scope configuration type. Valid values:
	//
	// - **1**: scan by instance
	//
	// - **3**: scan all
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateCheckScopeConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckScopeConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCheckScopeConfigResponseBodyData) GetAutoConfig() *string {
	return s.AutoConfig
}

func (s *UpdateCheckScopeConfigResponseBodyData) GetAutoType() *int32 {
	return s.AutoType
}

func (s *UpdateCheckScopeConfigResponseBodyData) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateCheckScopeConfigResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateCheckScopeConfigResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateCheckScopeConfigResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *UpdateCheckScopeConfigResponseBodyData) SetAutoConfig(v string) *UpdateCheckScopeConfigResponseBodyData {
	s.AutoConfig = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBodyData) SetAutoType(v int32) *UpdateCheckScopeConfigResponseBodyData {
	s.AutoType = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBodyData) SetConfigId(v string) *UpdateCheckScopeConfigResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBodyData) SetGmtCreate(v string) *UpdateCheckScopeConfigResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBodyData) SetGmtModified(v string) *UpdateCheckScopeConfigResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBodyData) SetType(v int32) *UpdateCheckScopeConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *UpdateCheckScopeConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
