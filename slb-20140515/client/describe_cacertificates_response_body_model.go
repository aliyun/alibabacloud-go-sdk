// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCACertificates(v *DescribeCACertificatesResponseBodyCACertificates) *DescribeCACertificatesResponseBody
	GetCACertificates() *DescribeCACertificatesResponseBodyCACertificates
	SetRequestId(v string) *DescribeCACertificatesResponseBody
	GetRequestId() *string
}

type DescribeCACertificatesResponseBody struct {
	// The information about the CA certificate.
	CACertificates *DescribeCACertificatesResponseBodyCACertificates `json:"CACertificates,omitempty" xml:"CACertificates,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCACertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponseBody) GetCACertificates() *DescribeCACertificatesResponseBodyCACertificates {
	return s.CACertificates
}

func (s *DescribeCACertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCACertificatesResponseBody) SetCACertificates(v *DescribeCACertificatesResponseBodyCACertificates) *DescribeCACertificatesResponseBody {
	s.CACertificates = v
	return s
}

func (s *DescribeCACertificatesResponseBody) SetRequestId(v string) *DescribeCACertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCACertificatesResponseBody) Validate() error {
	if s.CACertificates != nil {
		if err := s.CACertificates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCACertificatesResponseBodyCACertificates struct {
	CACertificate []*DescribeCACertificatesResponseBodyCACertificatesCACertificate `json:"CACertificate,omitempty" xml:"CACertificate,omitempty" type:"Repeated"`
}

func (s DescribeCACertificatesResponseBodyCACertificates) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificatesResponseBodyCACertificates) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponseBodyCACertificates) GetCACertificate() []*DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	return s.CACertificate
}

func (s *DescribeCACertificatesResponseBodyCACertificates) SetCACertificate(v []*DescribeCACertificatesResponseBodyCACertificatesCACertificate) *DescribeCACertificatesResponseBodyCACertificates {
	s.CACertificate = v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificates) Validate() error {
	if s.CACertificate != nil {
		for _, item := range s.CACertificate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCACertificatesResponseBodyCACertificatesCACertificate struct {
	// The CA certificate ID.
	//
	// example:
	//
	// 139a00604bd-cn-east-hangzho****
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	// The CA certificate name.
	//
	// example:
	//
	// test
	CACertificateName *string `json:"CACertificateName,omitempty" xml:"CACertificateName,omitempty"`
	// The domain name of the CA certificate.
	//
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The time when the CA certificate was created. The time is in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2021-08-31T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp when the CA certificate was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1504147745000
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The time when the CA certificate expires. The time is in the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2024-11-21T06:04:25Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The timestamp that indicates when the CA certificate expires. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1732169065000
	ExpireTimeStamp *int64 `json:"ExpireTimeStamp,omitempty" xml:"ExpireTimeStamp,omitempty"`
	// The fingerprint of the CA certificate.
	//
	// example:
	//
	// 79:43:fb:7d:a4:7f:44:32:61:16:57:17:e3:e8:b7:36:03:57:f6:89
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The region of the CA certificate.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag.
	Tags *DescribeCACertificatesResponseBodyCACertificatesCACertificateTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeCACertificatesResponseBodyCACertificatesCACertificate) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificatesResponseBodyCACertificatesCACertificate) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetCACertificateName() *string {
	return s.CACertificateName
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetExpireTimeStamp() *int64 {
	return s.ExpireTimeStamp
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) GetTags() *DescribeCACertificatesResponseBodyCACertificatesCACertificateTags {
	return s.Tags
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCACertificateId(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CACertificateId = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCACertificateName(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CACertificateName = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCommonName(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CommonName = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCreateTime(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CreateTime = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetCreateTimeStamp(v int64) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetExpireTime(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.ExpireTime = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetExpireTimeStamp(v int64) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.ExpireTimeStamp = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetFingerprint(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.Fingerprint = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetRegionId(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.RegionId = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetResourceGroupId(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) SetTags(v *DescribeCACertificatesResponseBodyCACertificatesCACertificateTags) *DescribeCACertificatesResponseBodyCACertificatesCACertificate {
	s.Tags = v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificate) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCACertificatesResponseBodyCACertificatesCACertificateTags struct {
	Tag []*DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCACertificatesResponseBodyCACertificatesCACertificateTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificatesResponseBodyCACertificatesCACertificateTags) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificateTags) GetTag() []*DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag {
	return s.Tag
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificateTags) SetTag(v []*DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag) *DescribeCACertificatesResponseBodyCACertificatesCACertificateTags {
	s.Tag = v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificateTags) Validate() error {
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

type DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag) SetTagKey(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag) SetTagValue(v string) *DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeCACertificatesResponseBodyCACertificatesCACertificateTagsTag) Validate() error {
	return dara.Validate(s)
}
