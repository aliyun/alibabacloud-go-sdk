// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchExpressionVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SwitchExpressionVariableRequest
	GetLang() *string
	SetDataVersion(v int64) *SwitchExpressionVariableRequest
	GetDataVersion() *int64
	SetId(v int64) *SwitchExpressionVariableRequest
	GetId() *int64
	SetRegId(v string) *SwitchExpressionVariableRequest
	GetRegId() *string
	SetStatus(v string) *SwitchExpressionVariableRequest
	GetStatus() *string
}

type SwitchExpressionVariableRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Data version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataVersion *int64 `json:"dataVersion,omitempty" xml:"dataVersion,omitempty"`
	// Variable ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2556
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Status.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SwitchExpressionVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchExpressionVariableRequest) GoString() string {
	return s.String()
}

func (s *SwitchExpressionVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *SwitchExpressionVariableRequest) GetDataVersion() *int64 {
	return s.DataVersion
}

func (s *SwitchExpressionVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *SwitchExpressionVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *SwitchExpressionVariableRequest) GetStatus() *string {
	return s.Status
}

func (s *SwitchExpressionVariableRequest) SetLang(v string) *SwitchExpressionVariableRequest {
	s.Lang = &v
	return s
}

func (s *SwitchExpressionVariableRequest) SetDataVersion(v int64) *SwitchExpressionVariableRequest {
	s.DataVersion = &v
	return s
}

func (s *SwitchExpressionVariableRequest) SetId(v int64) *SwitchExpressionVariableRequest {
	s.Id = &v
	return s
}

func (s *SwitchExpressionVariableRequest) SetRegId(v string) *SwitchExpressionVariableRequest {
	s.RegId = &v
	return s
}

func (s *SwitchExpressionVariableRequest) SetStatus(v string) *SwitchExpressionVariableRequest {
	s.Status = &v
	return s
}

func (s *SwitchExpressionVariableRequest) Validate() error {
	return dara.Validate(s)
}
