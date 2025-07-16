// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlidingAssistantInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalStatus(v int32) *GetAlidingAssistantInfoResponseBody
	GetApprovalStatus() *int32
	SetProcessInstanceId(v string) *GetAlidingAssistantInfoResponseBody
	GetProcessInstanceId() *string
	SetRequestId(v string) *GetAlidingAssistantInfoResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetAlidingAssistantInfoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetAlidingAssistantInfoResponseBody
	GetVendorType() *string
}

type GetAlidingAssistantInfoResponseBody struct {
	// example:
	//
	// 1
	ApprovalStatus *int32 `json:"approvalStatus,omitempty" xml:"approvalStatus,omitempty"`
	// example:
	//
	// 123
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetAlidingAssistantInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlidingAssistantInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlidingAssistantInfoResponseBody) GetApprovalStatus() *int32 {
	return s.ApprovalStatus
}

func (s *GetAlidingAssistantInfoResponseBody) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetAlidingAssistantInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlidingAssistantInfoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetAlidingAssistantInfoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetAlidingAssistantInfoResponseBody) SetApprovalStatus(v int32) *GetAlidingAssistantInfoResponseBody {
	s.ApprovalStatus = &v
	return s
}

func (s *GetAlidingAssistantInfoResponseBody) SetProcessInstanceId(v string) *GetAlidingAssistantInfoResponseBody {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetAlidingAssistantInfoResponseBody) SetRequestId(v string) *GetAlidingAssistantInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlidingAssistantInfoResponseBody) SetVendorRequestId(v string) *GetAlidingAssistantInfoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetAlidingAssistantInfoResponseBody) SetVendorType(v string) *GetAlidingAssistantInfoResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetAlidingAssistantInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
