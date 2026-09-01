// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckScopeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCheckScopeConfigResponseBody
	GetCode() *string
	SetData(v *GetCheckScopeConfigResponseBodyData) *GetCheckScopeConfigResponseBody
	GetData() *GetCheckScopeConfigResponseBodyData
	SetMessage(v string) *GetCheckScopeConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCheckScopeConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCheckScopeConfigResponseBody
	GetSuccess() *bool
}

type GetCheckScopeConfigResponseBody struct {
	// The result code. A value of **200*	- indicates success. Any other value indicates failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetCheckScopeConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
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

func (s GetCheckScopeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckScopeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckScopeConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCheckScopeConfigResponseBody) GetData() *GetCheckScopeConfigResponseBodyData {
	return s.Data
}

func (s *GetCheckScopeConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCheckScopeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckScopeConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCheckScopeConfigResponseBody) SetCode(v string) *GetCheckScopeConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetCheckScopeConfigResponseBody) SetData(v *GetCheckScopeConfigResponseBodyData) *GetCheckScopeConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetCheckScopeConfigResponseBody) SetMessage(v string) *GetCheckScopeConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetCheckScopeConfigResponseBody) SetRequestId(v string) *GetCheckScopeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckScopeConfigResponseBody) SetSuccess(v bool) *GetCheckScopeConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetCheckScopeConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCheckScopeConfigResponseBodyData struct {
	// The automatic scan configuration as a JSON string. The following fields are included:
	//
	// - **autoInclude**: specifies whether to enable automatic scanning. Valid values: **true**: enabled. **false**: disabled.
	//
	// - **autoRule**: the enablement configuration.
	//
	// - **ruleOperator**: the enablement configuration rule. The value is **include**.
	//
	// - **operator**: the logical operator. The value is **or**.
	//
	// - **rule**: the rule.
	//
	// - **condition**: the rule condition. Valid values: **vendor**: vendor. **assetType**: level-1 asset type. **assetSubType**: level-2 asset type.
	//
	// > For more information, see the [GetCloudAssetCriteria](~~GetCloudAssetCriteria~~) operation.
	//
	// example:
	//
	// "{\\"autoInclude\\":true,\\"autoRule\\":{\\"ruleOperator\\":\\"include\\",\\"operator\\":\\"or\\",\\"rule\\":[{\\"condition\\":\\"assetSubType\\",\\"ruleOperator\\":\\"include\\",\\"value\\":[{\\"vendor\\":\\"0\\",\\"assetType\\":\\"0\\",\\"assetSubType\\":\\"100\\"}]}]}}"
	AutoConfig *string `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	// The automatic scan configuration type. Valid values:
	//
	// - **0**: automatic scanning is disabled
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
	// 2026-01-09 10:19:57
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

func (s GetCheckScopeConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCheckScopeConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCheckScopeConfigResponseBodyData) GetAutoConfig() *string {
	return s.AutoConfig
}

func (s *GetCheckScopeConfigResponseBodyData) GetAutoType() *int32 {
	return s.AutoType
}

func (s *GetCheckScopeConfigResponseBodyData) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetCheckScopeConfigResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetCheckScopeConfigResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetCheckScopeConfigResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *GetCheckScopeConfigResponseBodyData) SetAutoConfig(v string) *GetCheckScopeConfigResponseBodyData {
	s.AutoConfig = &v
	return s
}

func (s *GetCheckScopeConfigResponseBodyData) SetAutoType(v int32) *GetCheckScopeConfigResponseBodyData {
	s.AutoType = &v
	return s
}

func (s *GetCheckScopeConfigResponseBodyData) SetConfigId(v string) *GetCheckScopeConfigResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *GetCheckScopeConfigResponseBodyData) SetGmtCreate(v string) *GetCheckScopeConfigResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetCheckScopeConfigResponseBodyData) SetGmtModified(v string) *GetCheckScopeConfigResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetCheckScopeConfigResponseBodyData) SetType(v int32) *GetCheckScopeConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetCheckScopeConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
