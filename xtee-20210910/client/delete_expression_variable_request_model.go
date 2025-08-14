// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressionVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteExpressionVariableRequest
	GetLang() *string
	SetDataVersion(v int64) *DeleteExpressionVariableRequest
	GetDataVersion() *int64
	SetId(v int64) *DeleteExpressionVariableRequest
	GetId() *int64
	SetRegId(v string) *DeleteExpressionVariableRequest
	GetRegId() *string
}

type DeleteExpressionVariableRequest struct {
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
}

func (s DeleteExpressionVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressionVariableRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressionVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteExpressionVariableRequest) GetDataVersion() *int64 {
	return s.DataVersion
}

func (s *DeleteExpressionVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteExpressionVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteExpressionVariableRequest) SetLang(v string) *DeleteExpressionVariableRequest {
	s.Lang = &v
	return s
}

func (s *DeleteExpressionVariableRequest) SetDataVersion(v int64) *DeleteExpressionVariableRequest {
	s.DataVersion = &v
	return s
}

func (s *DeleteExpressionVariableRequest) SetId(v int64) *DeleteExpressionVariableRequest {
	s.Id = &v
	return s
}

func (s *DeleteExpressionVariableRequest) SetRegId(v string) *DeleteExpressionVariableRequest {
	s.RegId = &v
	return s
}

func (s *DeleteExpressionVariableRequest) Validate() error {
	return dara.Validate(s)
}
