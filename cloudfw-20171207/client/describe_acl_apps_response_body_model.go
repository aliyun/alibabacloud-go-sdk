// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclApps(v []*DescribeAclAppsResponseBodyAclApps) *DescribeAclAppsResponseBody
	GetAclApps() []*DescribeAclAppsResponseBodyAclApps
	SetRequestId(v string) *DescribeAclAppsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAclAppsResponseBody
	GetTotalCount() *int64
}

type DescribeAclAppsResponseBody struct {
	AclApps []*DescribeAclAppsResponseBodyAclApps `json:"AclApps,omitempty" xml:"AclApps,omitempty" type:"Repeated"`
	// example:
	//
	// 9063AB86-6FFA-5B2D-A16D-697C966D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAclAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclAppsResponseBody) GetAclApps() []*DescribeAclAppsResponseBodyAclApps {
	return s.AclApps
}

func (s *DescribeAclAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclAppsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAclAppsResponseBody) SetAclApps(v []*DescribeAclAppsResponseBodyAclApps) *DescribeAclAppsResponseBody {
	s.AclApps = v
	return s
}

func (s *DescribeAclAppsResponseBody) SetRequestId(v string) *DescribeAclAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclAppsResponseBody) SetTotalCount(v int64) *DescribeAclAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAclAppsResponseBody) Validate() error {
	if s.AclApps != nil {
		for _, item := range s.AclApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAclAppsResponseBodyAclApps struct {
	// example:
	//
	// 1.0
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	Popular   *int32    `json:"Popular,omitempty" xml:"Popular,omitempty"`
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// 1
	SupportFqdn *int32 `json:"SupportFqdn,omitempty" xml:"SupportFqdn,omitempty"`
}

func (s DescribeAclAppsResponseBodyAclApps) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclAppsResponseBodyAclApps) GoString() string {
	return s.String()
}

func (s *DescribeAclAppsResponseBodyAclApps) GetAppId() *int32 {
	return s.AppId
}

func (s *DescribeAclAppsResponseBodyAclApps) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAclAppsResponseBodyAclApps) GetPopular() *int32 {
	return s.Popular
}

func (s *DescribeAclAppsResponseBodyAclApps) GetProtocols() []*string {
	return s.Protocols
}

func (s *DescribeAclAppsResponseBodyAclApps) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeAclAppsResponseBodyAclApps) GetSupportFqdn() *int32 {
	return s.SupportFqdn
}

func (s *DescribeAclAppsResponseBodyAclApps) SetAppId(v int32) *DescribeAclAppsResponseBodyAclApps {
	s.AppId = &v
	return s
}

func (s *DescribeAclAppsResponseBodyAclApps) SetAppName(v string) *DescribeAclAppsResponseBodyAclApps {
	s.AppName = &v
	return s
}

func (s *DescribeAclAppsResponseBodyAclApps) SetPopular(v int32) *DescribeAclAppsResponseBodyAclApps {
	s.Popular = &v
	return s
}

func (s *DescribeAclAppsResponseBodyAclApps) SetProtocols(v []*string) *DescribeAclAppsResponseBodyAclApps {
	s.Protocols = v
	return s
}

func (s *DescribeAclAppsResponseBodyAclApps) SetRiskLevel(v int32) *DescribeAclAppsResponseBodyAclApps {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAclAppsResponseBodyAclApps) SetSupportFqdn(v int32) *DescribeAclAppsResponseBodyAclApps {
	s.SupportFqdn = &v
	return s
}

func (s *DescribeAclAppsResponseBodyAclApps) Validate() error {
	return dara.Validate(s)
}
