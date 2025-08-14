// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyCustVariableRequest
	GetLang() *string
	SetCondition(v string) *ModifyCustVariableRequest
	GetCondition() *string
	SetDataVersion(v int64) *ModifyCustVariableRequest
	GetDataVersion() *int64
	SetDescription(v string) *ModifyCustVariableRequest
	GetDescription() *string
	SetEventCodes(v string) *ModifyCustVariableRequest
	GetEventCodes() *string
	SetId(v int64) *ModifyCustVariableRequest
	GetId() *int64
	SetName(v string) *ModifyCustVariableRequest
	GetName() *string
	SetOutputs(v string) *ModifyCustVariableRequest
	GetOutputs() *string
	SetRegId(v string) *ModifyCustVariableRequest
	GetRegId() *string
}

type ModifyCustVariableRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Condition value.
	//
	// example:
	//
	// {\\"currentId\\":1,\\"deepCount\\":0,\\"list\\":[{\\"currentId\\":2,\\"deepCount\\":1,\\"left\\":{\\"code\\":\\"__hit_rules\\",\\"fieldType\\":\\"STRING\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"hitRules\\",\\"operatorCode\\":\\"deNotInNameList\\"},\\"operatorCode\\":\\"deNotInNameList\\",\\"parentId\\":1,\\"right\\":{\\"code\\":\\"\\"nl_Xcufc8wV6624\\"\\",\\"name\\":\\"nl_Xcufc8wV6624\\",\\"rightVariableType\\":\\"variable\\"}},{\\"currentId\\":3,\\"deepCount\\":1,\\"list\\":[{\\"currentId\\":4,\\"deepCount\\":2,\\"left\\":{\\"code\\":\\"__v_safde\\",\\"fieldType\\":\\"DOUBLE\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":true,\\"name\\":\\"score\\",\\"operatorCode\\":\\"equals\\"},\\"operatorCode\\":\\"equals\\",\\"parentId\\":3,\\"right\\":{\\"code\\":\\"\\"99\\"\\",\\"name\\":\\"99\\",\\"rightVariableType\\":\\"constant\\"}},{\\"currentId\\":5,\\"deepCount\\":2,\\"left\\":{\\"code\\":\\"queryPhoneSimulatorInfo(deviceToken)?\\",\\"fieldType\\":\\"STRING\\",\\"functionName\\":\\"\\",\\"hasRightVariable\\":false,\\"name\\":\\"__device-test01__\\",\\"operatorCode\\":\\"isNotEmptyWrapper\\"},\\"operatorCode\\":\\"isNotEmptyWrapper\\",\\"parentId\\":3,\\"right\\":{\\"functionName\\":\\"\\",\\"name\\":\\"\\",\\"rightVariableType\\":\\"constant\\"}}],\\"parentId\\":1,\\"relationship\\":\\"and\\"}],\\"parentId\\":0,\\"relationship\\":\\"and\\"}
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Data version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	DataVersion *int64 `json:"dataVersion,omitempty" xml:"dataVersion,omitempty"`
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code.
	//
	// example:
	//
	// account_abuse_pro,account_abuse
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Variable ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 376773
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Variable name
	//
	// This parameter is required.
	//
	// example:
	//
	// dxkkLw8fe18
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Output
	//
	// example:
	//
	// STRING
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s ModifyCustVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustVariableRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyCustVariableRequest) GetCondition() *string {
	return s.Condition
}

func (s *ModifyCustVariableRequest) GetDataVersion() *int64 {
	return s.DataVersion
}

func (s *ModifyCustVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCustVariableRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *ModifyCustVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyCustVariableRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCustVariableRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *ModifyCustVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModifyCustVariableRequest) SetLang(v string) *ModifyCustVariableRequest {
	s.Lang = &v
	return s
}

func (s *ModifyCustVariableRequest) SetCondition(v string) *ModifyCustVariableRequest {
	s.Condition = &v
	return s
}

func (s *ModifyCustVariableRequest) SetDataVersion(v int64) *ModifyCustVariableRequest {
	s.DataVersion = &v
	return s
}

func (s *ModifyCustVariableRequest) SetDescription(v string) *ModifyCustVariableRequest {
	s.Description = &v
	return s
}

func (s *ModifyCustVariableRequest) SetEventCodes(v string) *ModifyCustVariableRequest {
	s.EventCodes = &v
	return s
}

func (s *ModifyCustVariableRequest) SetId(v int64) *ModifyCustVariableRequest {
	s.Id = &v
	return s
}

func (s *ModifyCustVariableRequest) SetName(v string) *ModifyCustVariableRequest {
	s.Name = &v
	return s
}

func (s *ModifyCustVariableRequest) SetOutputs(v string) *ModifyCustVariableRequest {
	s.Outputs = &v
	return s
}

func (s *ModifyCustVariableRequest) SetRegId(v string) *ModifyCustVariableRequest {
	s.RegId = &v
	return s
}

func (s *ModifyCustVariableRequest) Validate() error {
	return dara.Validate(s)
}
