// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryForCssOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryForCssOrderResponseBody
	GetCode() *string
	SetData(v string) *QueryForCssOrderResponseBody
	GetData() *string
	SetMessage(v string) *QueryForCssOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryForCssOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryForCssOrderResponseBody
	GetSuccess() *bool
}

type QueryForCssOrderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryForCssOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryForCssOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryForCssOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryForCssOrderResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryForCssOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryForCssOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryForCssOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryForCssOrderResponseBody) SetCode(v string) *QueryForCssOrderResponseBody {
	s.Code = &v
	return s
}

func (s *QueryForCssOrderResponseBody) SetData(v string) *QueryForCssOrderResponseBody {
	s.Data = &v
	return s
}

func (s *QueryForCssOrderResponseBody) SetMessage(v string) *QueryForCssOrderResponseBody {
	s.Message = &v
	return s
}

func (s *QueryForCssOrderResponseBody) SetRequestId(v string) *QueryForCssOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryForCssOrderResponseBody) SetSuccess(v bool) *QueryForCssOrderResponseBody {
	s.Success = &v
	return s
}

func (s *QueryForCssOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
