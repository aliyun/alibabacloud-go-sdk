// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDownloadInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaderSignatureInfo(v *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) *GetFileDownloadInfoResponseBody
	GetHeaderSignatureInfo() *GetFileDownloadInfoResponseBodyHeaderSignatureInfo
	SetProtocol(v string) *GetFileDownloadInfoResponseBody
	GetProtocol() *string
	SetRequestId(v string) *GetFileDownloadInfoResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetFileDownloadInfoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetFileDownloadInfoResponseBody
	GetVendorType() *string
}

type GetFileDownloadInfoResponseBody struct {
	HeaderSignatureInfo *GetFileDownloadInfoResponseBodyHeaderSignatureInfo `json:"headerSignatureInfo,omitempty" xml:"headerSignatureInfo,omitempty" type:"Struct"`
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
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetFileDownloadInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoResponseBody) GetHeaderSignatureInfo() *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	return s.HeaderSignatureInfo
}

func (s *GetFileDownloadInfoResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *GetFileDownloadInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileDownloadInfoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetFileDownloadInfoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetFileDownloadInfoResponseBody) SetHeaderSignatureInfo(v *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) *GetFileDownloadInfoResponseBody {
	s.HeaderSignatureInfo = v
	return s
}

func (s *GetFileDownloadInfoResponseBody) SetProtocol(v string) *GetFileDownloadInfoResponseBody {
	s.Protocol = &v
	return s
}

func (s *GetFileDownloadInfoResponseBody) SetRequestId(v string) *GetFileDownloadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileDownloadInfoResponseBody) SetVendorRequestId(v string) *GetFileDownloadInfoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetFileDownloadInfoResponseBody) SetVendorType(v string) *GetFileDownloadInfoResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetFileDownloadInfoResponseBody) Validate() error {
	if s.HeaderSignatureInfo != nil {
		if err := s.HeaderSignatureInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileDownloadInfoResponseBodyHeaderSignatureInfo struct {
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

func (s GetFileDownloadInfoResponseBodyHeaderSignatureInfo) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoResponseBodyHeaderSignatureInfo) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) GetExpirationSeconds() *int32 {
	return s.ExpirationSeconds
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) GetInternalResourceUrls() []*string {
	return s.InternalResourceUrls
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) GetRegion() *string {
	return s.Region
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) GetResourceUrls() []*string {
	return s.ResourceUrls
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetExpirationSeconds(v int32) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.ExpirationSeconds = &v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetHeaders(v map[string]*string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.Headers = v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetInternalResourceUrls(v []*string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.InternalResourceUrls = v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetRegion(v string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.Region = &v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) SetResourceUrls(v []*string) *GetFileDownloadInfoResponseBodyHeaderSignatureInfo {
	s.ResourceUrls = v
	return s
}

func (s *GetFileDownloadInfoResponseBodyHeaderSignatureInfo) Validate() error {
	return dara.Validate(s)
}
