// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPlaybookInputOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExeConfig(v string) *ModifyPlaybookInputOutputRequest
	GetExeConfig() *string
	SetInputParams(v string) *ModifyPlaybookInputOutputRequest
	GetInputParams() *string
	SetLang(v string) *ModifyPlaybookInputOutputRequest
	GetLang() *string
	SetOutputParams(v string) *ModifyPlaybookInputOutputRequest
	GetOutputParams() *string
	SetParamType(v string) *ModifyPlaybookInputOutputRequest
	GetParamType() *string
	SetPlaybookUuid(v string) *ModifyPlaybookInputOutputRequest
	GetPlaybookUuid() *string
}

type ModifyPlaybookInputOutputRequest struct {
	// The execution method for the playbook. This parameter is in the JSONObject format.
	ExeConfig *string `json:"ExeConfig,omitempty" xml:"ExeConfig,omitempty"`
	// The input parameter configuration for the playbook. This parameter is in the JSONArray format.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "typeName": "String",
	//
	//         "dataClass": "normal",
	//
	//         "dataType": "String",
	//
	//         "description": "period",
	//
	//         "example": "",
	//
	//         "name": "period",
	//
	//         "required": false
	//
	//     }
	//
	// ]
	InputParams *string `json:"InputParams,omitempty" xml:"InputParams,omitempty"`
	// The language of the request and response messages.
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Playbooks do not support output parameter configurations. This parameter is fixed to an empty value.
	//
	// This parameter is required.
	//
	// example:
	//
	// []
	OutputParams *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
	// The type of the input parameter for the playbook.
	//
	// - **template-ip**: IP request template.
	//
	// - **template-file**: file request template.
	//
	// - **template-process**: process request template.
	//
	// - **custom**: custom parameter.
	//
	// example:
	//
	// custom
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The UUID of the playbook.
	//
	// > Call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8baa6cff-319e-4ede-97bc-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s ModifyPlaybookInputOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPlaybookInputOutputRequest) GoString() string {
	return s.String()
}

func (s *ModifyPlaybookInputOutputRequest) GetExeConfig() *string {
	return s.ExeConfig
}

func (s *ModifyPlaybookInputOutputRequest) GetInputParams() *string {
	return s.InputParams
}

func (s *ModifyPlaybookInputOutputRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyPlaybookInputOutputRequest) GetOutputParams() *string {
	return s.OutputParams
}

func (s *ModifyPlaybookInputOutputRequest) GetParamType() *string {
	return s.ParamType
}

func (s *ModifyPlaybookInputOutputRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ModifyPlaybookInputOutputRequest) SetExeConfig(v string) *ModifyPlaybookInputOutputRequest {
	s.ExeConfig = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetInputParams(v string) *ModifyPlaybookInputOutputRequest {
	s.InputParams = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetLang(v string) *ModifyPlaybookInputOutputRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetOutputParams(v string) *ModifyPlaybookInputOutputRequest {
	s.OutputParams = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetParamType(v string) *ModifyPlaybookInputOutputRequest {
	s.ParamType = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) SetPlaybookUuid(v string) *ModifyPlaybookInputOutputRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ModifyPlaybookInputOutputRequest) Validate() error {
	return dara.Validate(s)
}
