// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseVideoConferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCause(v string) *CloseVideoConferenceResponseBody
	GetCause() *string
	SetCode(v int64) *CloseVideoConferenceResponseBody
	GetCode() *int64
	SetRequestId(v string) *CloseVideoConferenceResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *CloseVideoConferenceResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CloseVideoConferenceResponseBody
	GetVendorType() *string
}

type CloseVideoConferenceResponseBody struct {
	// example:
	//
	// success
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
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

func (s CloseVideoConferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseVideoConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceResponseBody) GetCause() *string {
	return s.Cause
}

func (s *CloseVideoConferenceResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CloseVideoConferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseVideoConferenceResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CloseVideoConferenceResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CloseVideoConferenceResponseBody) SetCause(v string) *CloseVideoConferenceResponseBody {
	s.Cause = &v
	return s
}

func (s *CloseVideoConferenceResponseBody) SetCode(v int64) *CloseVideoConferenceResponseBody {
	s.Code = &v
	return s
}

func (s *CloseVideoConferenceResponseBody) SetRequestId(v string) *CloseVideoConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseVideoConferenceResponseBody) SetVendorRequestId(v string) *CloseVideoConferenceResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CloseVideoConferenceResponseBody) SetVendorType(v string) *CloseVideoConferenceResponseBody {
	s.VendorType = &v
	return s
}

func (s *CloseVideoConferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
