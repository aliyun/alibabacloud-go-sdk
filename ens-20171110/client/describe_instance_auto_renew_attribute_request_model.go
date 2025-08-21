// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DescribeInstanceAutoRenewAttributeRequest
	GetInstanceIds() *string
	SetOwnerId(v int64) *DescribeInstanceAutoRenewAttributeRequest
	GetOwnerId() *int64
}

type DescribeInstanceAutoRenewAttributeRequest struct {
	// The ID of an instance. Separate multiple IDs with semicolons (;).
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5ci7l7k1m9m2zmhp4iw3o****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstanceAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetInstanceIds(v string) *DescribeInstanceAutoRenewAttributeRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeInstanceAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
