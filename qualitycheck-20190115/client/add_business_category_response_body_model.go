// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBusinessCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddBusinessCategoryResponseBody
	GetCode() *string
	SetData(v string) *AddBusinessCategoryResponseBody
	GetData() *string
	SetMessage(v string) *AddBusinessCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddBusinessCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddBusinessCategoryResponseBody
	GetSuccess() *bool
}

type AddBusinessCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 348193421
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 76DB5D8C-5BD9-42A7-B527-5AF3A5F83F12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddBusinessCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBusinessCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddBusinessCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddBusinessCategoryResponseBody) GetData() *string {
	return s.Data
}

func (s *AddBusinessCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddBusinessCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBusinessCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddBusinessCategoryResponseBody) SetCode(v string) *AddBusinessCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetData(v string) *AddBusinessCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetMessage(v string) *AddBusinessCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetRequestId(v string) *AddBusinessCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) SetSuccess(v bool) *AddBusinessCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *AddBusinessCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
