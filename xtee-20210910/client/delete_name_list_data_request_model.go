// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNameListDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteNameListDataRequest
	GetLang() *string
	SetRegId(v string) *DeleteNameListDataRequest
	GetRegId() *string
	SetVariableId(v string) *DeleteNameListDataRequest
	GetVariableId() *string
}

type DeleteNameListDataRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// 239
	VariableId *string `json:"variableId,omitempty" xml:"variableId,omitempty"`
}

func (s DeleteNameListDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNameListDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteNameListDataRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteNameListDataRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteNameListDataRequest) GetVariableId() *string {
	return s.VariableId
}

func (s *DeleteNameListDataRequest) SetLang(v string) *DeleteNameListDataRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNameListDataRequest) SetRegId(v string) *DeleteNameListDataRequest {
	s.RegId = &v
	return s
}

func (s *DeleteNameListDataRequest) SetVariableId(v string) *DeleteNameListDataRequest {
	s.VariableId = &v
	return s
}

func (s *DeleteNameListDataRequest) Validate() error {
	return dara.Validate(s)
}
