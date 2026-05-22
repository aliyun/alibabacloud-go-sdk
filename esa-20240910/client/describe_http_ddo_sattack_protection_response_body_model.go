// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSAttackProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalMode(v string) *DescribeHttpDDoSAttackProtectionResponseBody
	GetGlobalMode() *string
	SetRequestId(v string) *DescribeHttpDDoSAttackProtectionResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DescribeHttpDDoSAttackProtectionResponseBody
	GetSiteId() *int64
}

type DescribeHttpDDoSAttackProtectionResponseBody struct {
	// The level of HTTP DDoS attack protection. Valid values:
	//
	// 	- **very weak**: very loose.
	//
	// 	- **weak**: loose.
	//
	// 	- **default**: normal.
	//
	// 	- **hard**: strict.
	//
	// example:
	//
	// default
	GlobalMode *string `json:"GlobalMode,omitempty" xml:"GlobalMode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35C66C7B-671H-4297-9187-2C4477247A78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DescribeHttpDDoSAttackProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) GetGlobalMode() *string {
	return s.GlobalMode
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) SetGlobalMode(v string) *DescribeHttpDDoSAttackProtectionResponseBody {
	s.GlobalMode = &v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) SetRequestId(v string) *DescribeHttpDDoSAttackProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) SetSiteId(v int64) *DescribeHttpDDoSAttackProtectionResponseBody {
	s.SiteId = &v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
