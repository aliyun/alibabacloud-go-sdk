// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBizCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyBizCategoryResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyBizCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyBizCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBizCategoryResponseBody
	GetSuccess() *bool
}

type ModifyBizCategoryResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyBizCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBizCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBizCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyBizCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyBizCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBizCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBizCategoryResponseBody) SetCode(v string) *ModifyBizCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBizCategoryResponseBody) SetMessage(v string) *ModifyBizCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBizCategoryResponseBody) SetRequestId(v string) *ModifyBizCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBizCategoryResponseBody) SetSuccess(v bool) *ModifyBizCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBizCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
