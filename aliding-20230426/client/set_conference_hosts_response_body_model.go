// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetConferenceHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetConferenceHostsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetConferenceHostsResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *SetConferenceHostsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SetConferenceHostsResponseBody
	GetVendorType() *string
}

type SetConferenceHostsResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SetConferenceHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetConferenceHostsResponseBody) GoString() string {
	return s.String()
}

func (s *SetConferenceHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetConferenceHostsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetConferenceHostsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SetConferenceHostsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SetConferenceHostsResponseBody) SetRequestId(v string) *SetConferenceHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetConferenceHostsResponseBody) SetSuccess(v bool) *SetConferenceHostsResponseBody {
	s.Success = &v
	return s
}

func (s *SetConferenceHostsResponseBody) SetVendorRequestId(v string) *SetConferenceHostsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SetConferenceHostsResponseBody) SetVendorType(v string) *SetConferenceHostsResponseBody {
	s.VendorType = &v
	return s
}

func (s *SetConferenceHostsResponseBody) Validate() error {
	return dara.Validate(s)
}
