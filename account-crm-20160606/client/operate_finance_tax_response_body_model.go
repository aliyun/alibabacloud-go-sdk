// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateFinanceTaxResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OperateFinanceTaxResponseBody
	GetCode() *string
	SetData(v string) *OperateFinanceTaxResponseBody
	GetData() *string
	SetMessage(v string) *OperateFinanceTaxResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateFinanceTaxResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateFinanceTaxResponseBody
	GetSuccess() *bool
}

type OperateFinanceTaxResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateFinanceTaxResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateFinanceTaxResponseBody) GoString() string {
	return s.String()
}

func (s *OperateFinanceTaxResponseBody) GetCode() *string {
	return s.Code
}

func (s *OperateFinanceTaxResponseBody) GetData() *string {
	return s.Data
}

func (s *OperateFinanceTaxResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateFinanceTaxResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateFinanceTaxResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateFinanceTaxResponseBody) SetCode(v string) *OperateFinanceTaxResponseBody {
	s.Code = &v
	return s
}

func (s *OperateFinanceTaxResponseBody) SetData(v string) *OperateFinanceTaxResponseBody {
	s.Data = &v
	return s
}

func (s *OperateFinanceTaxResponseBody) SetMessage(v string) *OperateFinanceTaxResponseBody {
	s.Message = &v
	return s
}

func (s *OperateFinanceTaxResponseBody) SetRequestId(v string) *OperateFinanceTaxResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateFinanceTaxResponseBody) SetSuccess(v bool) *OperateFinanceTaxResponseBody {
	s.Success = &v
	return s
}

func (s *OperateFinanceTaxResponseBody) Validate() error {
	return dara.Validate(s)
}
