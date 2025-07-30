// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelRuleCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DelRuleCategoryResponseBody
	GetCode() *string
	SetData(v *DelRuleCategoryResponseBodyData) *DelRuleCategoryResponseBody
	GetData() *DelRuleCategoryResponseBodyData
	SetMessage(v string) *DelRuleCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DelRuleCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DelRuleCategoryResponseBody
	GetSuccess() *bool
}

type DelRuleCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DelRuleCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelRuleCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DelRuleCategoryResponseBody) GetData() *DelRuleCategoryResponseBodyData {
	return s.Data
}

func (s *DelRuleCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DelRuleCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelRuleCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DelRuleCategoryResponseBody) SetCode(v string) *DelRuleCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *DelRuleCategoryResponseBody) SetData(v *DelRuleCategoryResponseBodyData) *DelRuleCategoryResponseBody {
	s.Data = v
	return s
}

func (s *DelRuleCategoryResponseBody) SetMessage(v string) *DelRuleCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *DelRuleCategoryResponseBody) SetRequestId(v string) *DelRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelRuleCategoryResponseBody) SetSuccess(v bool) *DelRuleCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *DelRuleCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DelRuleCategoryResponseBodyData struct {
	// example:
	//
	// false
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
}

func (s DelRuleCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DelRuleCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DelRuleCategoryResponseBodyData) GetSelect() *bool {
	return s.Select
}

func (s *DelRuleCategoryResponseBodyData) SetSelect(v bool) *DelRuleCategoryResponseBodyData {
	s.Select = &v
	return s
}

func (s *DelRuleCategoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
