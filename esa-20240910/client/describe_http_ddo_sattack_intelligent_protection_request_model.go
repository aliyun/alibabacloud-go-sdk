// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSAttackIntelligentProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DescribeHttpDDoSAttackIntelligentProtectionRequest
	GetSiteId() *int64
}

type DescribeHttpDDoSAttackIntelligentProtectionRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSAttackIntelligentProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackIntelligentProtectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionRequest) SetSiteId(v int64) *DescribeHttpDDoSAttackIntelligentProtectionRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeHttpDDoSAttackIntelligentProtectionRequest) Validate() error {
	return dara.Validate(s)
}
