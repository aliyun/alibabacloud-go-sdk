// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecurityProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainSecurityProfileResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeDomainSecurityProfileResponseBodyResult) *DescribeDomainSecurityProfileResponseBody
	GetResult() []*DescribeDomainSecurityProfileResponseBodyResult
}

type DescribeDomainSecurityProfileResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Result []*DescribeDomainSecurityProfileResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeDomainSecurityProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecurityProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecurityProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSecurityProfileResponseBody) GetResult() []*DescribeDomainSecurityProfileResponseBodyResult {
	return s.Result
}

func (s *DescribeDomainSecurityProfileResponseBody) SetRequestId(v string) *DescribeDomainSecurityProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSecurityProfileResponseBody) SetResult(v []*DescribeDomainSecurityProfileResponseBodyResult) *DescribeDomainSecurityProfileResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDomainSecurityProfileResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSecurityProfileResponseBodyResult struct {
	// Indicates whether the global mitigation policy is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	GlobalEnable *bool `json:"GlobalEnable,omitempty" xml:"GlobalEnable,omitempty"`
	// The mode of the global mitigation policy. Valid values:
	//
	// 	- **weak**: the Low mode
	//
	// 	- **default**: the Normal mode
	//
	// 	- **hard**: the Strict mode
	//
	// example:
	//
	// default
	GlobalMode *string `json:"GlobalMode,omitempty" xml:"GlobalMode,omitempty"`
}

func (s DescribeDomainSecurityProfileResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecurityProfileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecurityProfileResponseBodyResult) GetGlobalEnable() *bool {
	return s.GlobalEnable
}

func (s *DescribeDomainSecurityProfileResponseBodyResult) GetGlobalMode() *string {
	return s.GlobalMode
}

func (s *DescribeDomainSecurityProfileResponseBodyResult) SetGlobalEnable(v bool) *DescribeDomainSecurityProfileResponseBodyResult {
	s.GlobalEnable = &v
	return s
}

func (s *DescribeDomainSecurityProfileResponseBodyResult) SetGlobalMode(v string) *DescribeDomainSecurityProfileResponseBodyResult {
	s.GlobalMode = &v
	return s
}

func (s *DescribeDomainSecurityProfileResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
