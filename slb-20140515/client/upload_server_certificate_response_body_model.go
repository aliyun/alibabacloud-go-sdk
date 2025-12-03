// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadServerCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliCloudCertificateId(v string) *UploadServerCertificateResponseBody
	GetAliCloudCertificateId() *string
	SetAliCloudCertificateName(v string) *UploadServerCertificateResponseBody
	GetAliCloudCertificateName() *string
	SetCommonName(v string) *UploadServerCertificateResponseBody
	GetCommonName() *string
	SetCreateTime(v string) *UploadServerCertificateResponseBody
	GetCreateTime() *string
	SetCreateTimeStamp(v int64) *UploadServerCertificateResponseBody
	GetCreateTimeStamp() *int64
	SetExpireTime(v string) *UploadServerCertificateResponseBody
	GetExpireTime() *string
	SetExpireTimeStamp(v int64) *UploadServerCertificateResponseBody
	GetExpireTimeStamp() *int64
	SetFingerprint(v string) *UploadServerCertificateResponseBody
	GetFingerprint() *string
	SetIsAliCloudCertificate(v int32) *UploadServerCertificateResponseBody
	GetIsAliCloudCertificate() *int32
	SetRegionId(v string) *UploadServerCertificateResponseBody
	GetRegionId() *string
	SetRequestId(v string) *UploadServerCertificateResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *UploadServerCertificateResponseBody
	GetResourceGroupId() *string
	SetServerCertificateId(v string) *UploadServerCertificateResponseBody
	GetServerCertificateId() *string
	SetServerCertificateName(v string) *UploadServerCertificateResponseBody
	GetServerCertificateName() *string
	SetSubjectAlternativeNames(v *UploadServerCertificateResponseBodySubjectAlternativeNames) *UploadServerCertificateResponseBody
	GetSubjectAlternativeNames() *UploadServerCertificateResponseBodySubjectAlternativeNames
}

type UploadServerCertificateResponseBody struct {
	// The AliCloud certificate ID.
	//
	// example:
	//
	// 775****
	AliCloudCertificateId *string `json:"AliCloudCertificateId,omitempty" xml:"AliCloudCertificateId,omitempty"`
	// The AliCloud certificate name.
	//
	// example:
	//
	// cloudcertificate****
	AliCloudCertificateName *string `json:"AliCloudCertificateName,omitempty" xml:"AliCloudCertificateName,omitempty"`
	// The domain name of the CA certificate.
	//
	// example:
	//
	// example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The time when the CA certificate is uploaded.
	//
	// example:
	//
	// 2022-02-21T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp generated when the CA certificate is uploaded.
	//
	// example:
	//
	// 1504147745000
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The time when the CA certificate expires.
	//
	// example:
	//
	// 2022-10-18T23:59:59Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The timestamp generated when the CA certificate expires.
	//
	// example:
	//
	// 1504147745000
	ExpireTimeStamp *int64 `json:"ExpireTimeStamp,omitempty" xml:"ExpireTimeStamp,omitempty"`
	// The fingerprint of the CA certificate.
	//
	// example:
	//
	// 8f:7d:cb:e5:f8:c8:33:9c:17:65:c1:92:30:9e:45:55:9c:3a:85:60
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// Indicates whether the certificate is provided by Alibaba Cloud Certificate Management Service. Valid values:
	//
	// - **0**: The certificate is not provided by Alibaba Cloud Certificate Management Service.
	//
	// - **1**: The certificate is provided by Alibaba Cloud Certificate Management Service.
	//
	// example:
	//
	// 0
	IsAliCloudCertificate *int32 `json:"IsAliCloudCertificate,omitempty" xml:"IsAliCloudCertificate,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rt****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the server certificate.
	//
	// example:
	//
	// 1321932713******_17f1b4b696b_1114720822_-1671******
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	// The name of the server certificate.
	//
	// The name must be 1 to 80 characters in length. It must start with an English letter. It can contain letters, numbers, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// mycert01
	ServerCertificateName *string `json:"ServerCertificateName,omitempty" xml:"ServerCertificateName,omitempty"`
	// The subject alternative names.
	//
	// A domain name list is supported. A maximum of 10 domain names are supported.
	SubjectAlternativeNames *UploadServerCertificateResponseBodySubjectAlternativeNames `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Struct"`
}

