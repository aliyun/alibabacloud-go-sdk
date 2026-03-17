// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSmartAccessGatewayVersionsResponseBody
	GetRequestId() *string
	SetSmartAGVersions(v *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions) *DescribeSmartAccessGatewayVersionsResponseBody
	GetSmartAGVersions() *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions
}

type DescribeSmartAccessGatewayVersionsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 765AB188-69BF-47C6-BEDD-B9FC72BFBB0
	RequestId       *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmartAGVersions *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions `json:"SmartAGVersions,omitempty" xml:"SmartAGVersions,omitempty" type:"Struct"`
}

func (s DescribeSmartAccessGatewayVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSmartAccessGatewayVersionsResponseBody) GetSmartAGVersions() *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions {
	return s.SmartAGVersions
}

func (s *DescribeSmartAccessGatewayVersionsResponseBody) SetRequestId(v string) *DescribeSmartAccessGatewayVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponseBody) SetSmartAGVersions(v *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions) *DescribeSmartAccessGatewayVersionsResponseBody {
	s.SmartAGVersions = v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponseBody) Validate() error {
	if s.SmartAGVersions != nil {
		if err := s.SmartAGVersions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions struct {
	SmartAGVersion []*DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion `json:"SmartAGVersion,omitempty" xml:"SmartAGVersion,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions) GetSmartAGVersion() []*DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion {
	return s.SmartAGVersion
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions) SetSmartAGVersion(v []*DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions {
	s.SmartAGVersion = v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersions) Validate() error {
	if s.SmartAGVersion != nil {
		for _, item := range s.SmartAGVersion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion struct {
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) GetType() *string {
	return s.Type
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) GetVersionName() *string {
	return s.VersionName
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) SetCreateTime(v int64) *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion {
	s.CreateTime = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) SetType(v string) *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion {
	s.Type = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) SetVersionCode(v string) *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion {
	s.VersionCode = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) SetVersionName(v string) *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion {
	s.VersionName = &v
	return s
}

func (s *DescribeSmartAccessGatewayVersionsResponseBodySmartAGVersionsSmartAGVersion) Validate() error {
	return dara.Validate(s)
}
