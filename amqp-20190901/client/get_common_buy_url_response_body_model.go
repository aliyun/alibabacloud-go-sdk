// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommonBuyUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetCommonBuyUrlResponseBody
	GetCode() *int32
	SetData(v string) *GetCommonBuyUrlResponseBody
	GetData() *string
	SetMessage(v string) *GetCommonBuyUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCommonBuyUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCommonBuyUrlResponseBody
	GetSuccess() *bool
}

type GetCommonBuyUrlResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCommonBuyUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCommonBuyUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommonBuyUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetCommonBuyUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetCommonBuyUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCommonBuyUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCommonBuyUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCommonBuyUrlResponseBody) SetCode(v int32) *GetCommonBuyUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetCommonBuyUrlResponseBody) SetData(v string) *GetCommonBuyUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetCommonBuyUrlResponseBody) SetMessage(v string) *GetCommonBuyUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetCommonBuyUrlResponseBody) SetRequestId(v string) *GetCommonBuyUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommonBuyUrlResponseBody) SetSuccess(v bool) *GetCommonBuyUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetCommonBuyUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