func (s UploadServerCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadServerCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadServerCertificateResponseBody) GetAliCloudCertificateId() *string {
	return s.AliCloudCertificateId
}

func (s *UploadServerCertificateResponseBody) GetAliCloudCertificateName() *string {
	return s.AliCloudCertificateName
}

func (s *UploadServerCertificateResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *UploadServerCertificateResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UploadServerCertificateResponseBody) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *UploadServerCertificateResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *UploadServerCertificateResponseBody) GetExpireTimeStamp() *int64 {
	return s.ExpireTimeStamp
}

func (s *UploadServerCertificateResponseBody) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *UploadServerCertificateResponseBody) GetIsAliCloudCertificate() *int32 {
	return s.IsAliCloudCertificate
}

func (s *UploadServerCertificateResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadServerCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadServerCertificateResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UploadServerCertificateResponseBody) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *UploadServerCertificateResponseBody) GetServerCertificateName() *string {
	return s.ServerCertificateName
}

func (s *UploadServerCertificateResponseBody) GetSubjectAlternativeNames() *UploadServerCertificateResponseBodySubjectAlternativeNames {
	return s.SubjectAlternativeNames
}

func (s *UploadServerCertificateResponseBody) SetAliCloudCertificateId(v string) *UploadServerCertificateResponseBody {
	s.AliCloudCertificateId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetAliCloudCertificateName(v string) *UploadServerCertificateResponseBody {
	s.AliCloudCertificateName = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetCommonName(v string) *UploadServerCertificateResponseBody {
	s.CommonName = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetCreateTime(v string) *UploadServerCertificateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetCreateTimeStamp(v int64) *UploadServerCertificateResponseBody {
	s.CreateTimeStamp = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetExpireTime(v string) *UploadServerCertificateResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetExpireTimeStamp(v int64) *UploadServerCertificateResponseBody {
	s.ExpireTimeStamp = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetFingerprint(v string) *UploadServerCertificateResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetIsAliCloudCertificate(v int32) *UploadServerCertificateResponseBody {
	s.IsAliCloudCertificate = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetRegionId(v string) *UploadServerCertificateResponseBody {
	s.RegionId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetRequestId(v string) *UploadServerCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetResourceGroupId(v string) *UploadServerCertificateResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetServerCertificateId(v string) *UploadServerCertificateResponseBody {
	s.ServerCertificateId = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetServerCertificateName(v string) *UploadServerCertificateResponseBody {
	s.ServerCertificateName = &v
	return s
}

func (s *UploadServerCertificateResponseBody) SetSubjectAlternativeNames(v *UploadServerCertificateResponseBodySubjectAlternativeNames) *UploadServerCertificateResponseBody {
	s.SubjectAlternativeNames = v
	return s
}

func (s *UploadServerCertificateResponseBody) Validate() error {
	if s.SubjectAlternativeNames != nil {
		if err := s.SubjectAlternativeNames.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadServerCertificateResponseBodySubjectAlternativeNames struct {
	SubjectAlternativeName []*string `json:"SubjectAlternativeName,omitempty" xml:"SubjectAlternativeName,omitempty" type:"Repeated"`
}

func (s UploadServerCertificateResponseBodySubjectAlternativeNames) String() string {
	return dara.Prettify(s)
}

func (s UploadServerCertificateResponseBodySubjectAlternativeNames) GoString() string {
	return s.String()
}

func (s *UploadServerCertificateResponseBodySubjectAlternativeNames) GetSubjectAlternativeName() []*string {
	return s.SubjectAlternativeName
}

func (s *UploadServerCertificateResponseBodySubjectAlternativeNames) SetSubjectAlternativeName(v []*string) *UploadServerCertificateResponseBodySubjectAlternativeNames {
	s.SubjectAlternativeName = v
	return s
}

func (s *UploadServerCertificateResponseBodySubjectAlternativeNames) Validate() error {
	return dara.Validate(s)
}
