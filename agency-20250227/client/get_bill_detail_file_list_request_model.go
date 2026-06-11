// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillDetailFileListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillMonth(v string) *GetBillDetailFileListRequest
	GetBillMonth() *string
	SetOssAccessKeyId(v string) *GetBillDetailFileListRequest
	GetOssAccessKeyId() *string
	SetOssAccessKeySecret(v string) *GetBillDetailFileListRequest
	GetOssAccessKeySecret() *string
	SetOssBucketName(v string) *GetBillDetailFileListRequest
	GetOssBucketName() *string
	SetOssEndpoint(v string) *GetBillDetailFileListRequest
	GetOssEndpoint() *string
	SetOssRegion(v string) *GetBillDetailFileListRequest
	GetOssRegion() *string
	SetOssSecurityToken(v string) *GetBillDetailFileListRequest
	GetOssSecurityToken() *string
}

type GetBillDetailFileListRequest struct {
	// Month
	//
	// This parameter is required.
	//
	// example:
	//
	// 202502
	BillMonth *string `json:"BillMonth,omitempty" xml:"BillMonth,omitempty"`
	// The AccessKeyID used to upload files to OSS.
	//
	// example:
	//
	// yourAccessKeyID
	OssAccessKeyId *string `json:"OssAccessKeyId,omitempty" xml:"OssAccessKeyId,omitempty"`
	// The AccessKeySecret used to upload files to OSS.
	//
	// example:
	//
	// yourAccessKeySecret
	OssAccessKeySecret *string `json:"OssAccessKeySecret,omitempty" xml:"OssAccessKeySecret,omitempty"`
	// OSS bucket.
	//
	// example:
	//
	// yourBucketName
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The Region of the edge zone where the OSS bucket corresponding to the file sharing is located.
	//
	// example:
	//
	// http://oss-cn-beijing.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The Region to which the current OSS bucket belongs.
	//
	// example:
	//
	// cn-beijing
	OssRegion *string `json:"OssRegion,omitempty" xml:"OssRegion,omitempty"`
	// STS token used to upload files to OSS
	//
	// example:
	//
	// yourSecurityToken
	OssSecurityToken *string `json:"OssSecurityToken,omitempty" xml:"OssSecurityToken,omitempty"`
}

func (s GetBillDetailFileListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBillDetailFileListRequest) GoString() string {
	return s.String()
}

func (s *GetBillDetailFileListRequest) GetBillMonth() *string {
	return s.BillMonth
}

func (s *GetBillDetailFileListRequest) GetOssAccessKeyId() *string {
	return s.OssAccessKeyId
}

func (s *GetBillDetailFileListRequest) GetOssAccessKeySecret() *string {
	return s.OssAccessKeySecret
}

func (s *GetBillDetailFileListRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetBillDetailFileListRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *GetBillDetailFileListRequest) GetOssRegion() *string {
	return s.OssRegion
}

func (s *GetBillDetailFileListRequest) GetOssSecurityToken() *string {
	return s.OssSecurityToken
}

func (s *GetBillDetailFileListRequest) SetBillMonth(v string) *GetBillDetailFileListRequest {
	s.BillMonth = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssAccessKeyId(v string) *GetBillDetailFileListRequest {
	s.OssAccessKeyId = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssAccessKeySecret(v string) *GetBillDetailFileListRequest {
	s.OssAccessKeySecret = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssBucketName(v string) *GetBillDetailFileListRequest {
	s.OssBucketName = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssEndpoint(v string) *GetBillDetailFileListRequest {
	s.OssEndpoint = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssRegion(v string) *GetBillDetailFileListRequest {
	s.OssRegion = &v
	return s
}

func (s *GetBillDetailFileListRequest) SetOssSecurityToken(v string) *GetBillDetailFileListRequest {
	s.OssSecurityToken = &v
	return s
}

func (s *GetBillDetailFileListRequest) Validate() error {
	return dara.Validate(s)
}
