// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainModes(v []*DescribeWebAccessModeResponseBodyDomainModes) *DescribeWebAccessModeResponseBody
	GetDomainModes() []*DescribeWebAccessModeResponseBodyDomainModes
	SetRequestId(v string) *DescribeWebAccessModeResponseBody
	GetRequestId() *string
}

type DescribeWebAccessModeResponseBody struct {
	// An array consisting of the modes in which the website service is added.
	DomainModes []*DescribeWebAccessModeResponseBodyDomainModes `json:"DomainModes,omitempty" xml:"DomainModes,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebAccessModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessModeResponseBody) GetDomainModes() []*DescribeWebAccessModeResponseBodyDomainModes {
	return s.DomainModes
}

func (s *DescribeWebAccessModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebAccessModeResponseBody) SetDomainModes(v []*DescribeWebAccessModeResponseBodyDomainModes) *DescribeWebAccessModeResponseBody {
	s.DomainModes = v
	return s
}

func (s *DescribeWebAccessModeResponseBody) SetRequestId(v string) *DescribeWebAccessModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebAccessModeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebAccessModeResponseBodyDomainModes struct {
	// The mode in which the website service is added. Valid values:
	//
	// 	- **0**: A record
	//
	// 	- **1**: anti-DDoS mode
	//
	// 	- **2**: origin redundancy mode
	//
	// example:
	//
	// 0
	AccessMode *int32 `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeWebAccessModeResponseBodyDomainModes) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessModeResponseBodyDomainModes) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessModeResponseBodyDomainModes) GetAccessMode() *int32 {
	return s.AccessMode
}

func (s *DescribeWebAccessModeResponseBodyDomainModes) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebAccessModeResponseBodyDomainModes) SetAccessMode(v int32) *DescribeWebAccessModeResponseBodyDomainModes {
	s.AccessMode = &v
	return s
}

func (s *DescribeWebAccessModeResponseBodyDomainModes) SetDomain(v string) *DescribeWebAccessModeResponseBodyDomainModes {
	s.Domain = &v
	return s
}

func (s *DescribeWebAccessModeResponseBodyDomainModes) Validate() error {
	return dara.Validate(s)
}
