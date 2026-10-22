// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAutoAuditUnblockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitAutoAuditUnblockResponseBody
	GetCode() *string
	SetData(v bool) *SubmitAutoAuditUnblockResponseBody
	GetData() *bool
	SetMessage(v string) *SubmitAutoAuditUnblockResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitAutoAuditUnblockResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitAutoAuditUnblockResponseBody
	GetSuccess() *bool
}

type SubmitAutoAuditUnblockResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6B57D35D-9DAC-5393-AE39-07697E37C2E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAutoAuditUnblockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAutoAuditUnblockResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAutoAuditUnblockResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitAutoAuditUnblockResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitAutoAuditUnblockResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitAutoAuditUnblockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAutoAuditUnblockResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAutoAuditUnblockResponseBody) SetCode(v string) *SubmitAutoAuditUnblockResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAutoAuditUnblockResponseBody) SetData(v bool) *SubmitAutoAuditUnblockResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitAutoAuditUnblockResponseBody) SetMessage(v string) *SubmitAutoAuditUnblockResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAutoAuditUnblockResponseBody) SetRequestId(v string) *SubmitAutoAuditUnblockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAutoAuditUnblockResponseBody) SetSuccess(v bool) *SubmitAutoAuditUnblockResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitAutoAuditUnblockResponseBody) Validate() error {
	return dara.Validate(s)
}
