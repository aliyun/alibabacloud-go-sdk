// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSceneGroupTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableSceneGroupTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableSceneGroupTemplateResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *DisableSceneGroupTemplateResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *DisableSceneGroupTemplateResponseBody
	GetVendorType() *string
}

type DisableSceneGroupTemplateResponseBody struct {
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

func (s DisableSceneGroupTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneGroupTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSceneGroupTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSceneGroupTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableSceneGroupTemplateResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *DisableSceneGroupTemplateResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *DisableSceneGroupTemplateResponseBody) SetRequestId(v string) *DisableSceneGroupTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSceneGroupTemplateResponseBody) SetSuccess(v bool) *DisableSceneGroupTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DisableSceneGroupTemplateResponseBody) SetVendorRequestId(v string) *DisableSceneGroupTemplateResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *DisableSceneGroupTemplateResponseBody) SetVendorType(v string) *DisableSceneGroupTemplateResponseBody {
	s.VendorType = &v
	return s
}

func (s *DisableSceneGroupTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
