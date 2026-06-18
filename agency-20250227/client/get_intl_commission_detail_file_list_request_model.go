// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntlCommissionDetailFileListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillMonth(v string) *GetIntlCommissionDetailFileListRequest
	GetBillMonth() *string
	SetOssAccessKeyId(v string) *GetIntlCommissionDetailFileListRequest
	GetOssAccessKeyId() *string
	SetOssAccessKeySecret(v string) *GetIntlCommissionDetailFileListRequest
	GetOssAccessKeySecret() *string
	SetOssBucketName(v string) *GetIntlCommissionDetailFileListRequest
	GetOssBucketName() *string
	SetOssEndpoint(v string) *GetIntlCommissionDetailFileListRequest
	GetOssEndpoint() *string
	SetOssRegion(v string) *GetIntlCommissionDetailFileListRequest
	GetOssRegion() *string
	SetOssSecurityToken(v string) *GetIntlCommissionDetailFileListRequest
	GetOssSecurityToken() *string
}

type GetIntlCommissionDetailFileListRequest struct {
	// The billing month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202502
	BillMonth *string `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
	// The AccessKey ID used to upload files to OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAccessKeyID
	OssAccessKeyId *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	// The AccessKey secret used to upload files to OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAccessKeySecret
	OssAccessKeySecret *string `json:"OssAccessKeySecret,omitempty" xml:"OssAccessKeySecret,omitempty"`
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourBucketName
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The endpoint of the region where the Object Storage Service (OSS) bucket for the file sharing resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://oss-cn-beijing.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The region where the current OSS bucket resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	OssRegion *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	// The STS token used to upload files to OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourSecurityToken
	OssSecurityToken *string `json:"OssSecurityToken,omitempty" xml:"OssSecurityToken,omitempty"`
}

func (s GetIntlCommissionDetailFileListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIntlCommissionDetailFileListRequest) GoString() string {
	return s.String()
}

func (s *GetIntlCommissionDetailFileListRequest) GetBillMonth() *string {
	return s.BillMonth
}

func (s *GetIntlCommissionDetailFileListRequest) GetOssAccessKeyId() *string {
	return s.OssAccessKeyId
}

func (s *GetIntlCommissionDetailFileListRequest) GetOssAccessKeySecret() *string {
	return s.OssAccessKeySecret
}

func (s *GetIntlCommissionDetailFileListRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetIntlCommissionDetailFileListRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *GetIntlCommissionDetailFileListRequest) GetOssRegion() *string {
	return s.OssRegion
}

func (s *GetIntlCommissionDetailFileListRequest) GetOssSecurityToken() *string {
	return s.OssSecurityToken
}

func (s *GetIntlCommissionDetailFileListRequest) SetBillMonth(v string) *GetIntlCommissionDetailFileListRequest {
	s.BillMonth = &v
	return s
}

func (s *GetIntlCommissionDetailFileListRequest) SetOssAccessKeyId(v string) *GetIntlCommissionDetailFileListRequest {
	s.OssAccessKeyId = &v
	return s
}

func (s *GetIntlCommissionDetailFileListRequest) SetOssAccessKeySecret(v string) *GetIntlCommissionDetailFileListRequest {
	s.OssAccessKeySecret = &v
	return s
}

func (s *GetIntlCommissionDetailFileListRequest) SetOssBucketName(v string) *GetIntlCommissionDetailFileListRequest {
	s.OssBucketName = &v
	return s
}

func (s *GetIntlCommissionDetailFileListRequest) SetOssEndpoint(v string) *GetIntlCommissionDetailFileListRequest {
	s.OssEndpoint = &v
	return s
}

func (s *GetIntlCommissionDetailFileListRequest) SetOssRegion(v string) *GetIntlCommissionDetailFileListRequest {
	s.OssRegion = &v
	return s
}

func (s *GetIntlCommissionDetailFileListRequest) SetOssSecurityToken(v string) *GetIntlCommissionDetailFileListRequest {
	s.OssSecurityToken = &v
	return s
}

func (s *GetIntlCommissionDetailFileListRequest) Validate() error {
	return dara.Validate(s)
}
