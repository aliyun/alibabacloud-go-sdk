// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSAttackProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *DescribeHttpDDoSAttackProtectionRequest
	GetSiteId() *int64
}

type DescribeHttpDDoSAttackProtectionRequest struct {
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSAttackProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackProtectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackProtectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeHttpDDoSAttackProtectionRequest) SetSiteId(v int64) *DescribeHttpDDoSAttackProtectionRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionRequest) Validate() error {
	return dara.Validate(s)
}
