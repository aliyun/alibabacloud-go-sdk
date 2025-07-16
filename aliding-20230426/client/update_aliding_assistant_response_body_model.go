// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlidingAssistantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAlidingAssistantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAlidingAssistantResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *UpdateAlidingAssistantResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateAlidingAssistantResponseBody
	GetVendorType() *string
}

type UpdateAlidingAssistantResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s UpdateAlidingAssistantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlidingAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlidingAssistantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlidingAssistantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAlidingAssistantResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateAlidingAssistantResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateAlidingAssistantResponseBody) SetRequestId(v string) *UpdateAlidingAssistantResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlidingAssistantResponseBody) SetSuccess(v bool) *UpdateAlidingAssistantResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAlidingAssistantResponseBody) SetVendorRequestId(v string) *UpdateAlidingAssistantResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateAlidingAssistantResponseBody) SetVendorType(v string) *UpdateAlidingAssistantResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateAlidingAssistantResponseBody) Validate() error {
	return dara.Validate(s)
}
