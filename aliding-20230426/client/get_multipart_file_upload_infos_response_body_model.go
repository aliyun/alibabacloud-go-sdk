// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipartFileUploadInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMultipartHeaderSignatureInfos(v []*GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) *GetMultipartFileUploadInfosResponseBody
	GetMultipartHeaderSignatureInfos() []*GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos
	SetRequestId(v string) *GetMultipartFileUploadInfosResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetMultipartFileUploadInfosResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetMultipartFileUploadInfosResponseBody
	GetVendorType() *string
}

type GetMultipartFileUploadInfosResponseBody struct {
	MultipartHeaderSignatureInfos []*GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos `json:"multipartHeaderSignatureInfos,omitempty" xml:"multipartHeaderSignatureInfos,omitempty" type:"Repeated"`
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

func (s GetMultipartFileUploadInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosResponseBody) GetMultipartHeaderSignatureInfos() []*GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos {
	return s.MultipartHeaderSignatureInfos
}

func (s *GetMultipartFileUploadInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultipartFileUploadInfosResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetMultipartFileUploadInfosResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetMultipartFileUploadInfosResponseBody) SetMultipartHeaderSignatureInfos(v []*GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) *GetMultipartFileUploadInfosResponseBody {
	s.MultipartHeaderSignatureInfos = v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBody) SetRequestId(v string) *GetMultipartFileUploadInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBody) SetVendorRequestId(v string) *GetMultipartFileUploadInfosResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBody) SetVendorType(v string) *GetMultipartFileUploadInfosResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos struct {
	HeaderSignatureInfo *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo `json:"HeaderSignatureInfo,omitempty" xml:"HeaderSignatureInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PartNumber *int32 `json:"PartNumber,omitempty" xml:"PartNumber,omitempty"`
}

func (s GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) GetHeaderSignatureInfo() *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	return s.HeaderSignatureInfo
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) GetPartNumber() *int32 {
	return s.PartNumber
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) SetHeaderSignatureInfo(v *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos {
	s.HeaderSignatureInfo = v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) SetPartNumber(v int32) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos {
	s.PartNumber = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfos) Validate() error {
	return dara.Validate(s)
}

type GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo struct {
	// example:
	//
	// 900
	ExpirationSeconds    *int32             `json:"ExpirationSeconds,omitempty" xml:"ExpirationSeconds,omitempty"`
	Headers              map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	InternalResourceUrls []*string          `json:"InternalResourceUrls,omitempty" xml:"InternalResourceUrls,omitempty" type:"Repeated"`
	// example:
	//
	// ZHANGJIAKOU
	Region       *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceUrls []*string `json:"ResourceUrls,omitempty" xml:"ResourceUrls,omitempty" type:"Repeated"`
}

func (s GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) GetExpirationSeconds() *int32 {
	return s.ExpirationSeconds
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) GetInternalResourceUrls() []*string {
	return s.InternalResourceUrls
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) GetRegion() *string {
	return s.Region
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) GetResourceUrls() []*string {
	return s.ResourceUrls
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetExpirationSeconds(v int32) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetHeaders(v map[string]*string) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.Headers = v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetInternalResourceUrls(v []*string) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.InternalResourceUrls = v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetRegion(v string) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.Region = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) SetResourceUrls(v []*string) *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo {
	s.ResourceUrls = v
	return s
}

func (s *GetMultipartFileUploadInfosResponseBodyMultipartHeaderSignatureInfosHeaderSignatureInfo) Validate() error {
	return dara.Validate(s)
}
