// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventVariableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventVariableListRequest
	GetLang() *string
	SetCreateType(v string) *DescribeEventVariableListRequest
	GetCreateType() *string
	SetFilterDTO(v string) *DescribeEventVariableListRequest
	GetFilterDTO() *string
	SetRefObjId(v string) *DescribeEventVariableListRequest
	GetRefObjId() *string
	SetRefObjType(v string) *DescribeEventVariableListRequest
	GetRefObjType() *string
	SetRegId(v string) *DescribeEventVariableListRequest
	GetRegId() *string
	SetType(v string) *DescribeEventVariableListRequest
	GetType() *string
}

type DescribeEventVariableListRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The creation type.
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// The filter object.
	//
	// example:
	//
	// {
	//
	//     "type": "EXPRESSION",
	//
	//     "name": "ex_NgR6nDVD821c"
	//
	// }
	FilterDTO *string `json:"filterDTO,omitempty" xml:"filterDTO,omitempty"`
	// The associated event code (eventCode).
	//
	// This parameter is required.
	//
	// example:
	//
	// de_awkhwh0314
	RefObjId *string `json:"refObjId,omitempty" xml:"refObjId,omitempty"`
	// The association type.
	//
	// This parameter is required.
	//
	// example:
	//
	// EVENT
	RefObjType *string `json:"refObjType,omitempty" xml:"refObjType,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The type.
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeEventVariableListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventVariableListRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeEventVariableListRequest) GetFilterDTO() *string {
	return s.FilterDTO
}

func (s *DescribeEventVariableListRequest) GetRefObjId() *string {
	return s.RefObjId
}

func (s *DescribeEventVariableListRequest) GetRefObjType() *string {
	return s.RefObjType
}

func (s *DescribeEventVariableListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventVariableListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListRequest) SetLang(v string) *DescribeEventVariableListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventVariableListRequest) SetCreateType(v string) *DescribeEventVariableListRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeEventVariableListRequest) SetFilterDTO(v string) *DescribeEventVariableListRequest {
	s.FilterDTO = &v
	return s
}

func (s *DescribeEventVariableListRequest) SetRefObjId(v string) *DescribeEventVariableListRequest {
	s.RefObjId = &v
	return s
}

func (s *DescribeEventVariableListRequest) SetRefObjType(v string) *DescribeEventVariableListRequest {
	s.RefObjType = &v
	return s
}

func (s *DescribeEventVariableListRequest) SetRegId(v string) *DescribeEventVariableListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventVariableListRequest) SetType(v string) *DescribeEventVariableListRequest {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListRequest) Validate() error {
	return dara.Validate(s)
}
