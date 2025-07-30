// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRuleCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddRuleCategoryResponseBody
	GetCode() *string
	SetData(v *AddRuleCategoryResponseBodyData) *AddRuleCategoryResponseBody
	GetData() *AddRuleCategoryResponseBodyData
	SetMessage(v string) *AddRuleCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddRuleCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRuleCategoryResponseBody
	GetSuccess() *bool
}

type AddRuleCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRuleCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddRuleCategoryResponseBody) GetData() *AddRuleCategoryResponseBodyData {
	return s.Data
}

func (s *AddRuleCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddRuleCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRuleCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRuleCategoryResponseBody) SetCode(v string) *AddRuleCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *AddRuleCategoryResponseBody) SetData(v *AddRuleCategoryResponseBodyData) *AddRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *AddRuleCategoryResponseBody) SetMessage(v string) *AddRuleCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *AddRuleCategoryResponseBody) SetRequestId(v string) *AddRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRuleCategoryResponseBody) SetSuccess(v bool) *AddRuleCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *AddRuleCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddRuleCategoryResponseBodyData struct {
	Select *bool  `json:"Select,omitempty" xml:"Select,omitempty"`
	Type   *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddRuleCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddRuleCategoryResponseBodyData) GetSelect() *bool {
	return s.Select
}

func (s *AddRuleCategoryResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *AddRuleCategoryResponseBodyData) SetSelect(v bool) *AddRuleCategoryResponseBodyData {
	s.Select = &v
	return s
}

func (s *AddRuleCategoryResponseBodyData) SetType(v int32) *AddRuleCategoryResponseBodyData {
	s.Type = &v
	return s
}

func (s *AddRuleCategoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
