// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteCustVariableRequest
	GetLang() *string
	SetCreateType(v string) *DeleteCustVariableRequest
	GetCreateType() *string
	SetDataVersion(v int64) *DeleteCustVariableRequest
	GetDataVersion() *int64
	SetRegId(v string) *DeleteCustVariableRequest
	GetRegId() *string
	SetVariableId(v string) *DeleteCustVariableRequest
	GetVariableId() *string
}

type DeleteCustVariableRequest struct {
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
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Data version.
	//
	// example:
	//
	// 1
	DataVersion *int64 `json:"dataVersion,omitempty" xml:"dataVersion,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Variable ID
	//
	// example:
	//
	// 235
	VariableId *string `json:"variableId,omitempty" xml:"variableId,omitempty"`
}

func (s DeleteCustVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustVariableRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteCustVariableRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DeleteCustVariableRequest) GetDataVersion() *int64 {
	return s.DataVersion
}

func (s *DeleteCustVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteCustVariableRequest) GetVariableId() *string {
	return s.VariableId
}

func (s *DeleteCustVariableRequest) SetLang(v string) *DeleteCustVariableRequest {
	s.Lang = &v
	return s
}

func (s *DeleteCustVariableRequest) SetCreateType(v string) *DeleteCustVariableRequest {
	s.CreateType = &v
	return s
}

func (s *DeleteCustVariableRequest) SetDataVersion(v int64) *DeleteCustVariableRequest {
	s.DataVersion = &v
	return s
}

func (s *DeleteCustVariableRequest) SetRegId(v string) *DeleteCustVariableRequest {
	s.RegId = &v
	return s
}

func (s *DeleteCustVariableRequest) SetVariableId(v string) *DeleteCustVariableRequest {
	s.VariableId = &v
	return s
}

func (s *DeleteCustVariableRequest) Validate() error {
	return dara.Validate(s)
}
