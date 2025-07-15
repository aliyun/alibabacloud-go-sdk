// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePostPayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpgradePostPayOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *UpgradePostPayOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradePostPayOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradePostPayOrderResponseBody
	GetSuccess() *bool
}

type UpgradePostPayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015A***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradePostPayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradePostPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpgradePostPayOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradePostPayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradePostPayOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradePostPayOrderResponseBody) SetCode(v int32) *UpgradePostPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradePostPayOrderResponseBody) SetMessage(v string) *UpgradePostPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradePostPayOrderResponseBody) SetRequestId(v string) *UpgradePostPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradePostPayOrderResponseBody) SetSuccess(v bool) *UpgradePostPayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradePostPayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
