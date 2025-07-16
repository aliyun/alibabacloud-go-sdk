// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDingTalkIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeDingTalkIdResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *ChangeDingTalkIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ChangeDingTalkIdResponseBody
	GetVendorType() *string
}

type ChangeDingTalkIdResponseBody struct {
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

func (s ChangeDingTalkIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeDingTalkIdResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDingTalkIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeDingTalkIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ChangeDingTalkIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ChangeDingTalkIdResponseBody) SetRequestId(v string) *ChangeDingTalkIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDingTalkIdResponseBody) SetVendorRequestId(v string) *ChangeDingTalkIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ChangeDingTalkIdResponseBody) SetVendorType(v string) *ChangeDingTalkIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *ChangeDingTalkIdResponseBody) Validate() error {
	return dara.Validate(s)
}
