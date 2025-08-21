// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *DescribeEpnInstanceAttributeRequest
	GetEPNInstanceId() *string
}

type DescribeEpnInstanceAttributeRequest struct {
	// The ID of the EPN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epn-xxxx
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
}

func (s DescribeEpnInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstanceAttributeRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *DescribeEpnInstanceAttributeRequest) SetEPNInstanceId(v string) *DescribeEpnInstanceAttributeRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
