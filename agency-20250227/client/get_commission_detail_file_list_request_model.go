// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommissionDetailFileListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillMonth(v string) *GetCommissionDetailFileListRequest
	GetBillMonth() *string
	SetOssAccessKeyId(v string) *GetCommissionDetailFileListRequest
	GetOssAccessKeyId() *string
	SetOssAccessKeySecret(v string) *GetCommissionDetailFileListRequest
	GetOssAccessKeySecret() *string
	SetOssBucketName(v string) *GetCommissionDetailFileListRequest
	GetOssBucketName() *string
	SetOssEndpoint(v string) *GetCommissionDetailFileListRequest
	GetOssEndpoint() *string
	SetOssRegion(v string) *GetCommissionDetailFileListRequest
	GetOssRegion() *string
	SetOssSecurityToken(v string) *GetCommissionDetailFileListRequest
	GetOssSecurityToken() *string
}

type GetCommissionDetailFileListRequest struct {
	// The billing month.
	//
	// example:
	//
	// 202501
	BillMonth *string `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
	// The AccessKey ID used to upload files to OSS.
	//
	// example:
	//
	// yourAccessKeyID
	OssAccessKeyId *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	// The AccessKey secret used to upload files to OSS.
	//
	// example:
	//
	// yourAccessKeySecret
	OssAccessKeySecret *string `json:"OssAccessKeySecret,omitempty" xml:"OssAccessKeySecret,omitempty"`
	// The OSS bucket name.
	//
	// example:
	//
	// yourBucketName
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The endpoint of the region where the Object Storage Service (OSS) bucket corresponding to the file sharing resides.
	//
	// example:
	//
	// http://oss-cn-beijing.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The region where the OSS bucket resides.
	//
	// example:
	//
	// cn-beijing
	OssRegion *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	// The STS token used to upload files to OSS.
	//
	// example:
	//
	// yourSecurityToken
	OssSecurityToken *string `json:"OssSecurityToken,omitempty" xml:"OssSecurityToken,omitempty"`
}

func (s GetCommissionDetailFileListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionDetailFileListRequest) GoString() string {
	return s.String()
}

func (s *GetCommissionDetailFileListRequest) GetBillMonth() *string {
	return s.BillMonth
}

func (s *GetCommissionDetailFileListRequest) GetOssAccessKeyId() *string {
	return s.OssAccessKeyId
}

func (s *GetCommissionDetailFileListRequest) GetOssAccessKeySecret() *string {
	return s.OssAccessKeySecret
}

func (s *GetCommissionDetailFileListRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetCommissionDetailFileListRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *GetCommissionDetailFileListRequest) GetOssRegion() *string {
	return s.OssRegion
}

func (s *GetCommissionDetailFileListRequest) GetOssSecurityToken() *string {
	return s.OssSecurityToken
}

func (s *GetCommissionDetailFileListRequest) SetBillMonth(v string) *GetCommissionDetailFileListRequest {
	s.BillMonth = &v
	return s
}

func (s *GetCommissionDetailFileListRequest) SetOssAccessKeyId(v string) *GetCommissionDetailFileListRequest {
	s.OssAccessKeyId = &v
	return s
}

func (s *GetCommissionDetailFileListRequest) SetOssAccessKeySecret(v string) *GetCommissionDetailFileListRequest {
	s.OssAccessKeySecret = &v
	return s
}

func (s *GetCommissionDetailFileListRequest) SetOssBucketName(v string) *GetCommissionDetailFileListRequest {
	s.OssBucketName = &v
	return s
}

func (s *GetCommissionDetailFileListRequest) SetOssEndpoint(v string) *GetCommissionDetailFileListRequest {
	s.OssEndpoint = &v
	return s
}

func (s *GetCommissionDetailFileListRequest) SetOssRegion(v string) *GetCommissionDetailFileListRequest {
	s.OssRegion = &v
	return s
}

func (s *GetCommissionDetailFileListRequest) SetOssSecurityToken(v string) *GetCommissionDetailFileListRequest {
	s.OssSecurityToken = &v
	return s
}

func (s *GetCommissionDetailFileListRequest) Validate() error {
	return dara.Validate(s)
}
