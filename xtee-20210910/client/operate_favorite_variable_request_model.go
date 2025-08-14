// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateFavoriteVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *OperateFavoriteVariableRequest
	GetLang() *string
	SetId(v int64) *OperateFavoriteVariableRequest
	GetId() *int64
	SetOperate(v string) *OperateFavoriteVariableRequest
	GetOperate() *string
	SetRegId(v string) *OperateFavoriteVariableRequest
	GetRegId() *string
}

type OperateFavoriteVariableRequest struct {
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
	// Variable ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Favorite operation
	//
	// This parameter is required.
	//
	// example:
	//
	// ADD
	Operate *string `json:"operate,omitempty" xml:"operate,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s OperateFavoriteVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateFavoriteVariableRequest) GoString() string {
	return s.String()
}

func (s *OperateFavoriteVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *OperateFavoriteVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *OperateFavoriteVariableRequest) GetOperate() *string {
	return s.Operate
}

func (s *OperateFavoriteVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *OperateFavoriteVariableRequest) SetLang(v string) *OperateFavoriteVariableRequest {
	s.Lang = &v
	return s
}

func (s *OperateFavoriteVariableRequest) SetId(v int64) *OperateFavoriteVariableRequest {
	s.Id = &v
	return s
}

func (s *OperateFavoriteVariableRequest) SetOperate(v string) *OperateFavoriteVariableRequest {
	s.Operate = &v
	return s
}

func (s *OperateFavoriteVariableRequest) SetRegId(v string) *OperateFavoriteVariableRequest {
	s.RegId = &v
	return s
}

func (s *OperateFavoriteVariableRequest) Validate() error {
	return dara.Validate(s)
}
