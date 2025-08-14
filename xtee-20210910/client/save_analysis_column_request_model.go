// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAnalysisColumnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SaveAnalysisColumnRequest
	GetLang() *string
	SetColumns(v string) *SaveAnalysisColumnRequest
	GetColumns() *string
	SetRegId(v string) *SaveAnalysisColumnRequest
	GetRegId() *string
}

type SaveAnalysisColumnRequest struct {
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
	// Custom columns
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"variableName\\":\\"requestId\\",\\"variableTitle\\":\\"RequestId\\",\\"variableType\\":\\"SYSTEM_BIND\\",\\"isDefault\\":true},{\\"variableName\\":\\"eventTime\\",\\"variableTitle\\":\\"事件时间\\",\\"variableType\\":\\"SYSTEM_BIND\\",\\"isDefault\\":true},{\\"variableName\\":\\"accountId\\",\\"variableTitle\\":\\"账号\\",\\"variableType\\":\\"SYSTEM_BIND\\",\\"isDefault\\":true},{\\"variableName\\":\\"deviceId\\",\\"variableTitle\\":\\"设备ID\\",\\"variableType\\":\\"SYSTEM_BIND\\",\\"isDefault\\":true},{\\"variableName\\":\\"eventCode\\",\\"variableTitle\\":\\"事件编码\\",\\"variableType\\":\\"SYSTEM_BIND\\",\\"isDefault\\":true},{\\"variableName\\":\\"ip\\",\\"variableTitle\\":\\"IP\\",\\"variableType\\":\\"SYSTEM_BIND\\",\\"isDefault\\":true},{\\"variableName\\":\\"score\\",\\"variableTitle\\":\\"分值\\",\\"variableType\\":\\"SYSTEM_BIND\\",\\"isDefault\\":true},{\\"variableName\\":\\"tags\\",\\"variableTitle\\":\\"标签\\",\\"variableType\\":\\"SYSTEM_BIND\\",\\"isDefault\\":true},{\\"variableName\\":\\"DEtest222\\",\\"variableTitle\\":\\"测试222\\",\\"variableType\\":\\"NATIVE\\",\\"isDefault\\":false}]
	Columns *string `json:"columns,omitempty" xml:"columns,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s SaveAnalysisColumnRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAnalysisColumnRequest) GoString() string {
	return s.String()
}

func (s *SaveAnalysisColumnRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveAnalysisColumnRequest) GetColumns() *string {
	return s.Columns
}

func (s *SaveAnalysisColumnRequest) GetRegId() *string {
	return s.RegId
}

func (s *SaveAnalysisColumnRequest) SetLang(v string) *SaveAnalysisColumnRequest {
	s.Lang = &v
	return s
}

func (s *SaveAnalysisColumnRequest) SetColumns(v string) *SaveAnalysisColumnRequest {
	s.Columns = &v
	return s
}

func (s *SaveAnalysisColumnRequest) SetRegId(v string) *SaveAnalysisColumnRequest {
	s.RegId = &v
	return s
}

func (s *SaveAnalysisColumnRequest) Validate() error {
	return dara.Validate(s)
}
