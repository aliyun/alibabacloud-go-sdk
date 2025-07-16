// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaderSignatureInfo(v *GetFileUploadInfoResponseBodyHeaderSignatureInfo) *GetFileUploadInfoResponseBody
	GetHeaderSignatureInfo() *GetFileUploadInfoResponseBodyHeaderSignatureInfo
	SetProtocol(v string) *GetFileUploadInfoResponseBody
	GetProtocol() *string
	SetRequestId(v string) *GetFileUploadInfoResponseBody
	GetRequestId() *string
	SetStorageDriver(v string) *GetFileUploadInfoResponseBody
	GetStorageDriver() *string
	SetUploadKey(v string) *GetFileUploadInfoResponseBody
	GetUploadKey() *string
	SetVendorRequestId(v string) *GetFileUploadInfoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetFileUploadInfoResponseBody
	GetVendorType() *string
}

type GetFileUploadInfoResponseBody struct {
	HeaderSignatureInfo *GetFileUploadInfoResponseBodyHeaderSignatureInfo `json:"headerSignatureInfo,omitempty" xml:"headerSignatureInfo,omitempty" type:"Struct"`
	// example:
	//
	// HEADER_SIGNATURE
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// example:
	//
	// upload_key
	UploadKey *string `json:"uploadKey,omitempty" xml:"uploadKey,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetFileUploadInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponseBody) GetHeaderSignatureInfo() *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	return s.HeaderSignatureInfo
}

func (s *GetFileUploadInfoResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *GetFileUploadInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileUploadInfoResponseBody) GetStorageDriver() *string {
	return s.StorageDriver
}

func (s *GetFileUploadInfoResponseBody) GetUploadKey() *string {
	return s.UploadKey
}

func (s *GetFileUploadInfoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetFileUploadInfoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetFileUploadInfoResponseBody) SetHeaderSignatureInfo(v *GetFileUploadInfoResponseBodyHeaderSignatureInfo) *GetFileUploadInfoResponseBody {
	s.HeaderSignatureInfo = v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetProtocol(v string) *GetFileUploadInfoResponseBody {
	s.Protocol = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetRequestId(v string) *GetFileUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetStorageDriver(v string) *GetFileUploadInfoResponseBody {
	s.StorageDriver = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetUploadKey(v string) *GetFileUploadInfoResponseBody {
	s.UploadKey = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetVendorRequestId(v string) *GetFileUploadInfoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) SetVendorType(v string) *GetFileUploadInfoResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetFileUploadInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileUploadInfoResponseBodyHeaderSignatureInfo struct {
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

func (s GetFileUploadInfoResponseBodyHeaderSignatureInfo) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoResponseBodyHeaderSignatureInfo) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) GetExpirationSeconds() *int32 {
	return s.ExpirationSeconds
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) GetInternalResourceUrls() []*string {
	return s.InternalResourceUrls
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) GetRegion() *string {
	return s.Region
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) GetResourceUrls() []*string {
	return s.ResourceUrls
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetExpirationSeconds(v int32) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetHeaders(v map[string]*string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.Headers = v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetInternalResourceUrls(v []*string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.InternalResourceUrls = v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetRegion(v string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.Region = &v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) SetResourceUrls(v []*string) *GetFileUploadInfoResponseBodyHeaderSignatureInfo {
	s.ResourceUrls = v
	return s
}

func (s *GetFileUploadInfoResponseBodyHeaderSignatureInfo) Validate() error {
	return dara.Validate(s)
}
