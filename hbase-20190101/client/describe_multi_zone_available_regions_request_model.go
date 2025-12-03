// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiZoneAvailableRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeMultiZoneAvailableRegionsRequest
	GetAcceptLanguage() *string
}

type DescribeMultiZoneAvailableRegionsRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeMultiZoneAvailableRegionsRequest) SetAcceptLanguage(v string) *DescribeMultiZoneAvailableRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsRequest) Validate() error {
	return dara.Validate(s)
}
