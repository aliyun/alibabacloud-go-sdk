// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServerCertificatesResponseBody
	GetRequestId() *string
	SetServerCertificates(v *DescribeServerCertificatesResponseBodyServerCertificates) *DescribeServerCertificatesResponseBody
	GetServerCertificates() *DescribeServerCertificatesResponseBodyServerCertificates
}

type DescribeServerCertificatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId          *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerCertificates *DescribeServerCertificatesResponseBodyServerCertificates `json:"ServerCertificates,omitempty" xml:"ServerCertificates,omitempty" type:"Struct"`
}

func (s DescribeServerCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServerCertificatesResponseBody) GetServerCertificates() *DescribeServerCertificatesResponseBodyServerCertificates {
	return s.ServerCertificates
}

func (s *DescribeServerCertificatesResponseBody) SetRequestId(v string) *DescribeServerCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBody) SetServerCertificates(v *DescribeServerCertificatesResponseBodyServerCertificates) *DescribeServerCertificatesResponseBody {
	s.ServerCertificates = v
	return s
}

func (s *DescribeServerCertificatesResponseBody) Validate() error {
	if s.ServerCertificates != nil {
		if err := s.ServerCertificates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServerCertificatesResponseBodyServerCertificates struct {
	ServerCertificate []*DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty" type:"Repeated"`
}

func (s DescribeServerCertificatesResponseBodyServerCertificates) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerCertificatesResponseBodyServerCertificates) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBodyServerCertificates) GetServerCertificate() []*DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	return s.ServerCertificate
}

func (s *DescribeServerCertificatesResponseBodyServerCertificates) SetServerCertificate(v []*DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) *DescribeServerCertificatesResponseBodyServerCertificates {
	s.ServerCertificate = v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificates) Validate() error {
	if s.ServerCertificate != nil {
		for _, item := range s.ServerCertificate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate struct {
	AliCloudCertificateId   *string                                                                                           `json:"AliCloudCertificateId,omitempty" xml:"AliCloudCertificateId,omitempty"`
	AliCloudCertificateName *string                                                                                           `json:"AliCloudCertificateName,omitempty" xml:"AliCloudCertificateName,omitempty"`
	CommonName              *string                                                                                           `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CreateTime              *string                                                                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeStamp         *int64                                                                                            `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	ExpireTime              *string                                                                                           `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ExpireTimeStamp         *int64                                                                                            `json:"ExpireTimeStamp,omitempty" xml:"ExpireTimeStamp,omitempty"`
	Fingerprint             *string                                                                                           `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	IsAliCloudCertificate   *int32                                                                                            `json:"IsAliCloudCertificate,omitempty" xml:"IsAliCloudCertificate,omitempty"`
	RegionId                *string                                                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId         *string                                                                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerCertificateId     *string                                                                                           `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
	ServerCertificateName   *string                                                                                           `json:"ServerCertificateName,omitempty" xml:"ServerCertificateName,omitempty"`
	SubjectAlternativeNames *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Struct"`
	Tags                    *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetAliCloudCertificateId() *string {
	return s.AliCloudCertificateId
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetAliCloudCertificateName() *string {
	return s.AliCloudCertificateName
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetExpireTimeStamp() *int64 {
	return s.ExpireTimeStamp
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetIsAliCloudCertificate() *int32 {
	return s.IsAliCloudCertificate
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetServerCertificateName() *string {
	return s.ServerCertificateName
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetSubjectAlternativeNames() *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames {
	return s.SubjectAlternativeNames
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) GetTags() *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags {
	return s.Tags
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetAliCloudCertificateId(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.AliCloudCertificateId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetAliCloudCertificateName(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.AliCloudCertificateName = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetCommonName(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.CommonName = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetCreateTime(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.CreateTime = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetCreateTimeStamp(v int64) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetExpireTime(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ExpireTime = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetExpireTimeStamp(v int64) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ExpireTimeStamp = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetFingerprint(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.Fingerprint = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetIsAliCloudCertificate(v int32) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.IsAliCloudCertificate = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetRegionId(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.RegionId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetResourceGroupId(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetServerCertificateId(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetServerCertificateName(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.ServerCertificateName = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetSubjectAlternativeNames(v *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.SubjectAlternativeNames = v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) SetTags(v *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate {
	s.Tags = v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificate) Validate() error {
	if s.SubjectAlternativeNames != nil {
		if err := s.SubjectAlternativeNames.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames struct {
	SubjectAlternativeName []*string `json:"SubjectAlternativeName,omitempty" xml:"SubjectAlternativeName,omitempty" type:"Repeated"`
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) GetSubjectAlternativeName() []*string {
	return s.SubjectAlternativeName
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) SetSubjectAlternativeName(v []*string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames {
	s.SubjectAlternativeName = v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateSubjectAlternativeNames) Validate() error {
	return dara.Validate(s)
}

type DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags struct {
	Tag []*DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags) GetTag() []*DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag {
	return s.Tag
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags) SetTag(v []*DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags {
	s.Tag = v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag) SetTagKey(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag) SetTagValue(v string) *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeServerCertificatesResponseBodyServerCertificatesServerCertificateTagsTag) Validate() error {
	return dara.Validate(s)
}
