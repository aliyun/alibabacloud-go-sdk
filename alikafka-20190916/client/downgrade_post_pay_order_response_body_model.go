// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradePostPayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DowngradePostPayOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *DowngradePostPayOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *DowngradePostPayOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DowngradePostPayOrderResponseBody
	GetSuccess() *bool
}

type DowngradePostPayOrderResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DowngradePostPayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DowngradePostPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DowngradePostPayOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DowngradePostPayOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DowngradePostPayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DowngradePostPayOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DowngradePostPayOrderResponseBody) SetCode(v int32) *DowngradePostPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *DowngradePostPayOrderResponseBody) SetMessage(v string) *DowngradePostPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *DowngradePostPayOrderResponseBody) SetRequestId(v string) *DowngradePostPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DowngradePostPayOrderResponseBody) SetSuccess(v bool) *DowngradePostPayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *DowngradePostPayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
