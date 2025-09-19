// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProtectVpcListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddVpcInstanceIdList(v string) *AddProtectVpcListRequest
	GetAddVpcInstanceIdList() *string
	SetDelVpcInstanceIdList(v string) *AddProtectVpcListRequest
	GetDelVpcInstanceIdList() *string
}

type AddProtectVpcListRequest struct {
	// Collection of new VPC instance IDs.
	//
	// > Call the [DescribeVpcList](~~DescribeVpcList~~) interface to obtain this parameter.
	//
	// example:
	//
	// ["vpc-bp1vnpgotyzay6p5i****","vpc-bp1vnpgotyzay6p5i****"]
	AddVpcInstanceIdList *string `json:"AddVpcInstanceIdList,omitempty" xml:"AddVpcInstanceIdList,omitempty"`
	// Collection of VPC instance IDs to be deleted.
	//
	// > Call the [DescribeVpcList](~~DescribeVpcList~~) interface to obtain this parameter.
	//
	// example:
	//
	// ["vpc-bp1vnpgotyzay6p5i****","vpc-bp1vnpgotyzay6p5i****"]
	DelVpcInstanceIdList *string `json:"DelVpcInstanceIdList,omitempty" xml:"DelVpcInstanceIdList,omitempty"`
}

func (s AddProtectVpcListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddProtectVpcListRequest) GoString() string {
	return s.String()
}

func (s *AddProtectVpcListRequest) GetAddVpcInstanceIdList() *string {
	return s.AddVpcInstanceIdList
}

func (s *AddProtectVpcListRequest) GetDelVpcInstanceIdList() *string {
	return s.DelVpcInstanceIdList
}

func (s *AddProtectVpcListRequest) SetAddVpcInstanceIdList(v string) *AddProtectVpcListRequest {
	s.AddVpcInstanceIdList = &v
	return s
}

func (s *AddProtectVpcListRequest) SetDelVpcInstanceIdList(v string) *AddProtectVpcListRequest {
	s.DelVpcInstanceIdList = &v
	return s
}

func (s *AddProtectVpcListRequest) Validate() error {
	return dara.Validate(s)
}
