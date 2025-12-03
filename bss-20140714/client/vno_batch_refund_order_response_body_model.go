// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVnoBatchRefundOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VnoBatchRefundOrderResponseBody
	GetCode() *string
	SetData(v string) *VnoBatchRefundOrderResponseBody
	GetData() *string
	SetMessage(v string) *VnoBatchRefundOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *VnoBatchRefundOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VnoBatchRefundOrderResponseBody
	GetSuccess() *bool
}

type VnoBatchRefundOrderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VnoBatchRefundOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VnoBatchRefundOrderResponseBody) GoString() string {
	return s.String()
}

func (s *VnoBatchRefundOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *VnoBatchRefundOrderResponseBody) GetData() *string {
	return s.Data
}

func (s *VnoBatchRefundOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VnoBatchRefundOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VnoBatchRefundOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VnoBatchRefundOrderResponseBody) SetCode(v string) *VnoBatchRefundOrderResponseBody {
	s.Code = &v
	return s
}

func (s *VnoBatchRefundOrderResponseBody) SetData(v string) *VnoBatchRefundOrderResponseBody {
	s.Data = &v
	return s
}

func (s *VnoBatchRefundOrderResponseBody) SetMessage(v string) *VnoBatchRefundOrderResponseBody {
	s.Message = &v
	return s
}

func (s *VnoBatchRefundOrderResponseBody) SetRequestId(v string) *VnoBatchRefundOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *VnoBatchRefundOrderResponseBody) SetSuccess(v bool) *VnoBatchRefundOrderResponseBody {
	s.Success = &v
	return s
}

func (s *VnoBatchRefundOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
