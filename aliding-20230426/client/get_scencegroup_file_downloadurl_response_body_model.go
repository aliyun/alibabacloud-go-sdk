// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScencegroupFileDownloadurlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *GetScencegroupFileDownloadurlResponseBody
	GetDownloadUrl() *string
	SetRequestId(v string) *GetScencegroupFileDownloadurlResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetScencegroupFileDownloadurlResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetScencegroupFileDownloadurlResponseBody
	GetVendorType() *string
}

type GetScencegroupFileDownloadurlResponseBody struct {
	// example:
	//
	// An https download connection
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
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

func (s GetScencegroupFileDownloadurlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScencegroupFileDownloadurlResponseBody) GoString() string {
	return s.String()
}

func (s *GetScencegroupFileDownloadurlResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetScencegroupFileDownloadurlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScencegroupFileDownloadurlResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetScencegroupFileDownloadurlResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetScencegroupFileDownloadurlResponseBody) SetDownloadUrl(v string) *GetScencegroupFileDownloadurlResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *GetScencegroupFileDownloadurlResponseBody) SetRequestId(v string) *GetScencegroupFileDownloadurlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScencegroupFileDownloadurlResponseBody) SetVendorRequestId(v string) *GetScencegroupFileDownloadurlResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetScencegroupFileDownloadurlResponseBody) SetVendorType(v string) *GetScencegroupFileDownloadurlResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetScencegroupFileDownloadurlResponseBody) Validate() error {
	return dara.Validate(s)
}
