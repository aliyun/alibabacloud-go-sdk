// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCACertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCACertificateId(v string) *UploadCACertificateResponseBody
	GetCACertificateId() *string
	SetCACertificateName(v string) *UploadCACertificateResponseBody
	GetCACertificateName() *string
	SetCommonName(v string) *UploadCACertificateResponseBody
	GetCommonName() *string
	SetCreateTime(v string) *UploadCACertificateResponseBody
	GetCreateTime() *string
	SetCreateTimeStamp(v int64) *UploadCACertificateResponseBody
	GetCreateTimeStamp() *int64
	SetExpireTime(v string) *UploadCACertificateResponseBody
	GetExpireTime() *string
	SetExpireTimeStamp(v int64) *UploadCACertificateResponseBody
	GetExpireTimeStamp() *int64
	SetFingerprint(v string) *UploadCACertificateResponseBody
	GetFingerprint() *string
	SetRequestId(v string) *UploadCACertificateResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *UploadCACertificateResponseBody
	GetResourceGroupId() *string
}

type UploadCACertificateResponseBody struct {
	// The ID of the CA certificate.
	//
	// example:
	//
	// 139a0******-cn-east-hangzhou-01
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	// The CA certificate name.
	//
	// example:
	//
	// mycacert01
	CACertificateName *string `json:"CACertificateName,omitempty" xml:"CACertificateName,omitempty"`
	// The domain name on the CA certificate.
	//
	// example:
	//
	// .example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The time when the CA certificate was created.
	//
	// example:
	//
	// 2017-08-31T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp when the CA certificate was created.
	//
	// example:
	//
	// 1504147745000
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The time when the CA certificate expires.
	//
	// example:
	//
	// 2024-11-21T06:04:25Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The timestamp when the server certificate expires.
	//
	// example:
	//
	// 1732169065000
	ExpireTimeStamp *int64 `json:"ExpireTimeStamp,omitempty" xml:"ExpireTimeStamp,omitempty"`
	// The fingerprint of the server certificate.
	//
	// example:
	//
	// 02:DF:AB:ED
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rt******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UploadCACertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCACertificateResponseBody) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *UploadCACertificateResponseBody) GetCACertificateName() *string {
	return s.CACertificateName
}

func (s *UploadCACertificateResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *UploadCACertificateResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UploadCACertificateResponseBody) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *UploadCACertificateResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *UploadCACertificateResponseBody) GetExpireTimeStamp() *int64 {
	return s.ExpireTimeStamp
}

func (s *UploadCACertificateResponseBody) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *UploadCACertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadCACertificateResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UploadCACertificateResponseBody) SetCACertificateId(v string) *UploadCACertificateResponseBody {
	s.CACertificateId = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetCACertificateName(v string) *UploadCACertificateResponseBody {
	s.CACertificateName = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetCommonName(v string) *UploadCACertificateResponseBody {
	s.CommonName = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetCreateTime(v string) *UploadCACertificateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetCreateTimeStamp(v int64) *UploadCACertificateResponseBody {
	s.CreateTimeStamp = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetExpireTime(v string) *UploadCACertificateResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetExpireTimeStamp(v int64) *UploadCACertificateResponseBody {
	s.ExpireTimeStamp = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetFingerprint(v string) *UploadCACertificateResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetRequestId(v string) *UploadCACertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadCACertificateResponseBody) SetResourceGroupId(v string) *UploadCACertificateResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *UploadCACertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
