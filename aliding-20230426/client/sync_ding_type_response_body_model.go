// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDingTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SyncDingTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncDingTypeResponseBody
	GetSuccess() *bool
	SetVendorRequestId(v string) *SyncDingTypeResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SyncDingTypeResponseBody
	GetVendorType() *string
}

type SyncDingTypeResponseBody struct {
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

func (s SyncDingTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncDingTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDingTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncDingTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncDingTypeResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SyncDingTypeResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SyncDingTypeResponseBody) SetRequestId(v string) *SyncDingTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDingTypeResponseBody) SetSuccess(v bool) *SyncDingTypeResponseBody {
	s.Success = &v
	return s
}

func (s *SyncDingTypeResponseBody) SetVendorRequestId(v string) *SyncDingTypeResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SyncDingTypeResponseBody) SetVendorType(v string) *SyncDingTypeResponseBody {
	s.VendorType = &v
	return s
}

func (s *SyncDingTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
